#!/bin/bash
cd /home/maxgor/legacycoin-explorer-service
exec /home/maxgor/legacycoin-explorer-service/explorer \
  -nodehost=127.0.0.1 \
  -nodeport=19556 \
  -rpcuser=legacycoin \
  -rpcpassword=legacycoin123 \
  -port=8088
