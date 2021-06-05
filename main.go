package main

import (
	"github.com/codonomado/awesomeProject/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "stuff-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.Goodie(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodie", hh)

	http.ListenAndServe(":9090", sm)
}