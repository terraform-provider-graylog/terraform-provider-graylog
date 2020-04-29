#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

set -eu

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

echo "" > coverage.txt

export TF_ACC=1
for d in $(go list ./...); do
  echo "$d"
  ee go test -v -race -coverprofile=profile.out -covermode=atomic "$d"
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done
