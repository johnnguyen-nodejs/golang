package handlers

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)
type Hello struct {
	l *log.Logger
}
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Hello %s", d)
}