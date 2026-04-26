package explorer

const sharedCSS = `:root{
  --bg:#08090d;
  --bg-soft:#11141a;
  --bg-deep:#0d1016;
  --panel:rgba(18,22,31,.92);
  --panel-strong:#151922;
  --panel-alt:rgba(255,255,255,.03);
  --line:rgba(255,255,255,.08);
  --line-strong:rgba(203,163,106,.22);
  --text:#f3f0ea;
  --muted:#bdb6aa;
  --muted-2:#8d877d;
  --gold:#cba36a;
  --gold-strong:#e0b77a;
  --cream:#f2ece0;
  --green:#44d78b;
  --red:#ef6a73;
  --blue:#87a8ff;
  --shadow:0 26px 80px rgba(0,0,0,.36);
  --radius:22px;
  --radius-sm:14px;
  --mono:'JetBrains Mono',monospace;
  --sans:'Space Grotesk',system-ui,sans-serif;
  --serif:'Cormorant Garamond',serif;
}
*{box-sizing:border-box;margin:0;padding:0}
html{scroll-behavior:smooth}
body{
  min-height:100vh;
  background:
    radial-gradient(circle at top left, rgba(203,163,106,.14), transparent 28%),
    radial-gradient(circle at 84% 14%, rgba(98,116,164,.12), transparent 26%),
    linear-gradient(180deg, #090b10 0%, #0d1016 42%, #090b10 100%);
  color:var(--text);
  font-family:var(--sans);
}
a{color:inherit;text-decoration:none}
a:hover{text-decoration:none}
.page-shell{position:relative;overflow:hidden;min-height:100vh}
.ambient{position:absolute;border-radius:999px;filter:blur(78px);opacity:.32;pointer-events:none}
.ambient-a{width:360px;height:360px;top:120px;left:-90px;background:rgba(203,163,106,.18)}
.ambient-b{width:420px;height:420px;top:520px;right:-130px;background:rgba(83,96,141,.20)}
.topbar{
  position:sticky;top:0;z-index:20;
  display:flex;align-items:center;justify-content:space-between;gap:24px;
  padding:18px 34px;
  background:rgba(9,11,16,.78);
  backdrop-filter:blur(18px);
  border-bottom:1px solid rgba(255,255,255,.05)
}
.brand{display:inline-flex;align-items:center;gap:14px;min-width:0}
.brand-mark{
  width:48px;height:48px;border-radius:14px;
  display:grid;place-items:center;
  background:
    linear-gradient(135deg, rgba(203,163,106,.18), rgba(255,255,255,.02)),
    var(--panel-strong);
  border:1px solid var(--line-strong);
  color:var(--cream);
  font-family:var(--serif);
  font-size:34px;
  font-weight:700;
  box-shadow:var(--shadow)
}
.brand-copy{display:flex;flex-direction:column;line-height:1}
.brand-copy strong{font-size:1rem;letter-spacing:.04em}
.brand-copy strong em{font-style:normal;color:var(--gold)}
.brand-copy span{
  margin-top:6px;
  font-size:.72rem;
  color:var(--muted);
  text-transform:uppercase;
  letter-spacing:.24em
}
.topnav{display:flex;gap:22px;color:var(--muted);font-size:.95rem}
.topnav a{padding:6px 0;position:relative}
.topnav a:hover,.topnav a.active{color:var(--cream)}
.topnav a.active::after{
  content:'';position:absolute;left:0;right:0;bottom:-14px;height:2px;
  border-radius:999px;background:linear-gradient(90deg,var(--gold),transparent)
}
.search-form{
  display:flex;align-items:center;overflow:hidden;
  min-width:min(360px,36vw);max-width:420px;
  border-radius:999px;border:1px solid var(--line);
  background:rgba(255,255,255,.03)
}
.search-form input{
  flex:1;min-width:0;background:none;border:none;outline:none;
  color:var(--text);font:500 13px var(--mono);padding:14px 16px
}
.search-form input::placeholder{color:var(--muted-2);font-family:var(--sans)}
.search-form button{
  border:none;background:linear-gradient(135deg,#d5ad73,#b88747);
  color:#17120d;font-weight:800;padding:14px 18px;cursor:pointer
}
.page{position:relative;z-index:1;width:min(1320px,calc(100% - 40px));margin:0 auto;padding:30px 0 60px}
.hero{
  display:grid;grid-template-columns:minmax(0,1.1fr) minmax(320px,.9fr);
  gap:22px;align-items:stretch;margin-bottom:22px
}
.hero-card,.panel,.stat,.table-wrap,.footer-card,.notice,.empty-card{
  border:1px solid var(--line);
  background:var(--panel);
  box-shadow:var(--shadow)
}
.hero-card{
  border-radius:28px;padding:32px;
  background:
    linear-gradient(160deg, rgba(203,163,106,.10), rgba(255,255,255,.02)),
    var(--panel-strong)
}
.eyebrow{
  display:inline-flex;align-items:center;gap:9px;
  color:var(--gold);font-size:.74rem;font-weight:700;
  text-transform:uppercase;letter-spacing:.26em
}
.eyebrow::before{
  content:'';width:30px;height:1px;background:linear-gradient(90deg,var(--gold),transparent)
}
.hero h1,.section-heading h2{
  margin:18px 0 0;
  font-family:var(--serif);
  line-height:.94;
  letter-spacing:-.04em
}
.hero h1{font-size:clamp(3rem,7vw,5.8rem);max-width:8ch}
.hero p{
  max-width:62ch;margin-top:22px;
  color:var(--muted);font-size:1.04rem;line-height:1.75
}
.hero-actions{display:flex;gap:12px;flex-wrap:wrap;margin-top:30px}
.hero-btn{
  display:inline-flex;align-items:center;justify-content:center;
  min-height:48px;padding:0 20px;border-radius:999px;
  font-size:.92rem;font-weight:700;transition:transform .18s ease
}
.hero-btn:hover{transform:translateY(-1px)}
.hero-btn.primary{background:linear-gradient(135deg,#d5ad73,#b88747);color:#16120d}
.hero-btn.secondary{border:1px solid var(--line);background:rgba(255,255,255,.03);color:var(--cream)}
.hero-side{display:grid;gap:16px}
.mini-card{
  border-radius:24px;padding:24px;
  background:
    linear-gradient(160deg, rgba(255,255,255,.03), rgba(255,255,255,.01)),
    var(--panel)
}
.mini-label,.panel-title{
  color:var(--muted);font-size:.72rem;font-weight:700;
  text-transform:uppercase;letter-spacing:.18em
}
.mini-value{margin-top:12px;font-family:var(--mono);font-size:1.5rem;font-weight:700;color:var(--cream)}
.mini-copy{margin-top:10px;color:var(--muted);font-size:.95rem;line-height:1.65}
.mini-grid{display:grid;grid-template-columns:1fr 1fr;gap:14px}
.stats-grid{
  display:grid;grid-template-columns:repeat(6,minmax(0,1fr));gap:14px;
  margin-bottom:22px
}
.stat{border-radius:20px;padding:18px 18px 16px;background:rgba(18,22,31,.88)}
.stat-label{
  color:var(--muted);font-size:.72rem;font-weight:700;
  text-transform:uppercase;letter-spacing:.18em
}
.stat-value{
  margin-top:11px;font-family:var(--mono);
  font-size:1.25rem;font-weight:700;color:var(--cream)
}
.stat-sub{margin-top:8px;color:var(--muted-2);font-size:.82rem;line-height:1.5}
.status-live{color:var(--green)}
.status-live::before{content:'● ';font-size:.9em}
.status-offline{color:var(--red)}
.status-offline::before{content:'● ';font-size:.9em}
.section-grid{display:grid;grid-template-columns:minmax(0,1.55fr) minmax(320px,.85fr);gap:18px;margin-bottom:20px}
.panel{border-radius:24px;padding:22px}
.panel-head{
  display:flex;align-items:flex-end;justify-content:space-between;gap:12px;
  margin-bottom:16px
}
.panel-copy h2,.panel-copy h3{font-family:var(--serif);font-size:2rem;line-height:1;margin:0}
.panel-copy p{margin-top:8px;color:var(--muted);font-size:.95rem}
.panel-link{color:var(--gold);font-size:.88rem;font-weight:700}
.table-wrap{border-radius:22px;overflow:hidden}
table{width:100%;border-collapse:collapse}
thead th{
  background:rgba(255,255,255,.03);color:var(--muted);
  font-size:.7rem;font-weight:700;text-transform:uppercase;
  letter-spacing:.18em;text-align:left;padding:14px 16px
}
tbody td{
  padding:14px 16px;border-top:1px solid rgba(255,255,255,.05);
  font-size:.92rem;vertical-align:middle
}
tbody tr:hover td{background:rgba(255,255,255,.02)}
.mono{font-family:var(--mono)}
.muted{color:var(--muted)}
.gold{color:var(--gold)}
.badge{
  display:inline-flex;align-items:center;gap:6px;
  padding:5px 10px;border-radius:999px;
  font-size:.72rem;font-weight:700;letter-spacing:.08em;
  text-transform:uppercase
}
.badge.ok{background:rgba(68,215,139,.10);color:var(--green);border:1px solid rgba(68,215,139,.18)}
.badge.warn{background:rgba(203,163,106,.10);color:var(--gold);border:1px solid rgba(203,163,106,.18)}
.badge.err{background:rgba(239,106,115,.10);color:var(--red);border:1px solid rgba(239,106,115,.18)}
.stack{display:grid;gap:14px}
.stack-card{
  background:rgba(255,255,255,.025);
  border:1px solid var(--line);
  border-radius:18px;padding:16px 18px
}
.stack-row{display:flex;align-items:flex-start;justify-content:space-between;gap:16px;padding:10px 0;border-top:1px solid rgba(255,255,255,.05)}
.stack-row:first-child{border-top:none;padding-top:0}
.stack-key{
  color:var(--muted);font-size:.74rem;font-weight:700;
  text-transform:uppercase;letter-spacing:.16em
}
.stack-val{font-size:.93rem;text-align:right;color:var(--cream);word-break:break-word}
.stack-val a{color:var(--gold)}
.section-heading{display:flex;align-items:flex-end;justify-content:space-between;gap:16px;margin-bottom:16px}
.section-heading p{max-width:56ch;color:var(--muted);font-size:.95rem;line-height:1.7}
.empty-row{text-align:center;color:var(--muted);padding:28px 16px}
.pager{display:flex;gap:10px;flex-wrap:wrap;margin-top:18px}
.pager a,.pager span{
  padding:10px 14px;border-radius:999px;border:1px solid var(--line);
  background:rgba(255,255,255,.03);font-size:.88rem;color:var(--muted)
}
.pager .cur{border-color:var(--line-strong);color:var(--cream)}
.block-hero{
  margin-bottom:18px;padding:28px;border-radius:28px;
  background:
    linear-gradient(160deg, rgba(203,163,106,.10), rgba(255,255,255,.02)),
    var(--panel-strong)
}
.back-link{display:inline-flex;align-items:center;gap:8px;color:var(--muted);font-size:.88rem}
.back-link:hover{color:var(--cream)}
.block-title{margin-top:18px;font-family:var(--serif);font-size:3.35rem;line-height:.92}
.block-sub{margin-top:12px;color:var(--muted);font:500 .88rem/1.7 var(--mono);word-break:break-all}
.block-summary{
  display:grid;grid-template-columns:repeat(4,minmax(0,1fr));
  gap:12px;margin-top:24px
}
.summary-pill{
  border:1px solid var(--line);border-radius:18px;
  background:rgba(255,255,255,.03);padding:14px 16px
}
.summary-pill span{
  display:block;color:var(--muted);font-size:.7rem;font-weight:700;
  text-transform:uppercase;letter-spacing:.16em
}
.summary-pill strong{
  display:block;margin-top:10px;color:var(--cream);
  font:700 1rem/1.3 var(--mono)
}
.detail-grid{display:grid;grid-template-columns:1fr 1fr;gap:18px;margin-bottom:18px}
.detail-card{border-radius:24px;padding:22px;background:var(--panel);border:1px solid var(--line)}
.detail-card h3{
  font-size:.74rem;font-weight:700;text-transform:uppercase;
  letter-spacing:.18em;color:var(--muted);margin-bottom:16px
}
.detail-row{display:flex;justify-content:space-between;gap:14px;padding:11px 0;border-top:1px solid rgba(255,255,255,.05)}
.detail-row:first-child{border-top:none;padding-top:0}
.detail-key{color:var(--muted);font-size:.8rem;text-transform:uppercase;letter-spacing:.12em}
.detail-val{color:var(--cream);font:.9rem/1.6 var(--mono);text-align:right;word-break:break-all}
.detail-val a{color:var(--gold)}
.tx-section{margin-top:20px}
.notice{
  border-radius:22px;padding:18px 20px;margin-bottom:18px;
  background:rgba(239,106,115,.08);border-color:rgba(239,106,115,.18);color:#ffd7db
}
.empty-card{
  max-width:540px;margin:90px auto;padding:34px;border-radius:28px;text-align:center
}
.empty-card h2{font-family:var(--serif);font-size:2.4rem;color:var(--cream)}
.empty-card p{margin-top:14px;color:var(--muted);line-height:1.8}
.empty-card a{
  display:inline-flex;align-items:center;justify-content:center;
  margin-top:22px;min-height:46px;padding:0 20px;border-radius:999px;
  background:linear-gradient(135deg,#d5ad73,#b88747);color:#16120d;font-weight:700
}
.footer-card{
  margin-top:22px;border-radius:24px;padding:20px 24px;
  background:rgba(255,255,255,.025);text-align:center;color:var(--muted);font-size:.92rem
}
.footer-card span{color:var(--gold)}
@media(max-width:1180px){
  .hero,.section-grid,.detail-grid{grid-template-columns:1fr}
  .block-summary,.stats-grid{grid-template-columns:repeat(3,minmax(0,1fr))}
}
@media(max-width:860px){
  .topbar{padding:16px 18px;flex-wrap:wrap}
  .search-form{min-width:100%;max-width:none;order:3}
  .page{width:min(100%,calc(100% - 24px))}
  .hero-card,.block-hero{padding:24px}
  .hero h1{font-size:clamp(2.7rem,12vw,4.5rem)}
  .block-title{font-size:2.5rem}
  .block-summary,.stats-grid{grid-template-columns:repeat(2,minmax(0,1fr))}
}
@media(max-width:560px){
  .topnav{gap:14px;font-size:.88rem}
  .brand-copy span{letter-spacing:.18em}
  .hero-actions{flex-direction:column}
  .hero-btn{width:100%}
  .mini-grid,.block-summary,.stats-grid{grid-template-columns:1fr}
  .panel-head,.section-heading{flex-direction:column;align-items:flex-start}
  thead th,tbody td{padding:12px 10px}
}
`

