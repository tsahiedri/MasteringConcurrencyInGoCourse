package Exercise1_MergeSort

func Merge(left, right []int) []int {

	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		merged = append(merged, left...)
	} else if len(right) > 0 {
		merged = append(merged, right...)
	}
	return merged
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return Merge(left, right)
}
