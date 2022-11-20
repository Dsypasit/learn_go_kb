package testlibrary

import (
	"testing"

	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		pp := &Person{FirstName: "Jane"}

		if assert.NotNil(t, pp) {
			assert.Equal(t, "Jane", pp.FirstName)
		}
	})
	t.Run("equal", func(t *testing.T) {
		want := 555
		got := 555

		assert.Equal(t, want, got, "they should equal")
	})
}

func TestIs(t *testing.T) {
	is := is.New(t)
	pp := &Person{FirstName: "Jane"}

	is.Equal(pp, &Person{FirstName: "Jane"})
}
