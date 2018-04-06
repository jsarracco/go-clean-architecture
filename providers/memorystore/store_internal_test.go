// internal tests for memorystore store private methods
package memorystore

import (
	"clean-go/domain"
	"reflect"
	"testing"
	"time"
)

// was time.Now().Add(12 * time.Hour)
// was time.Now().Add(24 * time.Hour)
// was time.Now()

func getTime(offsetHours int) time.Time {
	dur := time.Duration(offsetHours) * time.Hour
	timeStr := "2012-11-01T12:08:41+00:00"
	time, _ := time.Parse(time.RFC3339, timeStr)
	time = time.Add(dur)
	return time
}

func getMockData() []domain.Greeting {
	var returnVal []domain.Greeting
	g := domain.Greeting{Author: "post+12", ID: 5, Date: getTime(12), Content: "test"}
	returnVal = append(returnVal, g)
	g = domain.Greeting{Author: "postnow", ID: 1, Date: getTime(0), Content: "test"}
	returnVal = append(returnVal, g)
	g = domain.Greeting{Author: "post+24", ID: 15, Date: getTime(24), Content: "test"}
	returnVal = append(returnVal, g)
	return returnVal
}

func getSortedData() []domain.Greeting {
	var returnVal []domain.Greeting
	g := domain.Greeting{Author: "postnow", ID: 1, Date: getTime(0), Content: "test"}
	returnVal = append(returnVal, g)
	g = domain.Greeting{Author: "post+12", ID: 5, Date: getTime(12), Content: "test"}
	returnVal = append(returnVal, g)
	g = domain.Greeting{Author: "post+24", ID: 15, Date: getTime(24), Content: "test"}
	returnVal = append(returnVal, g)
	return returnVal
}

func Test_sortGreetingsByDate(t *testing.T) {
	type args struct {
		g []domain.Greeting
	}
	dataSorted := getSortedData()
	dataArgs := args{g: getMockData()}
	tests := []struct {
		name string
		args args
		want []domain.Greeting
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: dataArgs,
			want: dataSorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortGreetingsByDate(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortGreetingsByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortGreetingsByID(t *testing.T) {
	type args struct {
		g []domain.Greeting
	}
	dataSorted := getSortedData()
	dataArgs := args{g: getMockData()}
	tests := []struct {
		name string
		args args
		want []domain.Greeting
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: dataArgs,
			want: dataSorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortGreetingsByDate(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortGreetingsByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveGreeting(t *testing.T) {
	type args struct {
		g *domain.Greeting
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				g: &domain.Greeting{ID: 5, Author: "post+12", Content: "test", Date: getTime(12)},
			},
			want: "post+12",
		},
		{
			name: "test2",
			args: args{
				g: &domain.Greeting{ID: 0, Author: "post", Content: "test", Date: getTime(12)},
			},
			want: "post",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveGreeting(tt.args.g)
			g := getGreetingsByDate("asc", 0, 50)
			if got := g[len(g)-1].Author; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("author = %v, want %v", got, tt.want)
			}
		})
	}
}
