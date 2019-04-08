package main

import (
	"flag"
	"log"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	guff_proto "github.com/toohtrot/guff/go/generated"
)


func main() {
	flag.Parse()
	glog.Info("Don't take any guff from these swine.")
	l, err  := net.Listen("tcp", ":0")
	if err != nil {
		glog.Fatalf("failed to listen: %q", err)
	}

	foo := &guff_proto.UnimplementedDivisionsServiceServer{}
	s := grpc.NewServer()
	guff_proto.RegisterDivisionsServiceServer(s, foo)

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("s.Serve(%#v) = %q", l, err)
	}
}
