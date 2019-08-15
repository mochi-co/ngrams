package main

import (
	"encoding/json"
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
		errHandler(w, 500, err)
	}

	if len(b) == 0 {
		errHandler(w, 400, err)
	}

	tokens, err := index.Parse(string(b))
	if err != nil {
		errHandler(w, 500, err)
	}

	m, err := json.Marshal(map[string]interface{}{
		"total_tokens": len(tokens),
	})
	if err != nil {
		errHandler(w, 500, err)
	}

	w.Write(m)

}

// generateHandler is a GET request handler that generates a string of random
// text in the syntactic style of the trained ngrams.
func generateHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// Default length for the body to generate.
	tokenLen := 50

	if r.URL.Query().Get("limit") != "" {
		tokenLen, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			errHandler(w, 500, err)
		}
	}

	out, err := index.Babble("", tokenLen)
	if err != nil {
		if err == ngrams.ErrEmptyIndex {
			m, err := json.Marshal(map[string]interface{}{
				"err": "index is empty; please learn ngrams before generating.",
			})
			if err != nil {
				errHandler(w, 400, err)
			}

			w.Write(m)
			return
		}

		errHandler(w, 500, err)
	}

	m, err := json.Marshal(map[string]interface{}{
		"body":  out,
		"limit": tokenLen,
	})
	if err != nil {
		errHandler(w, 500, err)
	}

	w.Write(m)

}

// errHandler is a convenience function which writes and logs errors.
func errHandler(w http.ResponseWriter, code int, err error) {
	log.Println("Error:", err)
	http.Error(w, http.StatusText(code), code)
}
