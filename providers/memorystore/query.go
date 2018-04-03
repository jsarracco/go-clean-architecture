package memorystore

import (
	"clean-go/domain"
	"clean-go/engine"
)

type (
	// Query object for memory storage
	Query struct {
		OrderByColumn string
		Direction     string
	}
)

const (
	// OrderColumnID order by id
	OrderColumnID string = "id"

	// OrderColumnDate order by id
	OrderColumnDate string = "date"
)

// GetAll returns all results
func (q *Query) GetAll(g []*domain.Greeting) []*domain.Greeting {

	if q.OrderByColumn == OrderColumnID {
		g = getGreetingsByID(q.Direction)
	}

	if q.OrderByColumn == OrderColumnDate {
		g = getGreetingsByDate(q.Direction)
	}

	return g
}

// translateQuery converts an engine (use case layer)
// query spec into a memorystore provider query
// Note: Currently only processes order and count
func translateQuery(query *engine.Query) *Query {
	q := &Query{}

	for _, order := range query.Orders {
		q.OrderByColumn = order.Property
		switch order.Direction {
		case engine.Ascending:
			q.Direction = "asc"
		case engine.Descending:
			q.Direction = "desc"
		}
	}

	return q
}
