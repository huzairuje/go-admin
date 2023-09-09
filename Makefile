run:
	go run main.go -env=local

test:
	go test -race -short ./...