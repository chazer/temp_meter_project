
json_extract() {
  set -o pipefail
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

time_ms() {
  date +%s%N | cut -b1-13
}

random_float_in_range() {
  local a="$1"
  local b="$2"
  echo "$(( a + RANDOM % ( b - a ) )).$(( RANDOM % 999 ))"
}

function uriencode {
  jq -nr --arg v "$1" '$v|@uri'
}