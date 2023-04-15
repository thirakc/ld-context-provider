package main

import (
	"context"
	"ld-context-provider/pkg/controller/resthandler/ldcontext"
	ldcontextsvc "ld-context-provider/pkg/controller/service/ldcontext"
	"ld-context-provider/pkg/httpserver"
	"ld-context-provider/pkg/logz"
	ldcontextstore "ld-context-provider/pkg/store/ldcontext"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main() {
	logger := logz.NewLogger()
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		logger.Fatal("You must set your 'PORT' environmental variable.")
	}
	dbUri := os.Getenv("MONGODB_URI")
	if dbUri == "" {
		logger.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	handlers := make([]httpserver.HTTPHandler, 0)

	ldContextStore := ldcontextstore.NewLDContextStore(dbUri, "ldcontext", "c")
	defer func() {
		if err := ldContextStore.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	LDContextHandlers := ldcontext.NewHandler(ldcontextsvc.NewService(ldContextStore))

	handlers = append(handlers, LDContextHandlers.GetRESTHandlers()...)

	restHTTP := httpserver.New(gin.DebugMode, ":"+serverPort, handlers...)
	if err := restHTTP.Start(); err != nil {
		logger.Panic(err.Error())
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
}
