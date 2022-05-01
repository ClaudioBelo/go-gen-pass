run:
	@go run cmd/cli/main.go 30

install-linux:
	@go build -o dist/goGenPass cmd/cli/main.go
	@sudo cp dist/goGenPass /usr/local/bin/