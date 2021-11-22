protoc:
	protoc --go_out=. --go-grpc_out=. todo.proto
run-server:
	go run ./server