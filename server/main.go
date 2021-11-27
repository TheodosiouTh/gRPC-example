package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"todo/server/db"
	"todo/todo"

	"google.golang.org/grpc"
)

var port string

func main() {

	dbHost := "database"
	port = os.Getenv("CONTAINER_SERVER_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	dbPort, err := strconv.Atoi(os.Getenv("CONTAINER_DATABASE_PORT"))
	if err != nil {
		log.Fatalf("Could not convert  the Database: %v", err)
	}

	err = db.Init(dbHost, dbName, dbUser, dbPassword, dbPort)
	if err != nil {
		log.Fatalf("Could not connect to the Database: %v", err)
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

	serverPort := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		return fmt.Errorf("could not listen to port %s: %v", serverPort, err)
	}
	log.Printf("listening to port %s", serverPort)

	err = server.Serve(listener)
	if err != nil {
		return fmt.Errorf("could not serve: %v", err)
	}
	return nil
}