const navSnip = `<div class="page-shell"><div class="ambient ambient-a"></div><div class="ambient ambient-b"></div><header class="topbar"><a href="/" class="brand"><span class="brand-mark">L</span><span class="brand-copy"><strong><em>LegacyCoin</em> Explorer</strong><span>CPU-secured chain view</span></span></a><nav class="topnav"><a href="/" class="active">Dashboard</a><a href="/blocks">Blocks</a></nav><form class="search-form" action="/search" method="GET"><input type="text" name="q" placeholder="Block, txid, or wallet address"><button type="submit">Explore</button></form></header>`
const footSnip = `<div class="page footer-card"><span>LegacyCoin (LBTC)</span> public explorer · one CPU, one vote</div></div>`
const headSnip = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin><link href="https://fonts.googleapis.com/css2?family=Cormorant+Garamond:wght@600;700&family=JetBrains+Mono:wght@400;500;700&family=Space+Grotesk:wght@400;500;700&display=swap" rel="stylesheet"><style>` + sharedCSS + `</style>`

const allTemplates = `
{{define "home"}}` + headSnip + `<title>LegacyCoin Explorer</title></head><body>` + navSnip + `<main class="page">
{{if .Error}}<div class="notice">{{.Error}}</div>{{end}}
<section class="hero">
  <div class="hero-card">
    <div class="eyebrow">Live chain intelligence</div>
    <h1>LegacyCoin Explorer</h1>
    <p>Track the LBTC network with a cleaner public view of chain height, block flow, mining activity, confirmations, and on-chain history. The layout follows the LegacyCoin website and whitepaper language instead of the older bare explorer style.</p>
    <div class="hero-actions">
      <a class="hero-btn primary" href="/blocks">Browse blocks</a>
      <a class="hero-btn secondary" href="/search?q={{if .Info}}{{.Info.Blocks}}{{else}}0{{end}}">Jump to tip</a>
    </div>
  </div>
  <div class="hero-side">
    <div class="mini-card">
      <div class="mini-label">Chain posture</div>
      <div class="mini-value">{{if .NodeOnline}}Online{{else}}Offline{{end}}</div>
      <div class="mini-copy">A public chain view for LegacyCoin’s fixed-supply, Yespower-mined network with DGW3 retargeting and transparent block visibility.</div>
    </div>
    <div class="mini-grid">
      <div class="mini-card">
        <div class="mini-label">Network goal</div>
        <div class="mini-value">21M</div>
        <div class="mini-copy">Hard cap with Bitcoin-style emission and long-horizon halvings.</div>
      </div>
      <div class="mini-card">
        <div class="mini-label">Security model</div>
        <div class="mini-value">CPU</div>
        <div class="mini-copy">Focused on accessible mining and simple public verification.</div>
      </div>
    </div>
  </div>
