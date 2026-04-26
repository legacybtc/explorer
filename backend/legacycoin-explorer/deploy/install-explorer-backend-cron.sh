#!/bin/bash
(crontab -l 2>/dev/null | grep -v '/home/maxgor/legacycoin-explorer-service/run-explorer-backend.sh'; echo '@reboot /home/maxgor/legacycoin-explorer-service/run-explorer-backend.sh >/home/maxgor/legacycoin-explorer-service/explorer.out.log 2>/home/maxgor/legacycoin-explorer-service/explorer.err.log </dev/null &') | crontab -
crontab -l | tail -n 6
