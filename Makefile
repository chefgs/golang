BINARY_NAME=main.out

include src/api
include src/apiauth
include src/apicopy
include src/helloworld
include src/hiworld
 
all: build test
 
build:
    go build -o ${BINARY_NAME} main.go
 
test:
    go test -v main.go
 
run:
    go build -o ${BINARY_NAME} main.go
    ./${BINARY_NAME}
 
clean:
    go clean
    rm ${BINARY_NAME}
