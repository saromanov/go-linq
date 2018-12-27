package linq

import (
	"strings"
	"testing"
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
	if strings.Compare(resp[0], "Alfa Romeo") != 0 {
		t.Errorf("strings is not equal")
	}
}

func TestFirstEmpty(t *testing.T) {
	l, err := New([]string{})
	if err != nil {
		panic(err)
	}
	result := l.First(func(x string) bool {
		return true
	})
	resp := result.Result()
	if resp != nil {
		t.Errorf("result should be empty")
	}
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
	if strings.Compare(resp[0], "Toyota") != 0 {
		t.Errorf("strings is not equal")
	}
}
