package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dstgo/configure/api/configure/client"
	"github.com/dstgo/kratosx"
	"github.com/dstgo/kratosx/config"
	"github.com/dstgo/kratosx/pkg/printx"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/cron/internal/server/app"
	"github.com/dstgo/cron/internal/server/conf"
)

func main() {
	server := kratosx.New(
		kratosx.Config(client.NewFromEnv()),
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(
			kratos.AfterStart(func(ctx context.Context) error {
				kt := kratosx.MustContext(ctx)
				printx.ArtFont(fmt.Sprintf("Hello %s !", kt.Name()))
				return nil
			}),
		),
	)

	if err := server.Run(); err != nil {
		log.Println("run service fail", err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	cfg := &conf.Config{}
	// c.ScanWatch("business", func(value config.Value) {
	//	if err := value.Scan(cfg); err != nil {
	//		log.Printf("business配置变更失败：%s", err.Error())
	//	}
	// })
	// 注册服务
	app.New(cfg, hs, gs)
}
