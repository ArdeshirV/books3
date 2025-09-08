package main

import (
	"context"
	"log"
	"time"

	"github.com/ArdeshirV/book/go-solati/grpc01"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	const address = "localhost:50051"
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	c := grpc01.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &grpc01.HelloRequest{Name: "Ardeshir"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
