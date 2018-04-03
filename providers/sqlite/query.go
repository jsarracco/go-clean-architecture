package sqlite

import (
	"clean-go/domain"
	"clean-go/engine"
	"database/sql"
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

var db *sql.DB

// translateQuery converts an application query spec into
// a sqlite datastore specific query
// SQLite does not have a query object
func translateQuery(dbengine *sql.DB, query *engine.Query) *Query {
	q := &Query{}
	db = dbengine

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

// GetAll returns results based on query
func (q *Query) GetAll(g []*domain.Greeting) []*domain.Greeting {

	query := "select * from greetings"

	if q.OrderByColumn == OrderColumnID {
		query += " order by id " + q.Direction
	}
	if q.OrderByColumn == OrderColumnDate {
		query += " order by date " + q.Direction
	}

	rows, err := db.Query(query)
	checkErr(err)

	for rows.Next() {
		greeting := domain.Greeting{}
		err = rows.Scan(&greeting.ID, &greeting.Date, &greeting.Author, &greeting.Content)
		checkErr(err)

		// convert to []*domain.Greeting
		t := make([]*domain.Greeting, 0)
		t = append(t, &greeting)
		g = append(g, t[0])
	}

	return g
}
