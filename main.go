package main

import (
	"context"
	"github.com/codonomado/awesomeProject/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "stuff-api", log.LstdFlags)
	sh := handlers.NewStuffs(l)

	sm := http.NewServeMux()
	sm.Handle("/", sh)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 *time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if (err != nil) {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminal, graceful shutdown...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)
	s.Shutdown(tc)
}