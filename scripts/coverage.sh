#!/bin/bash

rm -rf ./testdata/coverage/*
go test -coverpkg=./... -coverprofile=./testdata/coverage/coverage.out ./...
go tool cover -html=./testdata/coverage/coverage.out -o ./testdata/coverage/coverage.html
go tool cover -func=./testdata/coverage/coverage.out -o ./testdata/coverage/coverage.txt