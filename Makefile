GO := go
GOFLAGS := GOEXPERIMENT=simd
BENCH ?= .

run:
	$(GOFLAGS) $(GO) run ./cmd/main/main.go

build:
	$(GOFLAGS) $(GO) build ./cmd/main/main.go

test:
	$(GOFLAGS) $(GO) test -v ./...

benchmark:
	$(GOFLAGS) $(GO) test ./matrix/... -bench='$(BENCH)' -benchmem
