package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	TemplateDir = "public"
)

var (
	//go:embed static
	Res embed.FS
)

func main() {
	port := flag.String("port", "8080", "default port is 8080")
	flag.Parse()
	addr := fmt.Sprintf(":%v", *port)
	_ = addr

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		tpl, err := template.ParseFS(Res, "static/index.html")
		if err != nil {
			c.Writer.Header().Set("Content-Type", "text/html")
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.Writer.Write([]byte(err.Error()))
			return
		}

		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusOK)
		if err := tpl.Execute(c.Writer, map[string]interface{}{
			// "email": "myemail@gmail.com",
		}); err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.Writer.Write([]byte(err.Error()))
			return
		}
	})
	router.StaticFile("/wasm_exec.js", "cmd/frontend/static/wasm_exec.js")
	router.StaticFile("/main.wasm", "cmd/frontend/static/main.wasm")

	server := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("Server is listening on port :%v\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
