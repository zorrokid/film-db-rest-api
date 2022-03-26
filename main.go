package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/zorrokid/film-db-rest-api/data"
	"github.com/zorrokid/film-db-rest-api/handlers"
)

func main() {
	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()
	r := mux.NewRouter()

	logger := log.New(os.Stdout, "movies-api ", log.LstdFlags)
	repository := data.NewMoviesTestDataRepository(logger)
	moviesHandler := handlers.NewMovies(logger, repository)

	r.HandleFunc("/movies", moviesHandler.GetMovies)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	logger.Println("Starting server")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until signal received
	<-ch

	// deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	logger.Println("Shutting down")
	os.Exit(0)
}
