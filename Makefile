PWD := $(shell pwd 2>/dev/null)
PKG_LIST := $(shell go list ./... | grep -v /utils/shm  2>/dev/null)

.PHONY: coverage
coverage:
	  sh ./scripts/coverage.sh