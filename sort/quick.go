package sort

func quick(s []int) []int {
	quickSort(s, 0, len(s)-1)
	return s
}

func quickSort(s []int, low int, high int) {
	if low < high {
		partitionIndex := partition(s, low, high)
		quickSort(s, low, partitionIndex-1)
		quickSort(s, partitionIndex+1, high)
	}
}

func partition(s []int, low int, high int) int {
	i := low - 1
	pivot := s[high]
	for j := low; j < high; j++ {
		if s[j] <= pivot {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i+1], s[high] = s[high], s[i+1]
	return i + 1
}
