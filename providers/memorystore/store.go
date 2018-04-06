package memorystore

import (
	"clean-go/domain"
	"sort"
	"time"
)

const (
	ascending  string = "asc"
	descending string = "desc"
)

var greetingsData []domain.Greeting

func clearData() {
	greetingsData = nil
}

func saveGreeting(g *domain.Greeting) {
	// auto increment id for static save
	if g.ID == 0 {
		g.ID = int64(len(greetingsData))
	}

	// save value
	greetingsData = append(greetingsData, *g)
}

func getGreetingsByDate(direction string, offset int, limit int) []*domain.Greeting {
	g := sortGreetingsByDate(greetingsData)
	if direction == descending {
		g = reverseOrder(g)
	}
	return convertForView(limitResults(g, offset, limit))
}

// This assumes ids are already sorted by id
// for memory store so we don't need to sort via slice data
func getGreetingsByID(direction string, offset int, limit int) []*domain.Greeting {
	g := sortGreetingsByID(greetingsData)
	if direction == descending {
		g = reverseOrder(g)
	}
	return convertForView(limitResults(g, offset, limit))
}

func convertForView(g []domain.Greeting) []*domain.Greeting {
	o := make([]*domain.Greeting, len(g))
	for i := range g {
		greeting := greeting{}.Greeting
		o[i] = &greeting
		o[i].ID = g[i].ID
		o[i].Author = g[i].Author
		o[i].Content = g[i].Content
		o[i].Date = g[i].Date
	}
	return o
}

func reverseOrder(g []domain.Greeting) []domain.Greeting {
	var returnVal []domain.Greeting
	if len(g) == 0 {
		return g
	}
	for i := len(g) - 1; i >= 0; i-- {
		returnVal = append(returnVal, g[i])
	}
	return returnVal
}

func limitResults(g []domain.Greeting, offset int, limit int) []domain.Greeting {
	if len(g) == 0 {
		return g
	}
	if len(g) > limit {
		g = g[offset:limit]
	}
	return g
}

func sortGreetingsByID(g []domain.Greeting) []domain.Greeting {
	var sortedGreetings []domain.Greeting

	// build date sortable greetings
	p := make(GreetingsSortByID, len(g))
	for i := range g {
		p[i].ID = g[i].ID
		p[i].Date = g[i].Date
		p[i].Author = g[i].Author
		p[i].Content = g[i].Content
	}

	// sort it
	sort.Sort(p)

	// convert back to sorted greeting
	for i := range p {
		greeting := greeting{}.Greeting
		greeting.ID = p[i].ID
		greeting.Author = p[i].Author
		greeting.Content = p[i].Content
		greeting.Date = p[i].Date
		sortedGreetings = append(sortedGreetings, greeting)
	}

	return sortedGreetings
}

func sortGreetingsByDate(g []domain.Greeting) []domain.Greeting {
	var sortedGreetings []domain.Greeting

	// build date sortable greetings
	p := make(GreetingsSortByDate, len(g))
	for i := range g {
		p[i].ID = g[i].ID
		p[i].Date = g[i].Date
		p[i].Author = g[i].Author
		p[i].Content = g[i].Content
	}

	// sort it
	sort.Sort(p)

	// convert back to sorted greeting
	for i := range p {
		greeting := greeting{}.Greeting
		greeting.ID = p[i].ID
		greeting.Author = p[i].Author
		greeting.Content = p[i].Content
		greeting.Date = p[i].Date
		sortedGreetings = append(sortedGreetings, greeting)
	}

	// set sorted greetings by date
	return sortedGreetings
}

type (

	// Pair is a data structure to hold key/value pairs
	Pair struct {
		Key   string
		Value int64
	}

	// TimePair is a data structure to hold key/value pairs
	TimePair struct {
		ID      int64
		Date    time.Time
		Author  string
		Content string
	}

	// PairList is a slice of pairs that implements sort.Interface to sort by values
	PairList []Pair

	// TimeList is a slice of time pairs
	TimeList []TimePair

	// GreetingsSortByID holds greetings and has sort methods
	GreetingsSortByID []domain.Greeting

	// GreetingsSortByDate holds greetings and has sort methods
	GreetingsSortByDate []domain.Greeting
)

func (p GreetingsSortByID) Len() int           { return len(p) }
func (p GreetingsSortByID) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p GreetingsSortByID) Less(i, j int) bool { return p[i].ID < p[j].ID }

func (p GreetingsSortByDate) Len() int           { return len(p) }
func (p GreetingsSortByDate) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p GreetingsSortByDate) Less(i, j int) bool { return p[i].Date.Before(p[j].Date) }
