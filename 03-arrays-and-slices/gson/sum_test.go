package sum

import (
	"reflect"
	"testing"
)

// Sum
func TestSum(t *testing.T) {

	t.Run("container size with 5", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := Sum(arr)
		want := 15
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("container with different size", func(t *testing.T) {
		slice := []int{1, 2, 3}
		got := Sum(slice)
		want := 6
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

// SumAll
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

// SumAllTail
func TestSumAllTail(t *testing.T) {
	t.Run("sum all tailed with non empty slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3, 4}, []int{4, 5})
		want := []int{9, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("sum all tailed with some empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{1, 2, 3, 4})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
