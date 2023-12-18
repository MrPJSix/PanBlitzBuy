package main

import (
	"flag"
	"fmt"

	"pan-blitz-buy/rpc/productevent/internal/config"
	eventserviceServer "pan-blitz-buy/rpc/productevent/internal/server/eventservice"
	productserviceServer "pan-blitz-buy/rpc/productevent/internal/server/productservice"
	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/productevent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		productevent.RegisterProductServiceServer(grpcServer, productserviceServer.NewProductServiceServer(ctx))
		productevent.RegisterEventServiceServer(grpcServer, eventserviceServer.NewEventServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
