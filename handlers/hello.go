package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is struct for a Hello Handler
type Hello struct {
	logger *log.Logger
}

// NewHello creates a new Hello Handler
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

// ServeHTTP serves an HTTP request
func (hello *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hello.logger.Println("Hello World")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", data)
}
