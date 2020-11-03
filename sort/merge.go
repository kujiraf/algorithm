package sort

func mergeSort(s []int) {
	if len(s) <= 1 {
		return
	}

	center := len(s) / 2
	shallowCopied := make([]int, len(s))
	copy(shallowCopied, s)
	left := shallowCopied[:center]  // deep copy : shallowCopied -> left
	right := shallowCopied[center:] // deep copy : shallowCopied -> right

	mergeSort(left)
	mergeSort(right)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			s[k] = left[i]
			i++
		} else {
			s[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		s[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		s[k] = right[j]
		j++
		k++
	}
}
