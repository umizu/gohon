build:
	@go build -o bin/gohon

run: build
	@./bin/gohon

test:
	@go test -v ./...
