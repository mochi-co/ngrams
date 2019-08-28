// protoc -I . --go_out=plugins=grpc:. ./v1/ngrams.proto

package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"

	"github.com/mochi-co/ngrams"
	v1 "github.com/mochi-co/ngrams/cmd/grpc/v1"

	"github.com/jamiealquiza/envy"
)

// main is our entrypoint for the service.
func main() {
	log.Println("Starting gRPC trigram service...")

	// Prepare signal catching to ensure the webserver and indexer
	// shuts down cleanly.
	sigs := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- errors.New("caught signal")
	}()

	// Optionally override the port the gRPC server serves on.
	port := flag.Int("port", 50051, "port to serve grpc on")
	envy.Parse("NGRAMS") // Expose environment variables as NGRAMS_PORT, etc.

	// Configure the Ngrams indexer to index trigrams and use the default options.
	server := &ngramService{
		index: ngrams.NewIndex(3, nil),
	}

	// Setup the gRPC server with our ngram service.
	l, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterNgramServiceServer(s, v1.NgramServiceServer(server))

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for signals...
	err = <-done
	if err != nil {
		log.Println(err)
	}

	// Gracefully shutdown the index.
	server.index.Close()

}

// ngramService is a gRPC service for serving ngrams.
type ngramService struct {

	// index is the ngrams index, which will store and tokenize ngrams.
	index *ngrams.Index
}

// Learn trains the indexer on a body of data.
func (s *ngramService) Learn(ctx context.Context, req *v1.LearnRequest) (resp *v1.LearnResponse, err error) {

	tokens, err := s.index.Parse(req.Body)
	if err != nil {
		return
	}

	resp = &v1.LearnResponse{
		ParsedTokens: int64(len(tokens)),
	}

	return
}

// Generate trains the indexer on a body of data.
func (s *ngramService) Generate(ctx context.Context, req *v1.GenerateRequest) (resp *v1.GenerateResponse, err error) {

	var tokenLen int64 = 50

	if req.Limit > 0 {
		tokenLen = req.Limit
	}

	out, err := s.index.Babble("", int(tokenLen))
	if err != nil {
		return
	}

	resp = &v1.GenerateResponse{
		Body:  out,
		Limit: tokenLen,
	}

	return
}
