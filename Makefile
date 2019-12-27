TEST_OPTS?=-race -tags no_ci -cover -timeout=5m

test: ## run tests
	- go clean -testcache
	go test ${TEST_OPTS} ./src/...
