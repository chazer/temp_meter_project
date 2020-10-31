
api_request() {
  set -o pipefail
  local method="$1"
  local uri="$2"
  local body="$3"
  if $DEBUG; then
    (
      (
        (
          (
            curl -X "$method" "$uri" -d "$body" -L -v -s --fail \
            -D /dev/fd/6 \
            --trace-ascii /dev/fd/3 \
            2>/dev/null | tee /dev/fd/6 >&8
          ) 3>&1 | sed -n '/=>/,/<=/p' | grep -v '^\(<=\|==\|=>\)' | sed 's/^[0-9a-f]*: //g' | sed -e 's/^/<= /' >&7
        ) 6>&1 | sed -e 's/^/=> /' >&7
      ) 7>&1 | sort -k1.1,1.1 -s | sed -e 's/^\(.*\)$/'$'\033[0;90m''\1'$'\033[0m''/' >&9
    ) 8>&1 9>&2
  else
    curl -X "$method" "$uri" -d "$body" -L -v -s --fail \
    2>/dev/null
  fi
}

#api_create_device() {
#  local email="$1"
#  local name="$2"
#  api_request POST "${API_URI}/devices" '{"device_name":"'${name}'","for_email":"'${email}'"}'
#}

api_get_user_token() {
  local email="$1"
  local name="$2"
  api_request POST "${API_URI}/auth/user/token" '{"user_email":"'${email}'"}'
}

api_get_device_token() {
  local name="$1"
  local email="$2"
  api_request POST "${API_URI}/auth/device/token" '{"device_name":"'${name}'","user_email":"'${email}'"}'
}

api_get_device() {
  local uuid="$1"
  local token="$2"
  api_request GET "${API_URI}/devices/byId?id=${uuid}&token=${token}"
}

api_get_device_log() {
  local uuid="$1"
  local token="$2"
  api_request GET "${API_URI}/devices/byId/log?id=${uuid}&token=${token}"
}

api_get_my_devices() {
  local token="$1"
  api_request GET "${API_URI}/devices/?token=${token}"
}

api_device_save_measurement() {
  local token="$1"
  local value="$2"
  local time_ms="$(time_ms)"
  api_request POST "${API_URI}/measurements/temp?token=${token}" '[{"time":'"${time_ms}"', "value":'"${value}"'}]'
}
