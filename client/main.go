package main

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "go-grpc-client/pb"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func greet(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name is required"})
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Greet(ctx, &pb.GreetRequest{Greeting: name})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"result": r.GetResult()})
}

func main() {
	e := echo.New()
	e.GET("/greet/:name", greet)

	log.Println("Client is running on port :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
