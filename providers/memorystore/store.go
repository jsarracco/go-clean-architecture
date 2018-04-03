package memorystore

import (
	"clean-go/domain"
	"fmt"
	"sort"
	"time"
)

const (
	ascending  string = "asc"
	descending string = "desc"
)

var greetingsData []domain.Greeting

func saveGreeting(g *domain.Greeting) {
	// increment id
	g.ID = int64(len(greetingsData))

	// save value
	greetingsData = append(greetingsData, *g)
}

func getGreetingsByDate(direction string) []*domain.Greeting {
	g := sortGreetingsByDate(greetingsData)
	if direction == descending {
		g = reverseOrder(g)
	}
	return convertForView(g)
}

// This assumes ids are already sorted by id
// for memory store so we don't need to sort via slice data
func getGreetingsByID(direction string) []*domain.Greeting {
	g := greetingsData
	if direction == descending {
		g = reverseOrder(g)
	}
	return convertForView(g)
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

func sortGreetingsByDate(g []domain.Greeting) []domain.Greeting {
	var sortedGreetings []domain.Greeting

	// build date sortable greetings
	p := make(TimeList, len(g))
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
		Value int
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
)

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func (p TimeList) Len() int           { return len(p) }
func (p TimeList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p TimeList) Less(i, j int) bool { return p[i].Date.Before(p[j].Date) }

//
// Testing Stuff - not used
//

func sortTesting() {
	noble := map[string]int{
		"Radon":   1,
		"Xenon":   5,
		"Krypton": 3,
		"Argon":   4,
		"Neon":    7,
		"Helium":  0,
	}
	sortGreetingsByID(noble, "desc")
}

func sortGreetingsByID(mapCollection map[string]int, direction string) {
	p := make(PairList, len(mapCollection))
	i := 0
	for k, v := range mapCollection {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	fmt.Println(p)
	fmt.Println(p[0].Key)
}

func getMockData() []domain.Greeting {
	var returnVal []domain.Greeting
	g := domain.Greeting{Author: "post+12", ID: 5, Date: time.Now().Add(12 * time.Hour), Content: "test"}
	returnVal = append(returnVal, g)
	g = domain.Greeting{Author: "postnow", ID: 1, Date: time.Now(), Content: "test"}
	returnVal = append(returnVal, g)
	g = domain.Greeting{Author: "post+24", ID: 15, Date: time.Now().Add(24 * time.Hour), Content: "test"}
	returnVal = append(returnVal, g)
	return returnVal
}

func sortGreetingsNotUsed(g []domain.Greeting) []domain.Greeting {
	// TODO: use sorting based on query passed in from
	// engine, currently only uses order and direction

	// Just flip the array
	g = reverseOrder(g)

	return g
}
