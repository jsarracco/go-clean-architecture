package engine

import (
	"golang.org/x/net/context"

	"clean-go/domain"
)

type (
	// ListGreetingsRequest contains the count
	ListGreetingsRequest struct {
		Count int
	}

	// ListGreetingsResponse is a slice of greetings
	ListGreetingsResponse struct {
		Greetings []*domain.Greeting
	}
)

// List method (greeter.List) is the application business rule that
// all providers should be able to process
// This is the application presenter for the view
func (g *greeter) List(c context.Context, r *ListGreetingsRequest) *ListGreetingsResponse {
	q := NewQuery("greeting").Order("date", Descending).Slice(0, r.Count)
	return &ListGreetingsResponse{
		Greetings: g.repository.List(c, q),
	}
}
