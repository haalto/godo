package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/haalto/godo/data"
	"github.com/haalto/godo/handlers"
	"github.com/haalto/godo/models"
)

func main() {
	db := data.GetDB()
	db.AutoMigrate(&models.Todo{})

	logger := log.New(os.Stdout, "godo-api", log.LstdFlags)
	todoHandler := handlers.NewTodos(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/todos", todoHandler)

	server := &http.Server{
		Addr:    ":8081",
		Handler: serveMux,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}

	}()

	logger.Println("Server started")

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	logger.Printf("Received terminate, terminating: %s", sig)

	tc, _ := context.WithTimeout(context.Background(), 5*time.Second)

	server.Shutdown(tc)
}
