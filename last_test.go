package linq

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	l, err := New([]string{"Alfa Romeo", "Aston Martin", "Audi", "Nissan", "Chevrolet", "Chrysler", "Dodge", "BMW",
		"Ferrari", "Bentley", "Ford", "Lexus", "Mercedes", "Toyota", "Volvo", "Subaru", "Ferrari"})
	if err != nil {
		panic(err)
	}
	result := l.Last(func(x string) bool {
		return true
	})
	resp := result.Result().([]string)
	assert.Equal(t, resp[0], "Ferrari", "strings should be equal")
}

func TestLastEmpty(t *testing.T) {
	l, err := New([]string{})
	if err != nil {
		panic(err)
	}
	result := l.Last(func(x string) bool {
		return true
	})
	assert.Nil(t, result.Result())
}

func TestLastWithPrefix(t *testing.T) {
	l, err := New([]string{"Alfa Romeo", "Aston Martin", "Audi", "Nissan", "Chevrolet", "Chrysler", "Dodge", "BMW",
		"Ferrari", "Bentley", "Ford", "Lexus", "Mercedes", "Toyota", "Volvo", "Subaru", "Жигули :)"})
	if err != nil {
		panic(err)
	}
	result := l.Last(func(x string) bool {
		return strings.HasPrefix(x, "T")
	})
	resp := result.Result().([]string)
	assert.Equal(t, resp[0], "Toyota", "strings should be equal")
}

func TestLastInvalidMethod(t *testing.T) {
	l, err := New([]string{})
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	l.Last(func(x string) {
		return
	})

}
