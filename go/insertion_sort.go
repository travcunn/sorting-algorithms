package sort

func InsertionSort(values []int) {
	var x, j int
	for i := 0; i < len(values); i++ {
		x = values[i]
		j = i - 1
		for j >= 0 && values[j] > x {
			values[j+1] = values[j]
			j = j - 1
		}
		values[j+1] = x
	}
}
