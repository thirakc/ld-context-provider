package main

import (
	"context"
	"ld-context-provider/pkg/config"
	"ld-context-provider/pkg/controller/resthandler/ldcontext"
	ldcontextsvc "ld-context-provider/pkg/controller/service/ldcontext"
	"ld-context-provider/pkg/httpserver"
	"ld-context-provider/pkg/logz"
	ldcontextstore "ld-context-provider/pkg/store/ldcontext"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "time/tzdata"
)

var cfg config.Config

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
}

func main() {
	logger := logz.NewLogger()

	handlers := make([]httpserver.HTTPHandler, 0)

	ldContextStore := ldcontextstore.NewLDContextStore(cfg.MongodbUri, "ldcontext", "c")
	defer func() {
		if err := ldContextStore.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	LDContextHandlers := ldcontext.NewHandler(ldcontextsvc.NewService(ldContextStore))

	handlers = append(handlers, LDContextHandlers.GetRESTHandlers()...)

	restHTTP := httpserver.New(gin.DebugMode, handlers, httpserver.WithCustomPort(":"+cfg.ServerPort))
	if err := restHTTP.Start(); err != nil {
		logger.Panic(err.Error())
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
}
