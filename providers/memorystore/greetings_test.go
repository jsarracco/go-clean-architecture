package memorystore_test

import (
	"clean-go/domain"
	"clean-go/engine"
	"context"
	"reflect"
	"testing"
	"time"
)

type (
	fakeGreetingRepository struct{}

	// greeting is the internal struct we use for persistence
	// it allows us to attach the datastore PropertyLoadSaver
	// methods to the struct that we don't own
	greeting struct {
		domain.Greeting
	}
)

var (
	savedData []*domain.Greeting
)

func newFakeGreetingRepository() engine.GreetingRepository {
	return &fakeGreetingRepository{}
}

func (r fakeGreetingRepository) Type() string {
	return "fakerepository"
}

func (r fakeGreetingRepository) Put(c context.Context, g *domain.Greeting) {
	savedData = append(savedData, g)
}

func (r fakeGreetingRepository) List(c context.Context, query *engine.Query) []*domain.Greeting {
	return savedData
}

// Gets a static time with offset
func getTime(offsetHours int) time.Time {
	dur := time.Duration(offsetHours) * time.Hour
	timeStr := "2012-11-01T12:08:41+00:00"
	time, _ := time.Parse(time.RFC3339, timeStr)
	time = time.Add(dur)
	return time
}

func getMockData() []*domain.Greeting {
	var returnVal []*domain.Greeting
	g1 := domain.Greeting{Author: "post+12", ID: 5, Date: getTime(12), Content: "test"}
	returnVal = append(returnVal, &g1)
	g2 := domain.Greeting{Author: "postnow", ID: 1, Date: getTime(0), Content: "test"}
	returnVal = append(returnVal, &g2)
	g3 := domain.Greeting{Author: "post+24", ID: 15, Date: getTime(24), Content: "test"}
	returnVal = append(returnVal, &g3)
	return returnVal
}

func getGreetingValue(g *domain.Greeting) domain.Greeting {
	return domain.Greeting{ID: g.ID, Author: g.Author, Content: g.Content, Date: g.Date}
}

func Test_greetingRepository_Type(t *testing.T) {
	greetingRepo := newFakeGreetingRepository()
	type (
		greetingRepository engine.GreetingRepository
	)
	tests := []struct {
		name string
		r    greetingRepository
		want string
	}{
		{
			name: "test1",
			r:    greetingRepo,
			want: "fakerepository",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := greetingRepo
			if got := r.Type(); got != tt.want {
				t.Errorf("greetingRepository.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greetingRepository_PutAndList(t *testing.T) {
	// setup mock repo
	greetingRepo := newFakeGreetingRepository()

	// setup mock context
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// setup mock query
	query := engine.NewQuery("greeting").Order("id", engine.Ascending).Slice(0, 5)

	// format date string
	// dateString := getTime(0).Format(time.RFC3339)
	mockData := getMockData()

	type (
		greetingRepository engine.GreetingRepository
	)
	type args struct {
		c context.Context
		g *domain.Greeting
	}
	tests := []struct {
		name string
		r    greetingRepository
		args args
		want domain.Greeting
	}{
		{
			name: "test1",
			r:    greetingRepo,
			args: args{
				c: ctx,
				g: &domain.Greeting{ID: 5, Author: "post+12", Content: "test", Date: getTime(12)},
			},
			want: getGreetingValue(mockData[0]),
		},
		{
			name: "test2",
			r:    greetingRepo,
			args: args{
				c: ctx,
				g: &domain.Greeting{ID: 1, Author: "postnow", Content: "test", Date: getTime(0)},
			},
			want: getGreetingValue(mockData[1]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := greetingRepo

			// save greeting
			r.Put(tt.args.c, tt.args.g)

			// get greeting to validate result
			g := r.List(ctx, query)
			greeting := getGreetingValue(g[len(g)-1])

			// have to do a deep equal here because of pointer comparisons
			if got := greeting; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
