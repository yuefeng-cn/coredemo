package main

import (
	"coredemo/framework"
	"coredemo/framework/middleware"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()
	// core.Use(
	// 	middleware.Test1(),
	// 	middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	core.Use(middleware.Timeout(3 * time.Second))

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
