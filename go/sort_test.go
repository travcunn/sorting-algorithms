package sort

import (
	"math/rand"
	"testing"
)

func sorted(values []int) bool {
	length := len(values)
	for i := 0; i < length-1; i++ {
		if values[i] > values[i+1] {
			return false
		}
	}
	return true
}

func generateList(length int) []int {
	//Returns a slice of random integers.

	unsorted_list := make([]int, length)
	for i := 0; i < length; i++ {
		unsorted_list[i] = rand.Int()
	}

	return unsorted_list
}

func TestBogoSort(t *testing.T) {
	list := generateList(6)

	BogoSort(list)

	if !sorted(list) {
		t.Errorf("The values were not sorted properly.")
	}
}

func TestBubbleSort(t *testing.T) {
	list := generateList(1000)

	BubbleSort(list)

	if !sorted(list) {
		t.Errorf("The values were not sorted properly.")
	}
}

func TestInsertionSort(t *testing.T) {
	list := generateList(1000)

	InsertionSort(list)

	if !sorted(list) {
		t.Errorf("The values were not sorted properly.")
	}
}

func TestMergeSort(t *testing.T) {
	list := generateList(1000)

	list = MergeSort(list)

	if !sorted(list) {
		t.Errorf("The values were not sorted properly.")
	}
}
