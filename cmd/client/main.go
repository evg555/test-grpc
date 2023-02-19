package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	pb "grpc-server/pkg/api"
	"log"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatalln("not enough arguments")
	}

	name := flag.Arg(0)

	age, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserClient(conn)
	resp, err := client.Create(context.Background(), &pb.UserRequest{
		Name: name,
		Age:  uint32(age),
	})
	if err != nil {
		return
	}

	log.Println(resp.GetResult())
}
