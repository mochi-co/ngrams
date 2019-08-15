package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/mochi-co/ngrams"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jamiealquiza/envy"
)

var (
	// index is the ngrams index, which will store and tokenize ngrams.
	index *ngrams.Index
)

// main is our entrypoint for the service.
func main() {
	log.Println("Starting trigram service...")

	// Prepare signal catching to ensure the webserver and indexer
	// shuts down cleanly.
	sigs := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- errors.New("caught signal")
	}()

	// Optionally change the port the webserver listens on, for convenience.
	// In a cloud-native environment this would typically be set using environment
	// variables. We can achieve this in a single-line using one of my favourite new
	// packages; `envy`, which takes all the flags and transmutes them into env vars.
	port := flag.Int("port", 8080, "port to serve webserver on")
	envy.Parse("NGRAMS") // Expose environment variables as NGRAMS_PORT, etc.

	// Configure the Ngrams indexer to index trigrams and use the default options.
	index = ngrams.NewIndex(3, nil)

	// Setup our basic webserver; we'll use Chi here because it's a little bit
	// simpler than implementing a pure net/http design, has some convenient and
	// trusted middleware solutions, and is compatible with stdlib.
	r := chi.NewRouter()
	r.Use(middleware.Timeout(20 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Our mini REST router which points to the learn and generate endpoints.
	r.Post("/learn", learnHandler)
	r.Get("/generate", generateHandler)

	// We'll use the stdlib http to actually run the chi server so we can
	// implement our own signal catching. There's a way to do this using chi's
	// `valve` middleware, but it's excessive for this use-case.
	srv := http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: r,
	}
	go srv.ListenAndServe()
	log.Printf("Listening on localhost:%d\n", *port)

	// Wait for signals...
	err := <-done
	if err != nil {
		log.Println(err)
	}

	// Gracefully shutdown the index.
	index.Close()

}

// learnHandler is a POST request handler that takes a body of text and indexes it
// as trigrams. The body of text must be sent as raw plain-text.
func learnHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, http.StatusText(500), 500)
	}

	if len(b) == 0 {
		http.Error(w, http.StatusText(400), 400)
	}

	log.Println("learning body", string(b))

}

// generateHandler is a GET request handler that generates a string of random
// text in the syntactic style of the trained ngrams.
func generateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("generating")
}
