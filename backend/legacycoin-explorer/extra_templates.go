package explorer

const extraTemplates = `
{{define "tx"}}` + headSnip + `<title>Transaction {{.Tx.TxID}} вЂ” LegacyCoin Explorer</title></head><body>` + navSnip + `<main class="page">
<section class="block-hero">
  <a class="back-link" href="/block/{{.Tx.BlockHash}}">в†ђ Back to containing block</a>
  <div class="eyebrow" style="margin-top:20px">Transaction detail</div>
  <div class="block-title">Transaction</div>
  <div class="block-sub">{{.Tx.TxID}}</div>
  <div class="block-summary">
    <div class="summary-pill"><span>Block height</span><strong>{{.Tx.BlockHeight}}</strong></div>
    <div class="summary-pill"><span>Confirmations</span><strong>{{.Tx.Confirmations}}</strong></div>
    <div class="summary-pill"><span>Block position</span><strong>#{{.Tx.Index}}</strong></div>
    <div class="summary-pill"><span>Block tx count</span><strong>{{.Tx.BlockTxCount}}</strong></div>
  </div>
</section>
<section class="detail-grid">
  <div class="detail-card">
    <h3>Chain placement</h3>
    <div class="detail-row"><div class="detail-key">Transaction ID</div><div class="detail-val">{{.Tx.TxID}}</div></div>
    <div class="detail-row"><div class="detail-key">Block hash</div><div class="detail-val"><a href="/block/{{.Tx.BlockHash}}">{{.Tx.BlockHash}}</a></div></div>
    <div class="detail-row"><div class="detail-key">Block height</div><div class="detail-val"><a href="/block/{{.Tx.BlockHeight}}">{{.Tx.BlockHeight}}</a></div></div>
    <div class="detail-row"><div class="detail-key">Time</div><div class="detail-val">{{formatTime .Tx.BlockTime}}</div></div>
  </div>
  <div class="detail-card">
    <h3>Explorer note</h3>
    <div class="detail-row"><div class="detail-key">Status</div><div class="detail-val">Confirmed on active chain</div></div>
    <div class="detail-row"><div class="detail-key">Confirmations</div><div class="detail-val">{{.Tx.Confirmations}}</div></div>
    <div class="detail-row"><div class="detail-key">Index in block</div><div class="detail-val">{{.Tx.Index}}</div></div>
    <div class="detail-row"><div class="detail-key">Lookup mode</div><div class="detail-val">Transaction placement scan</div></div>
  </div>
</section>
</main>` + footSnip + `</body></html>{{end}}

{{define "address"}}` + headSnip + `<title>Address {{.Address.Address}} вЂ” LegacyCoin Explorer</title></head><body>` + navSnip + `<main class="page">
<section class="block-hero">
  <a class="back-link" href="/">в†ђ Back to dashboard</a>
  <div class="eyebrow" style="margin-top:20px">Wallet address</div>
  <div class="block-title">Address summary</div>
  <div class="block-sub">{{.Address.Address}}</div>
  <div class="block-summary">
    <div class="summary-pill"><span>Confirmed balance</span><strong>{{printf "%.8f LBTC" .Address.Balance}}</strong></div>
    <div class="summary-pill"><span>Network</span><strong>LegacyCoin</strong></div>
    <div class="summary-pill"><span>Type</span><strong>P2PKH</strong></div>
    <div class="summary-pill"><span>Explorer mode</span><strong>RPC balance lookup</strong></div>
  </div>
</section>
<section class="detail-grid">
  <div class="detail-card">
    <h3>Address</h3>
    <div class="detail-row"><div class="detail-key">Wallet address</div><div class="detail-val">{{.Address.Address}}</div></div>
    <div class="detail-row"><div class="detail-key">Confirmed balance</div><div class="detail-val">{{printf "%.8f LBTC" .Address.Balance}}</div></div>
    <div class="detail-row"><div class="detail-key">Ticker</div><div class="detail-val">LBTC</div></div>
  </div>
  <div class="detail-card">
    <h3>Explorer note</h3>
    <div class="detail-row"><div class="detail-key">Lookup type</div><div class="detail-val">Current balance from node RPC</div></div>
    <div class="detail-row"><div class="detail-key">Source</div><div class="detail-val">getaddressbalance</div></div>
    <div class="detail-row"><div class="detail-key">Search support</div><div class="detail-val">Paste wallet addresses directly into the explorer search bar.</div></div>
  </div>
</section>
</main>` + footSnip + `</body></html>{{end}}
`
