package linq

import "testing"

func TestWhereIntBasic(t *testing.T) {
	fp, _ := New([]int{2, 4, 5, 8})
	result := fp.Where(func(x int) bool {
		if x > 2 {
			return true
		}
		return false
	})
	expected := []int64{4, 5, 8}
	resultInt := result.Result().([]int64)
	for i, x := range expected {
		if resultInt[i] != x {
			t.Errorf("invalid response. Expected: %d. Found %d", x, resultInt[i])
		}
	}
}

func TestWhereIntMultiple(t *testing.T) {
	fp, _ := New([]int{2, 4, 5, 8, 15, 20})
	result := fp.Where(func(x int) bool {
		if x > 2 {
			return true
		}
		return false
	}).Where(func(x int) bool {
		if x > 8 {
			return true
		}
		return false
	})
	expected := []int64{15, 20}
	resultInt := result.Result().([]int64)
	for i, x := range expected {
		if resultInt[i] != x {
			t.Errorf("invalid response. Expected: %d. Found %d", x, resultInt[i])
		}
	}
}

func TestWhereFloatBasic(t *testing.T) {
	fp, _ := New([]float64{2.5, 2.7, 2.9, 3.2})
	result := fp.Where(func(x float64) bool {
		if x >= 2.9 {
			return true
		}
		return false
	})
	expected := []float64{2.9, 3.2}
	resultInt := result.Result().([]float64)
	for i, x := range expected {
		if resultInt[i] != x {
			t.Errorf("invalid response. Expected: %f. Found %f", x, resultInt[i])
		}
	}
}
