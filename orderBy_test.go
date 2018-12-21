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

func TestOrderbyBasicString(t *testing.T) {
	l, err := New([]string{"cde", "abd"})
	if err != nil {
		panic(err)
	}
	result := l.OrderBy(func(x string) string {
		return x
	})

	expected := []string{"abd", "cde"}
	resultString := result.Result().([]string)
	for i, x := range expected {
		if resultString[i] != x {
			t.Errorf("invalid response. Expected: %s. Found %s", x, resultString[i])
		}
	}
}

func TestOrderbyBasicFloat64(t *testing.T) {
	l, err := New([]float64{2.5, 3.5, 2, 1})
	if err != nil {
		panic(err)
	}
	result := l.OrderBy(func(x float64) float64 {
		return x
	})

	expected := []float64{1.0, 2.0, 2.5, 3.5}
	resultFloat64 := result.Result().([]float64)
	for i, x := range expected {
		if resultFloat64[i] != x {
			t.Errorf("invalid response. Expected: %f. Found %f", x, resultFloat64[i])
		}
	}
}

func TestOrderbyMultiple(t *testing.T) {
	l, err := New([]float64{2.5, 3.5, 2, 1, 6, 7})
	if err != nil {
		panic(err)
	}
	result := l.Where(func(x float64) bool {
		if x < 6 {
			return true
		}
		return false
	}).OrderBy(func(x float64) float64 {
		return x
	})

	expected := []float64{1.0, 2.0, 2.5, 3.5}
	resultFloat64 := result.Result().([]float64)
	for i, x := range expected {
		if resultFloat64[i] != x {
			t.Errorf("invalid response. Expected: %f. Found %f", x, resultFloat64[i])
		}
	}
}
