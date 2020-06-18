package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Goodbye is a struct of Goodbye Handler
type Goodbye struct {
	logger *log.Logger
}

// NewGoodbye returns a new Goodbye Handler
func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

// ServeHTTP serve http request
func (goodbye *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	goodbye.logger.Println("Goodbye World!")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye %s", data)
}
