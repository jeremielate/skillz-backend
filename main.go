package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	Host string = "localhost"
	Port int    = 8080
)

func init() {
	flag.StringVar(&Host, "host", Host, "host to listen")
	flag.IntVar(&Port, "port", Port, "port to listen")
}

func address() string {
	return fmt.Sprintf("%s:%d", Host, Port)
}

func main() {
	flag.Parse()
	if Port < 1 || Host == "" {
		return
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	handle := http.ServeMux{}
	handle.HandleFunc("/", mainPage)
	addr := address()
	srv := http.Server{
		Handler: &handle,
		Addr:    addr,
	}

	go func() {
		for c := range sig {
			log.Printf("received %s, quitting...\n", c)
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	log.Printf("listening to %s\n", addr)
	OpenUrl(fmt.Sprintf("http://" + addr + "/"))
	if err := srv.ListenAndServe(); err != nil {
		close(sig)
		if err != http.ErrServerClosed {
			log.Println(err)
		}
	}
}
