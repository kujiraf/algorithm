package sort

func insertion(s []int) []int {
	l := len(s)
	for i := 0; i < l-1; i++ {
		if s[i] > s[i+1] {
			reverse(s, i)
		}
	}
	return s
}

func reverse(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
	if (i > 0) && (s[i-1] > s[i]) {
		reverse(s, i-1)
	}
}
