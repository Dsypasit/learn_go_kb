package double

import "testing"

type FakeSearcher struct{}

func (fs FakeSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnEmptyStringForNoPerson(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeSearcher{}

	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Want '' got %s", phone)
	}
}
