package linq

import "testing"

func TestSelectBasic(t *testing.T) {
	fp, _ := New([]string{"foo", "bar", "foobar"})
	result := fp.Select(func(x string) int {
		return len(x)
	})
	expected := []int{3, 3, 6}
	resultInt := result.Result().([]int)
	for i, x := range expected {
		if resultInt[i] != x {
			t.Errorf("invalid response. Expected: %d. Found %d", x, resultInt[i])
		}
	}
}
