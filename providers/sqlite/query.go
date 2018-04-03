package sqlite

import (
	"clean-go/domain"
	"clean-go/engine"
	"database/sql"
	"strconv"
)

type (

	// Query object for sql
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

	if query.Offset > 0 {
		q.Offset = query.Offset
	}

	if query.Limit > 0 {
		q.Limit = query.Limit
	}

	return q
}

// GetAll returns results based on query
func (q *Query) GetAll(g []*domain.Greeting) []*domain.Greeting {

	query := "SELECT * FROM greetings"

	if q.OrderByColumn == OrderColumnID {
		query += " ORDER BY id " + q.Direction
	}
	if q.OrderByColumn == OrderColumnDate {
		query += " ORDER BY date " + q.Direction
	}
	if q.Limit > 0 {
		query += " LIMIT " + strconv.Itoa(q.Limit)
	}
	if q.Offset > 0 {
		query += " OFFSET " + strconv.Itoa(q.Offset)
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
