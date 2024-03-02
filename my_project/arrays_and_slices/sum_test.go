package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	numbers2 := []int{1, 2, 3, 4}
	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d give %v", got, want, numbers)
		}
	})
	t.Run("multiple sums of slices", func(t *testing.T) {
		got := SumAll(numbers, numbers2)
		want := []int{6, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given %v %v", got, want, numbers, numbers2)
		}
	})
	t.Run("sum of tails of slices", func(t *testing.T) {
		got := SumAllTails(numbers, numbers2)
		want := []int{5, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given %v %v", got, want, numbers, numbers2)
		}
	})
}
