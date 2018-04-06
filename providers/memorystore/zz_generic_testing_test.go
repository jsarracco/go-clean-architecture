package memorystore_test

import (
	"net/mail"
	"testing"
)

// Mail send supporting methods

type (
	MailMan struct{}

	EmailSender interface {
		Send(subject, body string, to ...*mail.Address)
	}

	testEmailSender struct {
		lastSubject string
		lastBody    string
		lastTo      []*mail.Address
	}
)

var (
	// make sure it satisfies the interface
	_ EmailSender = (*testEmailSender)(nil)
)

func (m *MailMan) Send(subject, body string, to ...*mail.Address) {
	// some code
}
func NewMailMan() *MailMan {
	return &MailMan{}
}
func NewTestEmailSender() *testEmailSender {
	return &testEmailSender{}
}

// Fib test supporting methods

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func (t *testEmailSender) Send(subject, body string, to ...*mail.Address) {
	t.lastSubject = subject
	t.lastBody = body
	t.lastTo = to
}

func SendWelcomeEmail(m EmailSender, to ...*mail.Address) {
	m.Send("Welcome", "you", to[0], to[1])
}

// Reverse testing supporting methods
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestSendWelcomeEmail(t *testing.T) {
	to1, _ := mail.ParseAddress("Alice <alice@example.com>")
	to2, _ := mail.ParseAddress("Alice2 <alice2@example.com>")

	// pointer instantiation
	sender := &testEmailSender{}

	// test new method as well
	sender = NewTestEmailSender()

	// test interface with fake mailman object
	mm := NewMailMan()
	SendWelcomeEmail(mm, to1, to2)

	// use interface with sender fake object
	SendWelcomeEmail(sender, to1, to2)

	// test fake object
	if sender.lastSubject != "Welcome" {
		t.Error("Subject line was wrong")
	}
	if sender.lastTo[0] != to1 && sender.lastTo[1] != to2 {
		t.Error("Wrong recipients")
	}
}

// testing tree
func TestFib(t *testing.T) {
	var fibTests = []struct {
		n        int // input
		expected int // expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestReverseToReturnReversedInputString(t *testing.T) {
	actualResult := Reverse("Hello")
	var expectedResult = "olleH"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
