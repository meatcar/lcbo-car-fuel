set -eo pipefail

export TABLES='cities stores products inventory'

CLEAR_EOL="\033[K"

status() {
    echo -ne "\r$*$CLEAR_EOL" >&2
}
