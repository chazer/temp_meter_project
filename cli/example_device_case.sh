#!/usr/bin/env sh

# shellcheck source=.init.sh
. "$(dirname "$0")/.init.sh"

create_device_token() {
  local name="$1"
  local email="$2"
  local JSON; JSON="$( api_get_device_token "$name" "$email" )" || return 1
  local TOKEN; TOKEN="$( json_extract "$JSON" "token" )" || return 1
  [ -n "$TOKEN" ] || return 1
  echo "$TOKEN"
}


EMAIL="$( gen_random_email 3 )"
DNAME="$( gen_random_device_name 10 )"

echo "Register new device:" >&2
DTOKEN="$( create_device_token "$DNAME" "$EMAIL")" || { echo "Error create token"; exit 1; }
echo "  device name = ${DNAME}"
echo "  device email = ${EMAIL}"
echo "  device auth token = ${DTOKEN}"
echo "" >&2

echo "Send some measurements data" >&2
for _ in $(seq 10); do
  api_device_save_measurement "$DTOKEN" "$(random_float_in_range -40 40 )" || { echo Error; exit 1; }
  printf ". "
done
echo "" >&2

echo "Done" >&2
