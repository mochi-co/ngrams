[![GoDoc](https://godoc.org/github.com/mochi-co/ngrams?status.svg)](https://godoc.org/github.com/mochi-co/ngrams/stores)

### Stores
Stores can be passed to the ngrams.NewIndex function to change the data storage mechanism. More details can be found in the [ngrams README](https://github.com/mochi-co/ngrams/blob/master/README.md).

##### Memory Only store (default)
```go 
m := NewMemoryStore()
```

New stores can be created by satisfying the `stores.Store` interface.
