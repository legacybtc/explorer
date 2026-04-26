package explorer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Server is the block explorer HTTP server.
type Server struct {
	rpc    *RPCClient
	cache  *Cache
	tmpl   *template.Template
	port   int
	mux    *http.ServeMux
}

// NewServer creates a new explorer server.
func NewServer(rpc *RPCClient, port int) *Server {
	s := &Server{
		rpc:   rpc,
		cache: NewCache(),
		port:  port,
		mux:   http.NewServeMux(),
	}
	s.tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"formatTime":     formatTime,
		"formatLBTC":     formatLBTC,
		"formatHashrate": formatHashrate,
		"truncate":       truncate,
		"add":            func(a, b int64) int64 { return a + b },
		"sub":            func(a, b int64) int64 { return a - b },
		"blockReward":    blockRewardForHeight,
		"safeHTML":       func(s string) template.HTML { return template.HTML(s) },
	}).Parse(allTemplates))
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/", s.handleHome)
	s.mux.HandleFunc("/block/", s.handleBlock)
	s.mux.HandleFunc("/blocks", s.handleBlocks)
	s.mux.HandleFunc("/search", s.handleSearch)
	s.mux.HandleFunc("/api/stats", s.handleAPIStats)
	s.mux.HandleFunc("/api/overview", s.handleAPIOverview)
	s.mux.HandleFunc("/api/blocks", s.handleAPIBlocks)
	s.mux.HandleFunc("/api/block/", s.handleAPIBlock)
}

// Start begins serving HTTP requests.
func (s *Server) Start() {
	log.Printf("Block explorer listening on http://0.0.0.0:%d", s.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux); err != nil {
		log.Fatalf("Explorer server error: %v", err)
	}
}

// ── Page handlers ─────────────────────────────────────────────────────────────

type homeData struct {
	NodeOnline   bool
	Info         *NodeInfo
	Mining       *MiningInfo
	RecentBlocks []*Block
	TipBlock     *Block
	Error        string
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := homeData{}

	if s.rpc.Ping() {
		data.NodeOnline = true
		if info, err := s.cachedInfo(); err == nil {
			data.Info = info
		}
		if mining, err := s.cachedMining(); err == nil {
			data.Mining = mining
		}
		if blocks, err := s.cachedRecentBlocks(20); err == nil {
			data.RecentBlocks = blocks
			if len(blocks) > 0 {
				data.TipBlock = blocks[0]
			}
		}
	} else {
		data.Error = "Cannot connect to legacycoind node. Is it running?"
	}

	s.render(w, "home", data)
}

func (s *Server) handleBlocks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page := int64(1)
	if p, err := strconv.ParseInt(pageStr, 10, 64); err == nil && p > 0 {
		page = p
	}

	perPage := int64(50)
	tip, _ := s.rpc.GetBlockCount()
	start := tip - (page-1)*perPage
	end := start - perPage + 1
	if end < 0 {
		end = 0
	}

	var blocks []*Block
	for h := start; h >= end; h-- {
		b, err := s.rpc.GetBlockAtHeight(h)
		if err != nil {
			break
		}
		b.Confirmations = tip - h + 1
		blocks = append(blocks, b)
	}

	s.render(w, "blocks", map[string]interface{}{
		"Blocks":   blocks,
		"Page":     page,
		"PrevPage": page - 1,
		"NextPage": page + 1,
		"HasPrev":  page > 1,
		"HasNext":  end > 0,
		"Tip":      tip,
	})
}

func (s *Server) handleBlock(w http.ResponseWriter, r *http.Request) {
	identifier := strings.TrimPrefix(r.URL.Path, "/block/")
	identifier = strings.TrimSpace(identifier)

	var block *Block
	var err error

	// Try as height first, then as hash
	if height, perr := strconv.ParseInt(identifier, 10, 64); perr == nil {
		block, err = s.rpc.GetBlockAtHeight(height)
	} else {
		block, err = s.rpc.GetBlock(identifier)
	}

	if err != nil {
		s.render(w, "error", map[string]interface{}{
			"Message": fmt.Sprintf("Block not found: %s", identifier),
		})
		return
	}

	tip, _ := s.rpc.GetBlockCount()
	block.Confirmations = tip - block.Height + 1

	s.render(w, "block", map[string]interface{}{
		"Block":  block,
		"Reward": blockRewardForHeight(block.Height),
	})
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// Try as block height
	if height, err := strconv.ParseInt(q, 10, 64); err == nil {
		http.Redirect(w, r, fmt.Sprintf("/block/%d", height), http.StatusFound)
		return
	}
	// Try as block hash (64 hex chars)
	if len(q) == 64 {
		http.Redirect(w, r, fmt.Sprintf("/block/%s", q), http.StatusFound)
		return
	}
	s.render(w, "error", map[string]interface{}{
		"Message": fmt.Sprintf("Not found: %s. Enter a block height or block hash.", q),
	})
}

// ── API handlers (JSON) ───────────────────────────────────────────────────────

