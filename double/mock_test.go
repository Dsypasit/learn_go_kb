package double

import "testing"

type MockSearch struct {
	phone        string
	methodToCall map[string]bool
}

func (ms *MockSearch) Search(people []*Person, firstName string, lastName string) *Person {
	ms.methodToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LasName:   lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearch) Create(people []*Person, firstName string, lastName string) *Person {
	ms.methodToCall["Create"] = true
	return &Person{
		FirstName: firstName,
		LasName:   lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearch) ExpectToCall(methodName string) {
	if ms.methodToCall == nil {
		ms.methodToCall = make(map[string]bool)
	}

	ms.methodToCall[methodName] = false
}

func (ms *MockSearch) Verify(t *testing.T) {
	for methodName, called := range ms.methodToCall {
		if !called {
			t.Errorf("Expected to call %s, but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnPersonUsingMock(t *testing.T) {
	fakePhone := "+00 00 000 000"
	phonebook := &Phonebook{}
	mock := &MockSearch{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phonebook.Find(mock, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '' got %s", phone)
	}

	mock.Verify(t)
}
