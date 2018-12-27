package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhereIntBasic(t *testing.T) {
	fp, _ := New([]int{2, 4, 5, 8})
	result := fp.Where(func(x int) bool {
		if x > 2 {
			return true
		}
		return false
	})
	expected := []int{4, 5, 8}
	resultInt := result.Result().([]int)
	assert.Equal(t, resultInt, expected, "slices is not equal")
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
	expected := []int{15, 20}
	resultInt := result.Result().([]int)
	assert.Equal(t, resultInt, expected, "slices is not equal")
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
	assert.Equal(t, resultInt, expected, "slices is not equal")
}
