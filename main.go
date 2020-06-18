package main

import (
	"log"
	"net/http"
	"os"
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

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
