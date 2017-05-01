#!/usr/bin/env bash
set -e

function join { local IFS="$1"; shift; echo "$*"; }

echo "== Starting lint..."
gometalinter --vendored-linters --vendor --disable-all --enable=gofmt --enable=goimports --enable=gocyclo --cyclo-over=10
echo "== Finished lint."

PACKAGES=$(go list -e ./... | grep -v /vendor/)
TEST_PKGS=$(find . -name '*_test.go' -exec dirname '{}' ';' | grep -v './vendor' | grep -v './.glide/' | sort -u | uniq)

echo "== Building tests..."
go test -i -race $TEST_PKGS
echo "== Starting tests..."
gocov test $TEST_PKGS -v -race -coverpkg $(join "," $PACKAGES) 1> >(grep -v "{\"Packages"  >&1) 2> >(grep -v "warning: no packages being tested depend on" >&2)
echo "== Building strict tests..."
go test -i -tags=strict -race $TEST_PKGS
echo "== Starting strict tests..."
gocov test $TEST_PKGS -v -tags=strict -race -coverpkg $(join "," $PACKAGES) 1> >(grep -v "{\"Packages"  >&1) 2> >(grep -v "warning: no packages being tested depend on" >&2)
echo "== Finished tests."