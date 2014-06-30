package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func generateList(length int) []int {
	//Returns a slice of random integers.

	unsorted_list := make([]int, length)
	for i := 0; i < length; i++ {
		unsorted_list[i] = rand.Int()
	}

	return unsorted_list
}

func TestBogo(t *testing.T) {
	list := generateList(10)

	BogoSort(list)

	sorted := sort.IsSorted(sort.IntSlice(list))
	if !sorted {
		t.Errorf("The values were not sorted properly.")
	}
}
