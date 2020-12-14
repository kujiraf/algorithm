package main

import "fmt"

// 二つのスライスを比較し、同じ数字を持っている場合、その数字を持っている数が少ない方から、その数字を削除する
// 例：A:=[1, 2, 2, 3], B:=[1, 2, 4, 5] -> A[1, 2, 2, 3] , B[1, 4, 5]
func main() {

	a := []int{1, 2, 3, 4, 4, 5, 5, 8, 10}
	b := []int{4, 5, 5, 5, 6, 7, 8, 8, 10}

	ma := make(map[int]int)
	mb := make(map[int]int)

	for _, v := range a {
		ma[v]++
	}
	for _, v := range b {
		mb[v]++
	}

	rmMapA := makeRmMap(ma, mb)
	rmMapB := makeRmMap(mb, ma)

	a2 := remove(a, rmMapA)
	b2 := remove(b, rmMapB)

	fmt.Println(a2)
	fmt.Println(b2)
}

func remove(s []int, m map[int]struct{}) []int {
	for i := 0; i < len(s); i++ {
		v := s[i]
		if _, ok := m[v]; ok {
			if i == 0 {
				s = s[1:]
				i--
				continue
			}
			if i == len(s)-1 {
				s = s[:i]
				i--
				continue
			}
			s = append(s[0:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func makeRmMap(m1 map[int]int, m2 map[int]int) map[int]struct{} {
	rmS1Map := make(map[int]struct{})
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			continue
		}

		if v1 < v2 {
			rmS1Map[k] = struct{}{}
		}
	}
	fmt.Println(rmS1Map)
	return rmS1Map
}
