build:
	@go build -o bin/gohon

run:
	@./bin/gohon

test:
	@go test -v ./...
