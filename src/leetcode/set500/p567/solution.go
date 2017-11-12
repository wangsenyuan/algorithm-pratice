package main

import (
	"reflect"
	"fmt"
)


func checkInclusion(s1 string, s2 string) bool {
	m1 := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
	}

	m2 := make(map[byte]int)

	j := 0
	for i := 0; i < len(s2); i++ {
		x := s2[i]
		m2[x]++

		if reflect.DeepEqual(m1, m2) {
			return true
		}

		for m2[x] > m1[x] && j <= i {
			y := s2[j]
			m2[y]--
			if m2[y] == 0 {
				delete(m2, y)
			}
			j++
		}
	}
	return false
}


func main() {
	s1 := "ab"
	s2 := "eidooab"
	fmt.Println(checkInclusion(s1, s2))
}