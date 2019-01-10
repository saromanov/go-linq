package linq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	l, err := New([]string{"Alfa Romeo", "Aston Martin", "Audi", "Nissan", "Chevrolet", "Chrysler", "Dodge", "BMW",
		"Ferrari", "Bentley", "Ford", "Lexus", "Mercedes", "Toyota", "Volvo", "Subaru"})
	if err != nil {
		t.Fatal(err)
	}
	result := l.Zip([]int{1, 2, 3}, func(i int, x string) string {
		return fmt.Sprintf("%d:%s", i, x)
	})
	res := result.Result().([]string)
	assert.Equal(t, res, []string{"1:Alfa Romeo", "2:Aston Martin", "3:Audi"}, "results should be equal")

	l, err = New([]int{2, 4, 5})
	if err != nil {
		t.Fatal(err)
	}
	result = l.Zip([]int{1, 2, 3}, func(i, j int) int {
		return i + j
	})
	res2 := result.Result().([]int)
	assert.Equal(t, res2, []int{3, 6, 8}, "results should be equal")

}

func TestZipInvalidMethod(t *testing.T) {
	l, err := New([]string{"Alfa Romeo", "Aston Martin", "Audi", "Nissan", "Chevrolet", "Chrysler", "Dodge", "BMW",
		"Ferrari", "Bentley", "Ford", "Lexus", "Mercedes", "Toyota", "Volvo", "Subaru"})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	l.Zip([]int{1, 2, 3}, 5)

}
