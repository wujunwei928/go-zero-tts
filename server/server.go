package main

import (
	"flag"
	"fmt"
	"server/internal/config"
	"server/internal/handler"
	"server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/server-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 使用自定义cors中间件
	//server.Use(middleware.NewCorsMiddleware().Handle)
	// 使用go-zero提供的cors中间件
	opt := rest.WithCors(c.CorsOriginList...)
	opt(server)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
