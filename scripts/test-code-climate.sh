#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

repo_name=${1:-}
if [ -z "$repo_name" ]; then
  echo "the repository name is required" >&2
  exit 1
fi

mkdir -p bin
export PATH=$PWD/bin:$PATH
ee curl -L -o bin/cc-test-reporter https://codeclimate.com/downloads/test-reporter/test-reporter-0.6.3-linux-amd64
ee chmod a+x bin/cc-test-reporter

ee cc-test-reporter before-build

ee mkdir -p .code-climate

export TF_ACC=1
for d in $(go list ./...); do
  echo "$d"
  profile=.code-climate/$d/profile.txt
  coverage=.code-climate/$d/coverage.json
  ee mkdir -p "$(dirname "$profile")" "$(dirname "$coverage")"
  ee go test -v -race -coverprofile="$profile" -covermode=atomic "$d"
  if [ "$(wc -l < "$profile")" -eq 1 ]; then
    continue
  fi
  ee cc-test-reporter format-coverage -t gocov -p "github.com/${repo_name}" -o "$coverage" "$profile"
done

result=.code-climate/codeclimate.total.json
# shellcheck disable=SC2046
ee cc-test-reporter sum-coverage $(find .code-climate -name coverage.json) -o "$result"
ee cc-test-reporter upload-coverage -i "$result"
