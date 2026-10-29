GO := go
GOFLAGS := GOEXPERIMENT=simd

run:
	$(GOFLAGS) $(GO) run ./cmd/main/main.go

build:
	$(GOFLAGS) $(GO) build ./cmd/main/main.go