</section>
<section class="stats-grid">
  <div class="stat">
    <div class="stat-label">Status</div>
    <div class="stat-value {{if .NodeOnline}}status-live{{else}}status-offline{{end}}">{{if .NodeOnline}}Node online{{else}}Node offline{{end}}</div>
    <div class="stat-sub">{{if .NodeOnline}}RPC is responding and chain data is available.{{else}}Explorer cannot reach legacycoind right now.{{end}}</div>
  </div>
  <div class="stat">
    <div class="stat-label">Chain height</div>
    <div class="stat-value">{{if .Info}}{{.Info.Blocks}}{{else}}—{{end}}</div>
    <div class="stat-sub">Current best block known to the connected node.</div>
  </div>
  <div class="stat">
    <div class="stat-label">Difficulty</div>
    <div class="stat-value">{{if .Info}}{{printf "%.9f" .Info.Difficulty}}{{else}}—{{end}}</div>
    <div class="stat-sub">Dynamic per-block retargeting under DGW3.</div>
  </div>
  <div class="stat">
    <div class="stat-label">Peers</div>
    <div class="stat-value">{{if .Info}}{{.Info.Connections}}{{else}}—{{end}}</div>
    <div class="stat-sub">Active node connections visible to the explorer backend.</div>
  </div>
  <div class="stat">
    <div class="stat-label">Hashrate</div>
    <div class="stat-value">{{if .Mining}}{{formatHashrate .Mining.HashesPerSec}}{{else}}—{{end}}</div>
    <div class="stat-sub">Local mining sample reported by the connected node.</div>
  </div>
  <div class="stat">
    <div class="stat-label">Mempool</div>
    <div class="stat-value">{{if .Mining}}{{.Mining.PooledTx}}{{else}}—{{end}}</div>
    <div class="stat-sub">Pending transactions waiting for inclusion in a block.</div>
  </div>
