package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handler "github.com/rikisan1993/go-nic-jackson-microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "dang", log.LstdFlags|log.Lshortfile)
	hello := handler.NewHello(logger)
	goodbye := handler.NewGoodbye(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", hello)
	serveMux.Handle("/goodbye", goodbye)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  50 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Printf("Receive %s, gracefully shutdown", sig)

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(context)
}
