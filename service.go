package main

import (
	"github.com/mochi-co/trigrams-test/ngrams"
)

// Service provides the handlers and data transports for the server.
// It's an effective way of ensuring the endpoints have access to the
// same data and methods without resorting to globals, and allows the
// service to be imported as a package in other functions.
type Service struct {

	// Index is the index of stored trigrams (ngrams).
	Index *ngrams.Index
}

// New returns a pointer to a new Service.
func New() *Service {
	return &Service{}
}
