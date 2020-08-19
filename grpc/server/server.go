package main

import (
	context "context"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"

	grpc "google.golang.org/grpc"
)

type HelloServiceServerImpl struct {
}

func (h *HelloServiceServerImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello " + args.GetValue()}
	log.Println(reply.Value)
	return reply, nil
}

func (h *HelloServiceServerImpl) HiStream(stream HelloService_HiStreamServer) error {
	log.Println("Start stream server")
	defer log.Println("End stream server")

	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{Value: "hi stream " + args.GetValue()}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	// 启动gops
	// if err := agent.Listen(agent.Options{
	// 	Addr: "0.0.0.0:8848",
	// 	// ConfigDir:       "/home/centos/gopsconfig", // 最好使用默认
	// 	ShutdownCleanup: true}); err != nil {
	// 	log.Fatal(err)
	// }

	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		log.Println("Starting pprof")
		if err := http.ListenAndServe(ip, nil); err != nil {
			log.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	// grpc
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceServerImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start grpc server")
	grpcServer.Serve(lis)
}
