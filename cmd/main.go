package main

// For readability, I like to separate my stdlib, internal (not pictured), and
// third-party imports with newlines.
import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jamiealquiza/envy"

	"github.com/mochi-co/trigrams-test"
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
	envy.Parse("TRIGRAMS") // Expose environment variables as TRIGRAMS_PORT, etc.

	log.Printf("Listening on localhost:%d\n", port)

	// Wait for signals...
	err := <-done
	if err != nil {
		log.Println(err)
	}

	// Gracefully shutdown the service.
	//service.Teardown()

}
