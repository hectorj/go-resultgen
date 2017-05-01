#!/usr/bin/env bash
set -e
set -v

go get -u github.com/alecthomas/gometalinter github.com/axw/gocov/gocov

gometalinter --vendored-linters --install