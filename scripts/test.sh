#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

go test -race -covermode=atomic ./...
export TF_ACC=1
go test -v ./graylog/resource/... -covermode=atomic
go test -v ./graylog/datasource/... -covermode=atomic
