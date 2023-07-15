package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/scriptdealer/clean-todo/services"
	"github.com/scriptdealer/clean-todo/storage"
	"github.com/scriptdealer/clean-todo/transport/rest"
)

func main() {
	//InitConfig
	//InitLogger

	//Injections
	db := storage.NewMemoryStorage()
	services := services.NewComposer(db)
	handlers := rest.InitHandlers(services)
	services.Server = &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        handlers,
		ReadTimeout:    14 * time.Second,
		WriteTimeout:   14 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Info: Starting web/http server on %s...\n", services.Server.Addr)
	// setup signal catching
	signal.Notify(services.Interruption, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-services.Interruption
		fmt.Println("RECEIVED SIGNAL:", s)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//shutdown the server
		err := services.Server.Shutdown(ctx)
		if err == nil {
			os.Exit(0)
		} else {
			fmt.Printf("Graceful shutdown error: %v\n", err)
			services.Server.Close()
		}
	}()
	servingError := services.Server.ListenAndServe()
	fmt.Println("Info:", servingError.Error())
}
