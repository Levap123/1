package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"a/funcs"
)

func main() {
	// Parsing JSON
	if err := funcs.ParseJson(); err != nil {
		log.Fatalf("unable to start server, groups information unavailable")
	}

	mux := http.NewServeMux()

	// Serving static files
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))

	// Handlers
	mux.HandleFunc("/", funcs.IndexHandler)
	mux.HandleFunc("/group/", funcs.GroupHandler)
	mux.HandleFunc("/search", funcs.SearchHandler)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server is listening on http://localhost:8080")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("fatal is starting server: %v\n", err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("server ...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("fatal in server shutdown: %v\n", err)
	}
}