</section>
<section class="section-grid">
  <div class="panel">
    <div class="panel-head">
      <div class="panel-copy">
        <h2>Latest blocks</h2>
        <p>Recent chain activity with block reward, size, transaction count, and confirmations.</p>
      </div>
      <a class="panel-link" href="/blocks">Open full block list</a>
    </div>
    <div class="table-wrap">
      <table>
        <thead><tr><th>Height</th><th>Hash</th><th>Time (UTC)</th><th>Txs</th><th>Reward</th><th>Size</th><th>Confs</th></tr></thead>
        <tbody>
          {{range .RecentBlocks}}<tr>
            <td><a class="gold mono" href="/block/{{.Height}}">{{.Height}}</a></td>
            <td class="mono"><a class="gold" href="/block/{{.Hash}}">{{truncate .Hash 32}}</a></td>
            <td class="muted">{{formatTime .Time}}</td>
            <td>{{len .Tx}}</td>
            <td class="gold mono">{{formatLBTC (blockReward .Height)}}</td>
            <td class="muted">{{.Size}} B</td>
            <td><span class="badge ok">{{.Confirmations}}</span></td>
          </tr>{{else}}<tr><td class="empty-row" colspan="7">No blocks available yet.</td></tr>{{end}}
        </tbody>
      </table>
    </div>
  </div>
  <div class="stack">
    <div class="panel">
      <div class="panel-head">
        <div class="panel-copy">
          <h3>Network summary</h3>
          <p>High-level chain and node facts from the current RPC endpoint.</p>
        </div>
      </div>
      <div class="stack-card">
        <div class="stack-row"><div class="stack-key">Consensus</div><div class="stack-val">Yespower 1.0</div></div>
        <div class="stack-row"><div class="stack-key">Retarget</div><div class="stack-val">Dark Gravity Wave v3</div></div>
        <div class="stack-row"><div class="stack-key">Emission</div><div class="stack-val">50 LBTC base reward</div></div>
        <div class="stack-row"><div class="stack-key">Node version</div><div class="stack-val">{{if .Info}}{{.Info.Version}}{{else}}—{{end}}</div></div>
        <div class="stack-row"><div class="stack-key">Mining</div><div class="stack-val">{{if .Mining}}{{if .Mining.Generate}}Enabled{{else}}Disabled{{end}}{{else}}—{{end}}</div></div>
        <div class="stack-row"><div class="stack-key">Explorer mode</div><div class="stack-val">RPC direct</div></div>
      </div>
    </div>
    <div class="panel">
      <div class="panel-head">
        <div class="panel-copy">
          <h3>Search tips</h3>
          <p>The current explorer resolves block heights and full block hashes.</p>
        </div>
      </div>
      <div class="stack-card">
        <div class="stack-row"><div class="stack-key">By height</div><div class="stack-val">Search <span class="gold mono">{{if .Info}}{{.Info.Blocks}}{{else}}336{{end}}</span> to open the latest block.</div></div>
        <div class="stack-row"><div class="stack-key">By hash</div><div class="stack-val">Paste a 64-character block hash or transaction ID into the search bar.</div></div>
        <div class="stack-row"><div class="stack-key">By address</div><div class="stack-val">Paste a LegacyCoin wallet address to view its current on-chain balance.</div></div>
      </div>
    </div>
  </div>
