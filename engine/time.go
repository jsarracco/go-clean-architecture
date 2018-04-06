package engine

import "time"

var (
	// now returns the current UTC time
	// It is a replaceable function to allow for easy unit testing
	now = defaultNow
)

type (
	// Time sets global timestamp as base time
	Time struct {
		time time.Time
	}

	// local func type
	fn func() time.Time
)

// set it back to this function to restore normal functionality
func defaultNow() time.Time {
	return time.Now().UTC()
}

// NewTime factory gets new time object
func NewTime(t time.Time) *Time {
	return &Time{
		time: t,
	}
}

// GetNow gets domain time
func (t Time) GetNow() time.Time {
	return now()
}

// SetNow gets domain time
func (t Time) SetNow(f fn) {
	now = f
}

// ResetNow sets back to defaultNow function
func (t Time) ResetNow() time.Time {
	now = defaultNow
	return now()
}
