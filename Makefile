test:
	@go test -race ./...

bench:
	@go test -run=^$$ -bench=. ./...

fmt-check:
	@gofmt -l .

fmt-write:
	@gofmt -l -w .

lint:
	@golangci-lint run

# calculate % of code covered by tests and display covered/uncovered code in a browser
coverage:
	@go test -coverprofile=c.out ./...
	@go tool cover -html=c.out
	@rm c.out
