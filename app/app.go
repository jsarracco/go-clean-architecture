// +build !appengine

package main

import (
	"clean-go/adapters/web"
	"clean-go/engine"
	"clean-go/providers/memorystore"
	"net/http"
)

// when running in traditional or 'standalone' mode
// we're going to use MongoDB as the storage provider
// and start the webserver running ourselves.
func main() {

	// get mongodb
	//store := mongodb.NewStorage(config.MongoURL)

	// get sqlite
	//store := sqlite.NewStorage()

	// get memory store
	store := memorystore.NewStorage()

	// set store
	e := engine.NewEngine(store)

	http.ListenAndServe(":8080", web.NewWebAdapter(e, true))

	//fmt.Println(config.MongoURL)
	//g1 := domain.NewGreeting("test1", "test1")
	//g2 := domain.NewGreeting("test1", "test1")
	//fmt.Println("hey1", g1)
	//fmt.Println("hey2", g2)
}
