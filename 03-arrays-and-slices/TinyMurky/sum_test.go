package tinymurky

import (
	"slices"
	"testing"
)

// Sum
func TestSum1(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	got := Sum1(arr)
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum2(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	got := Sum2(arr)
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestSumSlice(t *testing.T) {
	t.Run("it can be slice of 5", func(t *testing.T) {

		slice := []int{1, 2, 3, 4, 5}
		got := SumSlice(slice)
		want := 15
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("it can be slice of 3", func(t *testing.T) {
		slice := []int{1, 2, 3}
		got := SumSlice(slice)
		want := 6
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

// SumAll
func TestSumAll(t *testing.T) {
	t.Run("很多slice加在一起", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// if !reflect.DeepEqual(got ,want) {
		//   t.Errorf("got %+v want %+v", got, want)
		// }
		if !slices.Equal(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}

// SumAllTail
func TestSumAllTail(t *testing.T) {
	t.Run("不加第一個, non empty slice", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3, 4}, []int{1, 9})
		want := []int{9, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("不加第一個, with empty slice", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{1, 2, 3, 4})
		want := []int{0, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}
