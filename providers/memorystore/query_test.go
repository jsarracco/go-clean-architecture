package memorystore_test

import (
	"clean-go/domain"
	"clean-go/providers/memorystore"
	"reflect"
	"testing"
)

func TestQuery_GetAll(t *testing.T) {
	type fields struct {
		OrderByColumn string
		Direction     string
		Offset        int
		Limit         int
	}
	type args struct {
		g []*domain.Greeting
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*domain.Greeting
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &memorystore.Query{
				OrderByColumn: tt.fields.OrderByColumn,
				Direction:     tt.fields.Direction,
				Offset:        tt.fields.Offset,
				Limit:         tt.fields.Limit,
			}
			if got := q.GetAll(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
