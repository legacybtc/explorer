#!/bin/bash
pkill -f 'cloudflared tunnel --url http://127.0.0.1:8088' || true
setsid /home/maxgor/.local/bin/cloudflared tunnel --url http://127.0.0.1:8088 --no-autoupdate >/home/maxgor/cloudflared/explorer-tunnel.log 2>&1 < /dev/null &
sleep 10
grep -Eo 'https://[-a-z0-9]+\.trycloudflare\.com' /home/maxgor/cloudflared/explorer-tunnel.log | head -n 1
ps -ef | grep 'cloudflared tunnel --url http://127.0.0.1:8088' | grep -v grep || true
