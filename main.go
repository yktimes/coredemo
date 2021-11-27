package main

import (
	"context"
	"coredemo/framework"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()

	// 当前g等待信号量
	quit := make(chan os.Signal)
	// 监控信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
