package memorystore

import (
	"clean-go/engine"
	"reflect"
	"testing"
)

func Test_newGreetingRepository(t *testing.T) {

	// setup mock repo
	s := NewStorage()
	greetingRepo := s.NewGreetingRepository()

	tests := []struct {
		name string
		want engine.GreetingRepository
	}{
		{
			name: "test1",
			want: greetingRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newGreetingRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGreetingRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
