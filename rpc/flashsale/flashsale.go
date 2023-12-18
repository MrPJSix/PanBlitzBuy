package main

import (
	"flag"
	"fmt"

	"pan-blitz-buy/rpc/flashsale/flashsale"
	"pan-blitz-buy/rpc/flashsale/internal/config"
	"pan-blitz-buy/rpc/flashsale/internal/server"
	"pan-blitz-buy/rpc/flashsale/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/flashsale.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		flashsale.RegisterFlashsaleServiceServer(grpcServer, server.NewFlashsaleServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
