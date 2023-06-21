build:
	@go build -o /bin/basic-api

run:
	@./bin/basic-api

test:
	@go test -v ./...
