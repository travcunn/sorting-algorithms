package sort

func MergeSort(values []int) []int {
	return split(values)
}

func merge(left []int, right []int) []int {
	sorted_list := make([]int, 0)
	var x int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				x, left = left[0], left[1:]
				sorted_list = append(sorted_list, x)
			} else {
				x, right = right[0], right[1:]
				sorted_list = append(sorted_list, x)
			}
		} else if len(left) > 0 {
			x, left = left[0], left[1:]
			sorted_list = append(sorted_list, x)
		} else if len(right) > 0 {
			x, right = right[0], right[1:]
			sorted_list = append(sorted_list, x)
		}
	}
	return sorted_list
}

func split(values []int) []int {
	if len(values) == 1 {
		return values
	}

	split_value := len(values) / 2
	left := split(values[:split_value])
	right := split(values[split_value:])

	return merge(left, right)
}
