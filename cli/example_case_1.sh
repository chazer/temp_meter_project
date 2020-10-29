#!/usr/bin/env sh

# shellcheck source=.init.sh
. "$(dirname "$0")/.init.sh"

create_token() {
  local email="$1"
  local name="$2"
  local JSON; JSON="$( api_get_token "$email" "$name" )" || return 1
  local TOKEN; TOKEN="$( json_extract "$JSON" "token" )" || return 1
  [ -n "$TOKEN" ] || return 1
  echo "$TOKEN"
}

create_device() {
  local email="$1"
  local name="$2"
  local JSON; JSON="$( api_create_device "$email" "$name")" || return 1
  local UUID; UUID="$( json_extract "$JSON" "device.uuid" )" || return 1
  [ -n "$UUID" ] || return 1
  echo "$UUID"
}


EMAIL="$( gen_random_email 10 )"
NAME="$( gen_random_user_name 10 )"
DNAME="$( gen_random_device_name 10 )"

TOKEN="$( create_token "$EMAIL" "$NAME")" || { echo "Error create token"; exit 1; }
echo "USER NAME = ${NAME}"
echo "USER EMAIL = ${EMAIL}"
echo "USER AUTH TOKEN = ${TOKEN}"

UUID="$( create_device "$EMAIL" "$DNAME")" || { echo "Error create device"; exit 1; }
echo "DEVICE NAME = ${DNAME}"
echo "DEVICE UUID = ${UUID}"

DEVICE_JSON="$( api_get_device "$UUID" )" || { echo "Error get device by uuid"; exit 1; }
echo "Device = ${DEVICE_JSON}"

DEVICES_JSON="$( api_get_my_devices "$TOKEN" )" || { echo Error; exit 1; }
echo "User devices = ${DEVICES_JSON}"
