package sort

func buble(s []int) []int {
	for l := len(s); l > 1; l-- {
		for i := 0; i < l-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}
