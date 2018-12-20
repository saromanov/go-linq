package linq

import "testing"

func TestOrderbyBasic(t *testing.T) {
	l, err := New([]int{2, 5, 1, 3, 9, 8, 2, 3})
	if err != nil {
		panic(err)
	}
	result := l.OrderBy(func(x int) int {
		return x
	})

	expected := []int{1, 2, 2, 3, 3, 5, 8, 9}
	resultInt := result.Result().([]int)
	for i, x := range expected {
		if resultInt[i] != x {
			t.Errorf("invalid response. Expected: %d. Found %d", x, resultInt[i])
		}
	}
}
