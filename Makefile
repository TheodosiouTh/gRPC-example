protoc:
	protoc --go_out=. --go-grpc_out=. todo.proto
start:
	docker-compose up --build
stop: 
	docker-compose down --rmi all --volumes 
install: 
	go install 