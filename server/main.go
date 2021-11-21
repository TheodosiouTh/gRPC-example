package main

import (
	"context"
	"fmt"
	"gRPC-example/item"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	errorMessage := initializeServer()
	if errorMessage != "" {
		log.Fatalf(errorMessage)
	}
}

type itemsServer struct {
	item.UnimplementedItemsServer
}

const PORT = ":8888"

func initializeServer() string {
	server := grpc.NewServer()

	var items itemsServer
	item.RegisterItemsServer(server, items)

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		return fmt.Sprintf("could not listen to port %s: %v", PORT, err)
	}
	log.Printf("listening to port %s", PORT)

	err = server.Serve(listener)
	if err != nil {
		return fmt.Sprintf("Could not serve: %v", err)
	}
	return ""
}

func (server itemsServer) Add(context.Context, *item.Item) (*item.Item, error) {
	fmt.Println("Add")
	return &item.Item{}, nil
}

func (server itemsServer) List(context.Context, *item.Void) (*item.ItemList, error) {
	fmt.Println("List")
	return &item.ItemList{}, nil
}

func (server itemsServer) Find(context.Context, *item.ItemId) (*item.Item, error) {
	fmt.Println("Find")
	return &item.Item{}, nil
}

func (server itemsServer) Delete(context.Context, *item.ItemId) (*item.ItemList, error) {
	fmt.Println("Delete")
	return &item.ItemList{}, nil

}
