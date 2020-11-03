package search

func linearSearch(s []int, value int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			return i
		}
	}
	return -1
}

func binarySearch(s []int, value int) int {
	left, right := 0, len(s)-1
	for left <= right {
		center := (left + right) / 2
		if s[center] == value {
			return center
		}
		if s[center] > value {
			right = center - 1
		} else {
			left = center + 1
		}
	}
	return -1
}

func recursiveBinarySearch(s []int, value int, left int, right int) int {
	if left > right {
		return -1
	}
	center := (left + right) / 2
	if s[center] == value {
		return center
	}
	if s[center] > value {
		return recursiveBinarySearch(s, value, left, center-1)
	}
	return recursiveBinarySearch(s, value, center+1, right)
}
