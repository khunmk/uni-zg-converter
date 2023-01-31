package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khunmk/uni-zg-converter/cmd/app/handler"
)

var (
	//go:embed static
	Res embed.FS

	port = flag.String("port", "8080", "default port is 8080")
)

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%v", *port)

	r := gin.Default()
	h := handler.NewHandler(&handler.HConfig{
		R:   r,
		Res: Res,
	})
	h.Register()

	server := http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    1 * time.Minute,
		WriteTimeout:   1 * time.Minute,
		MaxHeaderBytes: 10 << 20, //10 mb
	}

	go func() {
		log.Printf("Server is listening on port :%v\n", *port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Panic(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	//unregister services
	h.Destroy()

	//shutdown server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Panic(err.Error())
	}
}
