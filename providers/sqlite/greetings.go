package sqlite

import (
	"context"
	"database/sql"
	"log"

	"github.com/jsarracco/clean-go/domain"
	"github.com/jsarracco/clean-go/engine"
)

type (
	greetingRepository struct {
		db *sql.DB
	}

	// greeting is the internal struct we use for persistence
	// it allows us to attach the datastore PropertyLoadSaver
	// methods to the struct that we don't own
	greeting struct {
		domain.Greeting
	}
)

// This ensures greeting repository conforms to the interface GreetingRepository
// in this case greeting repo requires Put and List methods
func newGreetingRepository(db *sql.DB) engine.GreetingRepository {
	return &greetingRepository{db}
}

func (r greetingRepository) Type() string {
	return "sqlite"
}

func (r greetingRepository) Put(c context.Context, g *domain.Greeting) {
	query := "insert into greetings(author, content, date) values(?, ?, ?)"
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal("statement prepare fail", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.Author, g.Content, g.Date)
	if err != nil {
		log.Fatal("insert fail", err)
	}
	tx.Commit()
}

func (r greetingRepository) List(c context.Context, query *engine.Query) []*domain.Greeting {
	q := translateQuery(r.db, query)
	g := []*domain.Greeting{}
	g = q.GetAll(g)
	return g
}
