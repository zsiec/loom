run:
	go run ./cmd/loom

test:
	go test -count=1 -race ./...
