#!/bin/sh
set -e
while true; do
	node="node$(shuf -i 1-3 -n 1)"
	echo "will stop $node..."
	docker-compose stop "$node"
	s="$(shuf -i 20-50 -n 1)"
	echo "sleeping ${s}s"
	sleep "$s"
	docker-compose start "$node"
	s="$(shuf -i 1-20 -n 1)"
	echo "sleeping ${s}s"
	sleep "$s"
done
