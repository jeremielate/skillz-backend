package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	handle := http.ServeMux{}
	handle.HandleFunc("/", mainPage)
	srv := http.Server{Handler: &handle}

	go func() {
		for c := range sig {
			log.Printf("received %s, quitting...\n", c)
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		close(sig)
		if err != http.ErrServerClosed {
			log.Println(err)
		}
	}
}
