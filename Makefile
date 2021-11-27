protoc:
	protoc --go_out=. --go-grpc_out=. todo.proto
run:
	docker-compose up --build
stop: 
	docker-compose down
install: 
	go install 