func (s *Server) handleAPIStats(w http.ResponseWriter, r *http.Request) {
	info, err1 := s.cachedInfo()
	mining, err2 := s.cachedMining()
	if err1 != nil || err2 != nil {
		jsonError(w, "node unavailable", 503)
		return
	}
	jsonOK(w, map[string]interface{}{
		"blocks":       info.Blocks,
		"connections":  info.Connections,
		"difficulty":   info.Difficulty,
		"hashrate":     mining.HashesPerSec,
		"pooled_tx":    mining.PooledTx,
		"node_version": info.Version,
	})
}

func (s *Server) handleAPIOverview(w http.ResponseWriter, r *http.Request) {
	info, err1 := s.cachedInfo()
	mining, err2 := s.cachedMining()
	blocks, err3 := s.cachedRecentBlocks(10)
	if err1 != nil || err2 != nil || err3 != nil {
		jsonError(w, "node unavailable", 503)
		return
	}

	payload := map[string]interface{}{
		"network": map[string]interface{}{
			"name":      "LegacyCoin",
			"ticker":    "LBTC",
			"consensus": "Yespower 1.0",
			"retarget":  "Dark Gravity Wave v3",
			"supplyCap": "21M",
		},
		"stats": map[string]interface{}{
			"blocks":       info.Blocks,
			"connections":  info.Connections,
			"difficulty":   info.Difficulty,
			"hashrate":     mining.HashesPerSec,
			"hashrateText": formatHashrate(mining.HashesPerSec),
			"pooledTx":     mining.PooledTx,
			"nodeVersion":  info.Version,
			"mining":       mining.Generate,
		},
		"blocks": blocks,
	}

	if len(blocks) > 0 {
		payload["tip"] = map[string]interface{}{
			"height":      blocks[0].Height,
			"hash":        blocks[0].Hash,
			"time":        formatTime(blocks[0].Time),
			"reward":      formatLBTC(blockRewardForHeight(blocks[0].Height)),
			"size":        blocks[0].Size,
			"txCount":     len(blocks[0].Tx),
			"confirmations": blocks[0].Confirmations,
		}
	}

	jsonOK(w, payload)
}

func (s *Server) handleAPIBlocks(w http.ResponseWriter, r *http.Request) {
	blocks, err := s.cachedRecentBlocks(10)
	if err != nil {
		jsonError(w, "node unavailable", 503)
		return
	}
	jsonOK(w, blocks)
}

func (s *Server) handleAPIBlock(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/block/")
	var block *Block
	var err error
	if height, perr := strconv.ParseInt(id, 10, 64); perr == nil {
		block, err = s.rpc.GetBlockAtHeight(height)
	} else {
		block, err = s.rpc.GetBlock(id)
	}
	if err != nil {
		jsonError(w, "block not found", 404)
		return
	}
	jsonOK(w, block)
}

// ── Cache helpers ─────────────────────────────────────────────────────────────

func (s *Server) cachedInfo() (*NodeInfo, error) {
	if v, ok := s.cache.Get("info"); ok {
		return v.(*NodeInfo), nil
	}
	info, err := s.rpc.GetInfo()
	if err != nil {
		return nil, err
	}
	s.cache.Set("info", info, 5*time.Second)
	return info, nil
}

func (s *Server) cachedMining() (*MiningInfo, error) {
	if v, ok := s.cache.Get("mining"); ok {
		return v.(*MiningInfo), nil
	}
	mining, err := s.rpc.GetMiningInfo()
	if err != nil {
		return nil, err
	}
	s.cache.Set("mining", mining, 5*time.Second)
	return mining, nil
}

func (s *Server) cachedRecentBlocks(n int) ([]*Block, error) {
	key := fmt.Sprintf("blocks:%d", n)
	if v, ok := s.cache.Get(key); ok {
		return v.([]*Block), nil
	}
	blocks, err := s.rpc.GetRecentBlocks(n)
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, blocks, 10*time.Second)
	return blocks, nil
}

// ── Render helper ─────────────────────────────────────────────────────────────

func (s *Server) render(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := s.tmpl.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("Template error [%s]: %v", name, err)
		http.Error(w, "render error", 500)
	}
}

// ── Utility ───────────────────────────────────────────────────────────────────

func formatTime(unix uint32) string {
	t := time.Unix(int64(unix), 0).UTC()
	return t.Format("2006-01-02 15:04:05 UTC")
}

func formatLBTC(satoshis int64) string {
	lbtc := float64(satoshis) / 1e8
	return fmt.Sprintf("%.8f LBTC", lbtc)
}

func formatHashrate(hs int64) string {
	rate := float64(hs)
	switch {
	case rate >= 1e9:
		return fmt.Sprintf("%.2f GH/s", rate/1e9)
	case rate >= 1e6:
		return fmt.Sprintf("%.2f MH/s", rate/1e6)
	case rate >= 1e3:
		return fmt.Sprintf("%.2f kH/s", rate/1e3)
	default:
		return fmt.Sprintf("%d H/s", hs)
	}
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

func blockRewardForHeight(height int64) int64 {
	halvings := height / 210_000
	if halvings >= 64 {
		return 0
	}
	return (50 * 1e8) >> halvings
}

func jsonOK(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(v)
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
