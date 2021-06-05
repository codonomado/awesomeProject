package handlers

import (
	"github.com/codonomado/awesomeProject/data"
	"log"
	"net/http"
)

type Stuffs struct{
	l *log.Logger
}

func NewStuffs(l *log.Logger) *Stuffs {
	return &Stuffs{l}
}

func (s *Stuffs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ls := data.GetStuffs()
}
