#!/bin/sh

NAME=$(basename "$1")
CACHE_DIR="./cache"
mkdir -p "$CACHE_DIR"
CACHE="$CACHE_DIR/$NAME"

if [ -f "$CACHE" ]; then
	echo "fetching $NAME from cache" >&2
	jq <"$CACHE"
	exit
fi

echo "caching $NAME" >&2
curl 'https://www.lcbo.com/graphql' \
	-H 'content-type: application/json' \
	--compressed \
	--data-raw "{\"query\":\"$(sed 's/"/\\"/g' "$1" | tr -d '\n')\"}" |
	tee "$CACHE" | jq
