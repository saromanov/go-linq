package linq

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	l, err := New([]string{"Alfa Romeo", "Aston Martin", "Audi", "Nissan", "Chevrolet", "Chrysler", "Dodge", "BMW",
		"Ferrari", "Bentley", "Ford", "Lexus", "Mercedes", "Toyota", "Volvo", "Subaru", "Жигули :)"})
	if err != nil {
		panic(err)
	}
	result := l.First(func(x string) bool {
		return true
	})
	resp := result.Result().([]string)
	assert.Equal(t, resp[0], "Alfa Romeo", "strings should be equal")
}

func TestFirstEmpty(t *testing.T) {
	l, err := New([]string{})
	if err != nil {
		panic(err)
	}
	result := l.First(func(x string) bool {
		return true
	})
	assert.Nil(t, result.Result())
}

func TestFirstWithPrefix(t *testing.T) {
	l, err := New([]string{"Alfa Romeo", "Aston Martin", "Audi", "Nissan", "Chevrolet", "Chrysler", "Dodge", "BMW",
		"Ferrari", "Bentley", "Ford", "Lexus", "Mercedes", "Toyota", "Volvo", "Subaru", "Жигули :)"})
	if err != nil {
		panic(err)
	}
	result := l.First(func(x string) bool {
		return strings.HasPrefix(x, "T")
	})
	resp := result.Result().([]string)
	assert.Equal(t, resp[0], "Toyota", "strings should be equal")
}
