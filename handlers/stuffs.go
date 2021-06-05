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

	err := ls.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

	w.Write(d)
}
