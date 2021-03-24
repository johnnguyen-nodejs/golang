package handlers

import (
	"log"
	"io/ioutil"
	"net/http"
	"fmt"
)
type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye  {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	g.l.Println("Goodbye World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Goodbye %s", d)
}