</section>
</main>` + footSnip + `</body></html>{{end}}

{{define "blocks"}}` + headSnip + `<title>All Blocks — LegacyCoin Explorer</title></head><body>` + navSnip + `<main class="page">
<section class="section-heading">
  <div>
    <div class="eyebrow">Chain history</div>
    <h2>All blocks</h2>
    <p>Browse confirmed blocks from the current node tip backward through LegacyCoin chain history.</p>
  </div>
  <div class="mini-card" style="min-width:240px">
    <div class="mini-label">Current tip</div>
    <div class="mini-value">{{.Tip}}</div>
    <div class="mini-copy">Paginated newest-first block listing from the connected node.</div>
  </div>
</section>
<div class="table-wrap">
  <table>
    <thead><tr><th>Height</th><th>Hash</th><th>Time (UTC)</th><th>Txs</th><th>Reward</th><th>Size</th><th>Confs</th></tr></thead>
    <tbody>
      {{range .Blocks}}<tr>
        <td><a class="gold mono" href="/block/{{.Height}}">{{.Height}}</a></td>
        <td class="mono"><a class="gold" href="/block/{{.Hash}}">{{truncate .Hash 40}}</a></td>
        <td class="muted">{{formatTime .Time}}</td>
        <td>{{len .Tx}}</td>
        <td class="gold mono">{{formatLBTC (blockReward .Height)}}</td>
        <td class="muted">{{.Size}} B</td>
        <td><span class="badge ok">{{.Confirmations}}</span></td>
      </tr>{{else}}<tr><td class="empty-row" colspan="7">No blocks available.</td></tr>{{end}}
    </tbody>
  </table>
