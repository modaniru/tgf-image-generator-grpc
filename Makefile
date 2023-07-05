.PHONY: run
run: fmt
	go run ./src/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.DEFAULT := run