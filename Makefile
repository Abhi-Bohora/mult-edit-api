BINARY_NAME=main.out

build:
	mkdir -p bin
	go build -o bin/${BINARY_NAME} cmd/server/main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}