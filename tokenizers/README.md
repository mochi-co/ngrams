[![GoDoc](https://godoc.org/github.com/mochi-co/ngrams?status.svg)](https://godoc.org/github.com/mochi-co/ngrams/tokenizers)

### Tokenizers
Tokenizers can be passed to the ngrams.NewIndex function to change the data tokenization mechanism. More details can be found in the [ngrams README](https://github.com/mochi-co/ngrams/blob/master/README.md).

##### Default Word Tokenizer (default)
```go 

// New word tokenizer which includes line breaks as distinct tokens.
tk := NewDefaultWordTokenizer(false)

// New word tokenizer without tokenized line breaks.
tk := NewDefaultWordTokenizer(true)
```

New tokenizers can be created by satisfying the `tokenizers.Tokenizer` interface.
