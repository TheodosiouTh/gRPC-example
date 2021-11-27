package main

import (
	"fmt"
	"log"
	"net"
	"todo/server/db"
	"todo/todo"

	"google.golang.org/grpc"
)

const PORT = ":8080"

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("Could not connectto the Database: %v", err)
	}

	err = initializeServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

}

type todoServer struct {
	todo.UnimplementedTodoServer
}

func initializeServer() error {

	server := grpc.NewServer()

	var todoService todoServer
	todo.RegisterTodoServer(server, todoService)

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		return fmt.Errorf("could not listen to port %s: %v", PORT, err)
	}
	log.Printf("listening to port %s", PORT)

	err = server.Serve(listener)
	if err != nil {
		return fmt.Errorf("could not serve: %v", err)
	}
	return nil
}
