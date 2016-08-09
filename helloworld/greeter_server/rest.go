package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	gw "examples/helloworld/helloworld"
	"net"
	"log"
	"time"
	 google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)


type serverEcho struct{}

// Echo implements gw.Echo
func (s *serverEcho) Echo(ctx context.Context, in *gw.EchoMessage) (*gw.EchoMessage, error) {
	now := time.Now()
	date := google_protobuf1.Timestamp{
		Seconds: now.Unix(),
		Nanos: int32(now.UnixNano()),
	}

	date.String()
	return &gw.EchoMessage{Value: in.Value + " now! Only for `the date`", Time:&date}, nil
}


func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()


	//run grpc service
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gw.RegisterEchoServiceServer(s, &serverEcho{})
	go s.Serve(lis)


	//run api
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = gw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	http.ListenAndServe(":8080", mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
