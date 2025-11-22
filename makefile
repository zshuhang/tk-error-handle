.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o tk-error-handle.exe main.go