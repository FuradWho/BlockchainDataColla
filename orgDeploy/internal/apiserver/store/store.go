package store

var client Factory

// Factory defines the apiserver platform storage interface.
type Factory interface {
	Users() UserStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the apiserver store client.
func SetClient(factory Factory) {
	client = factory
}
