MAIN_PACKAGE_PATH := ./cmd/todo
BINARY_NAME := todo

.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME}


.PROXY: generate-structs
generate-structs:
	mkdir -p internal/todo_v1
	protoc --go_out=internal/todo_v1 --go_opt=paths=source_relative \
			proto/todo_v1/service.proto

.PROXY: generate
generate:
	mkdir -p internal/todo_v1
	protoc --go_out=internal/todo_v1 --go_opt=paths=import \
			--go-grpc_out=internal/todo_v1 --go-grpc_opt=paths=import \
			proto/todo_v1/service.proto
	mv internal/todo_v1/ToDoAppGrpc/internal/todo_v1/* internal/todo_v1/
	rm -rf internal/todo_v1/ToDoAppGrpc