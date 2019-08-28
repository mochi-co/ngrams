[![Build Status](https://travis-ci.com/mochi-co/ngrams.svg?token=59nqixhtefy2iQRwsPcu&branch=master)](https://travis-ci.com/mochi-co/ngrams)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/mochi-co/ngrams/issues)
[![codecov](https://codecov.io/gh/mochi-co/ngrams/branch/master/graph/badge.svg?token=6vBUgYVaVB)](https://codecov.io/gh/mochi-co/ngrams)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/52ee14bd3a924979a26eb7da11d22a8f)](https://www.codacy.com/app/mochi-co/ngrams?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mochi-co/ngrams&amp;utm_campaign=Badge_Grade)
[![GoDoc](https://godoc.org/github.com/mochi-co/ngrams?status.svg)](https://godoc.org/github.com/mochi-co/ngrams)

## What is Ngrams?
`Ngrams` is a simple N-gram index capable of learning from a corpus of data and generating a random output in the same style. The index and tokenization systems are implemented as interfaces, so you can roll your own solution.

## Quick Start
You can test `ngrams` by running the small REST webserver in `cmd/rest/trigrams.go`. There's also a gRPC-proto example in `cmd/grpc`.
```
$ git clone https://github.com/mochi-co/ngrams.git
$ cd ngrams
```
```
$ go test -cover ./...
```
```
$ go build cmd/rest/trigrams.go
or
$ go run cmd/rest/trigrams.go
```

The webserver will serve two endpoints:
##### POST `localhost:8080/learn` 
Indexes a plain-text body of data. Training texts can be found in `training`.
```
$ curl --data-binary @"training/pride-prejudice.txt" localhost:8080/learn
# {"parsed_tokens":139394}

curl -d 'posting a string of text' -H "Content-Type: text/plain" -X POST localhost:8080/learn
# {"parsed_tokens":5}
```

##### GET `localhost:8080/generate[?limit=n]` 
Generates a random output from the learned ngrams. The `limit` query param can be used to change the number of tokens used to create the output (default 50).
```
$ curl localhost:8080/generate
# {
  "body": "They have solaced their wretchedness, however, and had a most conscientious
  		and polite young man, that might be able to keep him quiet. The arrival of the
  		fugitives. Not so much wickedness existed in the only one daughter will be 
  		having a daughter married.",
  "limit": 50
}
$ curl localhost:8080/generate?limit=10
# {
	"body": "Of its late possessor, she added, so untidy.",
	"limit": 10
}
```

## Basic Usage
An example of usage as a library can be found in `cmd/rest/trigrams.go`. The trigrams example uses the `tokenizers.DefaultWord` tokenizer, which will parse and format ngrams based on general latin-alphabet rules. 

```go
import "github.com/mochi-co/ngrams"
```

```go
// Initialize a new ngram index for 3-grams (trigrams), with default options.
index = ngrams.NewIndex(3, nil)

// Parse and index ngrams from a dataset.
tokens, err := index.Parse("to be or not to be, that is the question.")

// Generate a random sequence from the indexed ngrams.
out, err := index.Babble("to be", 50)

```

### Custom Index Initialization
Both the data storage and tokenization mechanisms for the indexer can be replaced by satisfying the store and tokenizer interfaces, allowing the indexer to be adjusted for different purposes. The index supports bigrams, trigrams, quadgrams, etc, by way of changing the NewIndex `n` (3) value.

```go
// Initialize with custom tokenizers and memory stores.
// The DefaultWordTokenizer take a bool to strip linebreaks in parsed text.
index = ngrams.NewIndex(3, ngrams.Options{
	Store: stores.NewMemoryStore(),
	Tokenizer: tokenizers.NewDefaultWordTokenizer(true),
})
```

### Tokenizers [![GoDoc](https://godoc.org/github.com/mochi-co/ngrams?status.svg)](https://godoc.org/github.com/mochi-co/ngrams/tokenizers)
A tokenizer consists of a `Tokenize` method which is used to parse the input data into ngram tokens, and a `Format` method which pieces them back together in an expected format. The library uses the `tokenizers.DefaultWord` tokenizer by default, which is a simple tokenizer for parsing most latin-based languages (english, french, etc) into ngram tokens. 

The `Format` method will perform a best effort attempt at piecing any selected ngram tokens back together in a grammatically correct manner (or whichever is appropriate for the type of tokenization being performed). 

New tokenizers can be created by satisfying the `tokenizers.Tokenizer` interface, such as for the tokenization of CJK datasets or amino sequences.

### Stores [![GoDoc](https://godoc.org/github.com/mochi-co/ngrams?status.svg)](https://godoc.org/github.com/mochi-co/ngrams/stores)
By default, the index uses an in-memory store, `stores.MemoryStore`. This is a basic memory store which stores the ngrams as-is. It's great for small examples, but if you were indexing millions of tokens it would be good to think about compression or aliasing. 

New stores can be created by satisfying the `stores.Store` interface.


## Contributions
Contributions to `ngrams` is encouraged. Open an [issue](https://github.com/mochi-co/ngrams/issues) to report a bug or make a feature request. Well-considered pull requests are welcome!

## License
MIT.

Training texts are sourced from [Project Gutenberg](https://www.gutenberg.org).