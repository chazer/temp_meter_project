
api_request() {
  local method="$1"
  local uri="$2"
  local body="$3"
  if $DEBUG; then
    (
      (
        (
          (
            curl -X "$method" "$uri" -d "$body" -L -v -s \
            -D /dev/fd/6 \
            --trace-ascii /dev/fd/3 \
            2>/dev/null | tee /dev/fd/6 >&8
          ) 3>&1 | sed -n '/=>/,/<=/p' | grep -v '^\(<=\|==\|=>\)' | sed 's/^[0-9a-f]*: //g' | sed -e 's/^/<= /' >&7
        ) 6>&1 | sed -e 's/^/=> /' >&7
      ) 7>&1 | sort -k1.1,1.1 -s | sed -e 's/^\(.*\)$/'$'\033[0;90m''\1'$'\033[0m''/' >&9
    ) 8>&1 9>&2
  else
    curl -X "$method" "$uri" -d "$body" -L -v -s \
    2>/dev/null
  fi
}

api_create_device() {
  local email="$1"
  local name="$2"
  api_request POST "${API_URI}/devices" '{"device_name":"'${name}'","for_email":"'${email}'"}'
}

api_get_token() {
  local email="$1"
  local name="$2"
  api_request POST "${API_URI}/auth/token" '{"user_name":"'${name}'","user_email":"'${email}'"}'
}

api_get_device() {
  local uuid="$1"
  api_request GET "${API_URI}/devices/byId?id=${uuid}"
}

api_get_my_devices() {
  local token="$1"
  api_request GET "${API_URI}/devices/?token=${token}"
}
