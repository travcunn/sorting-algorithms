package sort

import (
	"math/rand"
	"sort"
)

func BogoSort(values []int) []int {
	for {
		if sort.IsSorted(sort.IntSlice(values)) {
			return values
		}

		// shuffle
		for x := range values {
			y := rand.Intn(x + 1)
			values[x], values[y] = values[y], values[x]
		}
	}
}
