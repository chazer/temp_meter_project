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

create_device_token() {
  local name="$1"
  local email="$2"
  local JSON; JSON="$( api_get_device_token "$name" "$email" )" || return 1
  local TOKEN; TOKEN="$( json_extract "$JSON" "token" )" || return 1
  [ -n "$TOKEN" ] || return 1
  echo "$TOKEN"
}

#create_device() {
#  local email="$1"
#  local name="$2"
#  local JSON; JSON="$( api_create_device "$email" "$name")" || return 1
#  local UUID; UUID="$( json_extract "$JSON" "device.uuid" )" || return 1
#  [ -n "$UUID" ] || return 1
#  echo "$UUID"
#}


EMAIL="$( gen_random_email 3 )"
DNAME="$( gen_random_device_name 10 )"

UTOKEN="$( create_user_token "$EMAIL" )" || { echo "Error create token"; exit 1; }
echo "USER EMAIL = ${EMAIL}"
echo "USER AUTH TOKEN = ${UTOKEN}"

DTOKEN="$( create_device_token "$DNAME" "$EMAIL")" || { echo "Error create token"; exit 1; }
echo "DEVICE NAME = ${DNAME}"
echo "DEVICE EMAIL = ${EMAIL}"
echo "DEVICE AUTH TOKEN = ${DTOKEN}"

#UUID="$( create_device "$EMAIL" "$DNAME")" || { echo "Error create device"; exit 1; }
#echo "DEVICE NAME = ${DNAME}"
#echo "DEVICE UUID = ${UUID}"


get_my_devices() {
  local token="$1"
  local JSON; JSON="$( api_get_my_devices "$token" )" || return 1
  local UUIDS; UUIDS="$( json_extract "$JSON" "items[].device.uuid" )" || return 1
  echo "$UUIDS"
}

DEVICES_UUIDs="$( get_my_devices "$UTOKEN" )" || { echo Error; exit 1; }

if [ ! "$DEVICES_UUIDs" ]; then
  echo "No devices:"
else
  echo "User devices:"
  n=0
  for UUID in $DEVICES_UUIDs; do
    n=$(( n + 1))
    echo "- Device #$n:"
    echo "  UUID: $UUID:"
    DEVICE_JSON="$( api_get_device "$UUID" )" || { echo "Error get device by uuid"; exit 1; }
    echo "  json: ${DEVICE_JSON}"
  done
fi
