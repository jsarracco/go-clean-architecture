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

	// get mongodb storage provider with config parameters
	//store := mongodb.NewStorage(config.MongoURL)

	// get sqlite storage provider
	//store := sqlite.NewStorage()

	// get memory storage provider
	store := memorystore.NewStorage()

	// get new engine and set storage provider
	e := engine.NewEngine(store)

	// send engine(w/ storage provider) web adapter
	w := web.NewWebAdapter(e, true)

	http.ListenAndServe(":8080", w)

	// additional debugging stuff

	//fmt.Println(config.MongoURL)
	//g1 := domain.NewGreeting("test1", "test1")
	//g2 := domain.NewGreeting("test1", "test1")
	//fmt.Println("hey1", g1)
	//fmt.Println("hey2", g2)
}
