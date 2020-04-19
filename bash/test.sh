#!/usr/bin/env bash

set -eu

run_test_and_write_report() {
    echo "" > coverage.txt
    for d in $(go list ./... | grep -v vendor); do
        go test -race -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    done
}

upload_codecov_report() {
    bash <(curl -s https://codecov.io/bash)
}

main() {
    run_test_and_write_report
    upload_codecov_report
}

main "$@" 
