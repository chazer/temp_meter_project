
json_extract() {
  echo "$1" |  jq -r -M ."$2"
}

gen_random_email() {
  local SIZE="$1"
  SIZE="${SIZE:-10}"
  local R=$(( ( RANDOM % SIZE )  + 1 ))
  echo "$R@email"
}

gen_random_user_name() {
  local SIZE="$1"
  SIZE="${SIZE:-1000}"
  local R=$(( ( RANDOM % SIZE )  + 1 ))
  echo "User-$R"
}

gen_random_device_name() {
  local SIZE="$1"
  SIZE="${SIZE:-1000}"
  local R=$(( ( RANDOM % SIZE )  + 1 ))
  echo "Device$R"
}
