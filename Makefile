OPTS?=GO111MODULE=on
TEST_OPTS?=-race -tags no_ci -cover -timeout=5m

dep: ## Sorts dependencies
	${OPTS} go mod vendor -v

build: dep bin ## Install dependencies, build binary. `go build` with ${OPTS} 

bin: ## build skywire-peering-daemon [`spd`]
	${OPTS} go build -o ./skywire-peering-daemon ./cmd/daemon

install: ## install `skywire-peering-daemon`
	${OPTS} go install ./cmd/daemon

test: ## run tests
	- go clean -testcache
	go test ${TEST_OPTS} ./src/...
