package main

import (
	"context"
	"fmt"
	"log"
	"time"

	todo "github.com/bagastri07/Todolist-gRPC-Gateway/protobuf/go"
	"google.golang.org/grpc"
)

const todoServiceAddress = "localhost:7000"

func main() {
	//create Connection to gRPC server
	conn, err := grpc.Dial(todoServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	//create New Service
	todoServiceClient := todo.NewToDoServiceClient(conn)

	//Create connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//CreateData
	todoData := todo.ToDo{
		Title:       "Masak",
		Description: "Nyalakan Kompor",
	}

	res, err := todoServiceClient.Create(ctx, &todo.CreateRequest{
		ToDo: &todoData,
	})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(res)

	//UpdateData
	todoDataUpdate := todo.ToDo{
		Id:          1,
		Title:       "Minum",
		Description: "Nyalakan Kompor",
		Completed:   1,
	}

	res2, err := todoServiceClient.Update(ctx, &todo.UpdateRequest{
		ToDo: &todoDataUpdate,
	})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(res2)

	//Read
	res3, err := todoServiceClient.Read(ctx, &todo.ReadRequest{
		Id: 2,
	})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(res3)

	//delete
	res4, err := todoServiceClient.Delete(ctx, &todo.DeleteRequest{
		Id: 2,
	})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(res4)

	//ReadAllData
	res5, err := todoServiceClient.ReadAll(ctx, &todo.ReadAllRequest{})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(res5)
}
