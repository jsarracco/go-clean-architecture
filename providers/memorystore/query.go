package memorystore

import (
	"github.com/jsarracco/clean-go/domain"
	"github.com/jsarracco/clean-go/engine"
)

type (
	// Query object for memory storage
	Query struct {
		OrderByColumn string
		Direction     string
		Offset        int
		Limit         int
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
		g = getGreetingsByID(q.Direction, q.Offset, q.Limit)
	}

	if q.OrderByColumn == OrderColumnDate {
		g = getGreetingsByDate(q.Direction, q.Offset, q.Limit)
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

	if query.Offset > 0 {
		q.Offset = query.Offset
	}

	if query.Limit > 0 {
		q.Limit = query.Limit
	}

	return q
}
