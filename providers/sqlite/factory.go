package sqlite

import (
	"database/sql"
	"log"

	"github.com/jsarracco/clean-go/engine"

	// Need to have an unused import here
	_ "github.com/mattn/go-sqlite3"
)

type (
	storageFactory struct {
		db *sql.DB
	}
)

// NewStorage creates a new instance of this datastore storage factory
func NewStorage(storagePath string) engine.StorageFactory {
	dbName := storagePath + "/guestbook.sqlite"
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table greetings (
		id integer not null primary key, 
		date datetime, 
		author varchar(50), 
		content varchar(50)
	);
	`

	// setup database schema
	_, err = db.Exec(sqlStmt)
	if err != nil && err.Error() != "table greetings already exists" {
		log.Printf("%q %s\n", err, sqlStmt)
	}

	return &storageFactory{db}
}

// NewGreetingRepository creates a new datastore greeting repository
func (f *storageFactory) NewGreetingRepository() engine.GreetingRepository {
	return newGreetingRepository(f.db)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
