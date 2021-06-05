package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h*Hello) ServeHttp(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "OPops", http.StatusBadRequest)
		return
	}

	//log.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Hello %s", d)
}