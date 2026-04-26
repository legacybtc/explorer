# LegacyCoin Explorer

Public explorer repository for LegacyCoin.

What is here:
- `index.html`
  GitHub Pages landing page served at `https://legacybtc.github.io/explorer/`
- `backend/legacycoin-explorer`
  Full Go explorer source code

Current structure:
- static public landing page at the repo root
- real Go explorer backend source under `backend/legacycoin-explorer`
- deployment helper scripts for the backend under `backend/legacycoin-explorer/deploy`

Notes:
- the GitHub Pages site is static and reads live data from a separate backend API
- the full explorer backend is deployed on the Debian server and exposed through HTTPS
- temporary local helper scripts used during deployment are ignored from version control
- the backend now supports block, transaction ID, and wallet address search

Transparency:
- the hosted landing page source is in this repo
- the real Go backend source is in this repo
- deployment helper scripts that matter to the backend are in this repo
