package sort

func BubbleSort(values []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(values)-1; i++ {
			if values[i] > values[i+1] {
				sorted = false
				values[i], values[i+1] = values[i+1], values[i]
			}
		}
	}
}
