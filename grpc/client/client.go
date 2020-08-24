package main

import (
	context "context"
	"io"
	"log"

	grpc "google.golang.org/grpc"
)

type HelloServiceClientImpl struct{}

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "client 1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.GetValue())

	// stream
	stream, err := client.HiStream(context.Background())
	if err != nil {
		log.Println(err)
	}

	go func() {
		if err := stream.Send(&String{Value: "client 1"}); err != nil {
			log.Println(err)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				log.Println(err)
				return
			}

			log.Println(reply.GetValue())
		}
	}()

	select {}
}
