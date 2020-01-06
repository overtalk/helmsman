PWD := $(shell pwd 2>/dev/null)
PKG_LIST := $(shell go list ./... | grep -v /utils/shm  2>/dev/null)

.PHONY: coverage
coverage:
	  rm -rf ./testdata/coverage/*
	  go test -coverpkg=./... -coverprofile=./testdata/coverage/coverage.out ./...
	  go tool cover -html=./testdata/coverage/coverage.out -o ./testdata/coverage/coverage.html
	  go tool cover -func=./testdata/coverage/coverage.out -o ./testdata/coverage/coverage.txt