</div>
<div class="pager">
  {{if .HasPrev}}<a href="/blocks?page={{.PrevPage}}">← Newer</a>{{else}}<span>← Newer</span>{{end}}
  <span class="cur">Page {{.Page}}</span>
  {{if .HasNext}}<a href="/blocks?page={{.NextPage}}">Older →</a>{{else}}<span>Older →</span>{{end}}
</div>
</main>` + footSnip + `</body></html>{{end}}

{{define "block"}}` + headSnip + `<title>Block {{.Block.Height}} — LegacyCoin Explorer</title></head><body>` + navSnip + `<main class="page">
<section class="block-hero">
  <a class="back-link" href="/blocks">← Back to blocks</a>
  <div class="eyebrow" style="margin-top:20px">Block detail</div>
  <div class="block-title">Block #{{.Block.Height}}</div>
  <div class="block-sub">{{.Block.Hash}}</div>
  <div class="block-summary">
    <div class="summary-pill"><span>Confirmations</span><strong>{{.Block.Confirmations}}</strong></div>
    <div class="summary-pill"><span>Transactions</span><strong>{{len .Block.Tx}}</strong></div>
    <div class="summary-pill"><span>Reward</span><strong>{{formatLBTC .Reward}}</strong></div>
    <div class="summary-pill"><span>Size</span><strong>{{.Block.Size}} B</strong></div>
  </div>
