package double

import (
	"log"
	"testing"
)

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LasName:   lastName,
		Phone:     ss.phone,
	}
}

func TestFindReturnPerson(t *testing.T) {
	fakePhone := "+00 00 000 000"
	phonebook := &Phonebook{}

	phone, _ := phonebook.Find(StubSearcher{fakePhone}, "Jan", "Dos")

	log.Printf("%v %v\n", phone, fakePhone)
	if phone != fakePhone {
		t.Errorf("Want %s, got %s", fakePhone, phone)
	}
}
