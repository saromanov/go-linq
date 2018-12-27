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
