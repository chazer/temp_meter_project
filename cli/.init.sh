
DEBUG=false

API_URI="${API_URI:-http://127.0.0.1:8080}"

while [ $# -gt 0 ]; do
  key="$1"
  case $key in
    -d|--debug)
      DEBUG=true
      shift
      ;;
    -s|--server)
      API_URI="http://$2"
      shift
      shift
      ;;
    *) # unknown option
      shift
      ;;
  esac
done


command -v curl >/dev/null || { echo "The curl binary is not found"; exit 1; }
command -v jq >/dev/null || { echo "The jq binary is not found"; exit 1; }

export DEBUG
export API_URI

# shellcheck source=.api.sh
. "$(dirname "$0")/.api.sh"

# shellcheck source=.api.sh
. "$(dirname "$0")/.functions.sh"
