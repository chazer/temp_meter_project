#!/usr/bin/env sh

# shellcheck source=.init.sh
. "$(dirname "$0")/.init.sh"

create_user_token() {
  local email="$1"
  local name="$2"
  local JSON; JSON="$( api_get_user_token "$email" )" || return 1
  local TOKEN; TOKEN="$( json_extract "$JSON" "token" )" || return 1
  [ -n "$TOKEN" ] || return 1
  echo "$TOKEN"
}

get_my_devices() {
  local token="$1"
  local JSON; JSON="$( api_get_my_devices "$token" )" || return 1
  local UUIDS; UUIDS="$( json_extract "$JSON" "items[].device.uuid" )" || return 1
  echo "$UUIDS"
}



EMAIL="$( gen_random_email 3 )"

echo "Register new user:" >&2
UTOKEN="$( create_user_token "$EMAIL" )" || { echo "Error create token"; exit 1; }
echo "  user email = ${EMAIL}"
echo "  user auth token = ${UTOKEN}"
echo "" >&2


DEVICES_UUIDs="$( get_my_devices "$UTOKEN" )" || { echo Error; exit 1; }

if [ ! "$DEVICES_UUIDs" ]; then
  echo "No devices:"
else
  echo "User devices:"
  n=0
  for UUID in $DEVICES_UUIDs; do
    n=$(( n + 1))
    echo "- Device #$n:"
    echo "  UUID: $UUID"
    DEVICE_JSON="$( api_get_device "$UUID" "$UTOKEN" )" || { echo "Error get device by uuid"; exit 1; }
    DEVICE_UPDATED_AT="$( json_extract "$DEVICE_JSON" "device.updatedAt" )"
    echo "  Updated At: ${DEVICE_UPDATED_AT}"
    DEVICE_LOG_JSON="$( api_get_device_log "$UUID" "$UTOKEN" )" || { echo "Error get device log data"; exit 1; }
    echo "  History:"
    json_extract "$DEVICE_LOG_JSON" 'items[] | "\(.time) \(.temperature)"' \
    | while read -r time temp; do
      printf "  - %s\t%s\n" "$time" "$temp"
    done
  done
fi

echo "Done" >&2
