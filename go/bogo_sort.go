package sort

import "math/rand"

func BogoSort(values []int) {
	for {
		if sorted(values) {
			break
		}

		// shuffle
		for x := range values {
			y := rand.Intn(x + 1)
			values[x], values[y] = values[y], values[x]
		}
	}
}
