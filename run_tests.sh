#!/usr/bin/env bash

# Shellcheck test
echo "Test Shellcheck" >&2
docker run --rm -ti -v "$(pwd)":/src -w /src koalaman/shellcheck-alpine shellcheck \
  --shell=ash \
  -x \
  -P cli/ \
  cli/example_case_1.sh || { echo "Error on Shellcheck" >&2; exit 1; }
echo "Shellcheck done" >&2


SCOPE="tmetertestproject"

echo "Test containers build" >&2
docker-compose -p "$SCOPE" build || { echo "Error on build containers" >&2; exit 1; }
echo "Build done" >&2

echo "Test containers up" >&2
docker-compose -p "$SCOPE" -f docker-compose.yml up -d || { echo "Error on up compose project" >&2; exit 1; }
echo "Up done" >&2

echo "Test API" >&2
docker-compose -p "$SCOPE" run cli example_case_1.sh || { echo "Test 'example_case_1' failed" >&2; exit 1; }
echo "API test done" >&2

echo "Test containers down" >&2
docker-compose -p "$SCOPE" -f docker-compose.yml down -v || { echo "Error on down compose project" >&2; exit 1; }
echo "Down done" >&2
