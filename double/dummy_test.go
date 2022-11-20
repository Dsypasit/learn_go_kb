package double

import (
	"testing"
)

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{}
}

func TestFindItShouldReturnErrorWhenFirstNameOrLastEmpty(t *testing.T) {
	phonebook := &Phonebook{}
	want := ErrMissingArgs

	_, got := phonebook.Find(DummySearcher{}, "", "")

	if got != want {
		t.Errorf("Want %s, got %s", want, got)
	}
}
