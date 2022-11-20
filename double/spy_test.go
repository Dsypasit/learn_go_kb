package double

import "testing"

type SpySearcher struct {
	phone            string
	searchWantCalled bool
}

func (sp *SpySearcher) Search(people []*Person, firstName string, lastName string) *Person {
	sp.searchWantCalled = true
	return &Person{
		FirstName: firstName,
		LasName:   lastName,
		Phone:     sp.phone,
	}
}

func TestFindCallSearchAndReturnPerson(t *testing.T) {
	fakePhone := "+00 00 000 000"
	phonebook := &Phonebook{}

	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWantCalled {
		t.Errorf("Search method doesn't call")
	}

	if phone != fakePhone {
		t.Errorf("Want %s, got %s", fakePhone, phone)
	}
}
