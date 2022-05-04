package main

import (
	"log"
	"net/http"
)

const (
	AddSrv      = ":8080"
	TemplateDir = "/home/zero/www/check/go-wasm/public"
)

func main() {
	fileSrv := http.FileServer(http.Dir(TemplateDir))

	log.Println("Server is listening on port :8080")

	if err := http.ListenAndServe(AddSrv, fileSrv); err != nil {
		panic(err.Error())
	}
}
