package handlers

import (
	"log"
	"net/http"
)

type Goodie struct {
	l *log.Logger
}


func NewGoodie(l *log.Logger) *Goodie {
	return &Goodie{l}
}

func (g *Goodie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Byee"))
}