</section>
<section class="detail-grid">
  <div class="detail-card">
    <h3>Header</h3>
    <div class="detail-row"><div class="detail-key">Height</div><div class="detail-val">{{.Block.Height}}</div></div>
    <div class="detail-row"><div class="detail-key">Hash</div><div class="detail-val">{{.Block.Hash}}</div></div>
    <div class="detail-row"><div class="detail-key">Previous</div><div class="detail-val">{{if .Block.PreviousBlockHash}}<a href="/block/{{.Block.PreviousBlockHash}}">{{.Block.PreviousBlockHash}}</a>{{else}}Genesis{{end}}</div></div>
    <div class="detail-row"><div class="detail-key">Merkle root</div><div class="detail-val">{{.Block.MerkleRoot}}</div></div>
    <div class="detail-row"><div class="detail-key">Time</div><div class="detail-val">{{formatTime .Block.Time}}</div></div>
  </div>
  <div class="detail-card">
    <h3>Consensus summary</h3>
    <div class="detail-row"><div class="detail-key">Bits</div><div class="detail-val">{{.Block.Bits}}</div></div>
    <div class="detail-row"><div class="detail-key">Nonce</div><div class="detail-val">{{.Block.Nonce}}</div></div>
    <div class="detail-row"><div class="detail-key">Version</div><div class="detail-val">{{.Block.Version}}</div></div>
    <div class="detail-row"><div class="detail-key">Algorithm</div><div class="detail-val">Yespower 1.0</div></div>
    <div class="detail-row"><div class="detail-key">Retarget</div><div class="detail-val">DGW3</div></div>
  </div>
</section>
<section class="tx-section">
  <div class="section-heading">
    <div>
      <div class="eyebrow">Included data</div>
      <h2>Transactions</h2>
      <p>All transaction IDs included in this block, with the coinbase entry highlighted at the top.</p>
    </div>
  </div>
  <div class="table-wrap">
    <table>
      <thead><tr><th>#</th><th>TXID</th><th>Type</th></tr></thead>
      <tbody>
        {{range $i, $tx := .Block.Tx}}<tr>
          <td class="muted">{{$i}}</td>
          <td class="mono"><a class="gold" href="/tx/{{$tx}}">{{$tx}}</a></td>
          <td>{{if eq $i 0}}<span class="badge warn">Coinbase</span>{{else}}<span class="badge ok">Transfer</span>{{end}}</td>
        </tr>{{else}}<tr><td class="empty-row" colspan="3">No transactions recorded.</td></tr>{{end}}
      </tbody>
    </table>
  </div>
  <div class="pager">
    {{if gt .Block.Height 0}}<a href="/block/{{sub .Block.Height 1}}">← Block {{sub .Block.Height 1}}</a>{{end}}
    <a href="/block/{{add .Block.Height 1}}">Block {{add .Block.Height 1}} →</a>
  </div>
</section>
</main>` + footSnip + `</body></html>{{end}}

{{define "error"}}` + headSnip + `<title>Not Found — LegacyCoin Explorer</title></head><body>` + navSnip + `<main class="page"><div class="empty-card"><div class="eyebrow" style="justify-content:center">Explorer response</div><h2>Not found</h2><p>{{.Message}}</p><a href="/">Return home</a></div></main>` + footSnip + `</body></html>{{end}}
`
