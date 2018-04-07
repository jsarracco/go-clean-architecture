// internal tests for memorystore query private methods
package memorystore

import (
	"reflect"
	"testing"

	"github.com/jsarracco/clean-go/engine"
)

func getEngineQuery() *engine.Query {
	q := engine.NewQuery("greeting").Order("date", engine.Descending).Slice(0, 3)
	return q
}

func Test_translateQuery(t *testing.T) {
	type args struct {
		query *engine.Query
	}
	tests := []struct {
		name string
		args args
		want *Query
	}{
		{
			name: "test1",
			args: args{
				query: engine.NewQuery("greeting").Order("date", engine.Descending).Slice(0, 3),
			},
			want: &Query{"date", "desc", 0, 3},
		},
		{
			name: "test2",
			args: args{
				query: engine.NewQuery("greeting").Order("id", engine.Ascending).Slice(15, 20),
			},
			want: &Query{"id", "asc", 15, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateQuery(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("translateQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
