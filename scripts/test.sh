#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

go test -race -covermode=atomic ./...
export TF_ACC=1
go test -v  -covermode=atomic -race ./graylog/resource/...
go test -v  -covermode=atomic -race ./graylog/datasource/...
