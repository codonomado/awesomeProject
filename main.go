package main

import (
	"awesomeProject/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "stuff-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.ListenAndServe(":9090", nil)
}