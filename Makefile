BINARY_NAME=grcflshtdwn

build:
	go build -o ${BINARY_NAME}

run:
	go run main.go

lint:
	revive -formatter friendly  ./...
	golangci-lint  run ./...

test:
	go test -v ./... -cover

clean:
	go clean

