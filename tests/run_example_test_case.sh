#/usr/bin/env bash

API_URI="http://127.0.0.1:8080"

json_extract() {
  echo "$1" |  python -c 'import json,sys;obj=json.load(sys.stdin);print obj'"$2" 2>/dev/null
}

gen_random_email() {
  R=$(( ( RANDOM % 10 )  + 1 ))
  echo "$R@email"
}

gen_random_user_name() {
  R=$(( ( RANDOM % 1000 )  + 1 ))
  echo "User-$R"
}

api_create_device() {
  curl -X POST "${API_URI}/devices" \
  -d '{"device_name":"Device 1","for_email":"'${1}'"}' 2>/dev/null
}

api_get_token() {
  local email="$1"
  local name="$2"
  curl -X POST "${API_URI}/auth/token" \
  -d '{"user_name":"'${name}'","user_email":"'${email}'"}' 2>/dev/null
}

api_get_device() {
  curl -X GET "${API_URI}/devices/byId?id=$1" 2>/dev/null
}

api_get_my_devices() {
  curl -X GET "${API_URI}/devices/?token=$1" 2>/dev/null
}

create_token() {
  local email="$1"
  local name="$2"
  local JSON="$( api_get_token "$email" "$name" )" || return 1
  local TOKEN="$( json_extract "$JSON" '["token"]' )" || return 1
  [ -n "$TOKEN" ] || return 1
  echo "$TOKEN"
}

create_uuid() {
  local JSON="$( api_create_device "$EMAIL")" || return 1
  local UUID="$( json_extract "$JSON" '["device"]["uuid"]' )" || return 1
  [ -n "$UUID" ] || return 1
  echo "$UUID"
}




EMAIL="$( gen_random_email )"
NAME="$( gen_random_user_name )"

UUID="$( create_uuid "$EMAIL")" || { echo "Error create device uuid"; exit 1; }
echo "UUID=${UUID}"

TOKEN="$( create_token "$EMAIL" "$NAME")" || { echo "Error create token"; exit 1; }
echo "TOKEN=${TOKEN}"

DEVICE_JSON="$( api_get_device "$UUID" )" || { echo "Error get device by uuid"; exit 1; }
echo "Device=${DEVICE_JSON}"

DEVICES_JSON="$( api_get_my_devices "$TOKEN" )" || { echo Error; exit 1; }
echo "Devices=${DEVICES_JSON}"
