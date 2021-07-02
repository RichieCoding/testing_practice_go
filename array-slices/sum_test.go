package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{5, 6, 7, 10})
		want := []int{6, 28}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})
}

func BenchmarkSumAll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{5, 6, 7, 10})
	}
}