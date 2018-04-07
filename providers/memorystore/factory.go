package memorystore

import "github.com/jsarracco/clean-go/engine"

type (
	storageFactory struct{}
)

// NewStorage creates a new instance of this datastore storage factory
func NewStorage() engine.StorageFactory {
	return &storageFactory{}
}

// NewGreetingRepository creates a new datastore greeting repository
func (f *storageFactory) NewGreetingRepository() engine.GreetingRepository {
	return newGreetingRepository()
}
