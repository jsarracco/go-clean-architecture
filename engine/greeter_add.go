package engine

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/jsarracco/clean-go/domain"
)

type (
	// AddGreetingRequest is a request facade
	AddGreetingRequest struct {
		Author  string
		Content string
	}

	// AddGreetingResponse is a response facade
	AddGreetingResponse struct {
		ID int64
	}
)

func (g *greeter) Add(c context.Context, r *AddGreetingRequest) *AddGreetingResponse {
	// this is where all our app logic would go - the
	// rules that apply to adding a greeting whether it
	// is being done via the web UI, a console app, or
	// whatever the internet has just been added to ...
	greeting := domain.NewGreeting(r.Author, r.Content)
	g.repository.Put(c, greeting)
	fmt.Println("testing add for id", greeting)
	return &AddGreetingResponse{
		ID: greeting.ID,
	}
}
