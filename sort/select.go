package sort

func selectSort(s []int) []int {
	l := len(s)
	for i := 0; i < l-1; i++ {
		tmp := s[i]
		for j := i + 1; j < l; j++ {
			if tmp > s[j] {
				tmp, s[j] = s[j], tmp
			}
		}
		if s[i] != tmp {
			s[i] = tmp
		}
	}
	return s
}
