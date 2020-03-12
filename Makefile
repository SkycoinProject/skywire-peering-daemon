OPTS?=GO111MODULE=on
TEST_OPTS?=-race -tags no_ci -cover -timeout=5m

bin: ## build skywire-peering-daemon [`spd`]
	${OPTS} go build -o ./skywire-peering-daemon ./cmd/daemon

build: dep bin ## Install dependencies, build binary. `go build` with ${OPTS}

check: lint test ## Run linters and tests

clean: ## Clean project: remove created binaries and apps
	-rm -f ./skywire-peering-daemon

dep: ## Sorts dependencies
	${OPTS} go mod vendor -v

format: ## Formats the code. Must have goimports installed (use make install-linters).
	${OPTS} goimports -w -local github.com/SkycoinProject/skywire-peering-daemon ./pkg
	${OPTS} goimports -w -local github.com/SkycoinProject/skywire-peering-daemon ./cmd

install: ## install `skywire-peering-daemon`
	${OPTS} go install ./cmd/daemon

install-linters: ## Install linters
	- VERSION=1.23.8 ./ci_scripts/install-golangci-lint.sh
	# GO111MODULE=off go get -u github.com/FiloSottile/vendorcheck
	# For some reason this install method is not recommended, see https://github.com/golangci/golangci-lint#install
	# However, they suggest `curl ... | bash` which we should not do
	# ${OPTS} go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	${OPTS} go get -u golang.org/x/tools/cmd/goimports

lint: ## Run linters. Use make install-linters first
	${OPTS} golangci-lint run -c .golangci.yml ./...
	# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	${OPTS} go vet -all ./...

test: ## run tests
	- go clean -testcache
	go test ${TEST_OPTS} ./pkg/...
