package memorystore

import (
	"clean-go/domain"
	"clean-go/engine"
	"context"
)

type (
	greetingRepository struct{}

	// greeting is the internal struct we use for persistence
	// it allows us to attach the datastore PropertyLoadSaver
	// methods to the struct that we don't own
	greeting struct {
		domain.Greeting
	}
)

func newGreetingRepository() engine.GreetingRepository {
	return &greetingRepository{}
}

func (r greetingRepository) Type() string {
	return "memory"
}

func (r greetingRepository) Put(c context.Context, g *domain.Greeting) {
	saveGreeting(g)
}

func (r greetingRepository) List(c context.Context, query *engine.Query) []*domain.Greeting {
	q := translateQuery(query)
	g := []*domain.Greeting{}
	g = q.GetAll(g)
	return g
}
