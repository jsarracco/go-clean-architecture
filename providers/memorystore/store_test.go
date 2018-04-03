package memorystore_test

import (
	"clean-go/domain"
	"testing"
)

func Test_sortGreetingsByDate(t *testing.T) {
	type args struct {
		g []domain.Greeting
	}
	tests := []struct {
		name string
		args args
		want []domain.Greeting
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: external comparisons
			//if got := sortGreetingsByDate(tt.args.g); !reflect.DeepEqual(got, tt.want) {
			//t.Errorf("sortGreetingsByDate() = %v, want %v", got, tt.want)
			//}
		})
	}
}
