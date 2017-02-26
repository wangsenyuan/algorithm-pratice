package main

import "fmt"

func main() {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	fmt.Println(findLongestWord(s, d))
}

func findLongestWord(s string, d []string) string {
	n := len(d)
	ps := make([]int, n)

	for i := 0; i < len(s); i++ {
		x := s[i]

		for j := 0; j < n; j++ {
			if ps[j] == len(d[j]) {
				continue
			}
			if x == d[j][ps[j]] {
				ps[j]++
			}
		}
	}

	mx := 0
	for j := 0; j < n; j++ {
		if ps[j] < len(d[j]) {
			continue
		}
		if ps[j] > mx {
			mx = ps[j]
		}
	}

	if mx == 0 {
		return ""
	}
	var res = ""

	for j := 0; j < n; j++ {
		if ps[j] == mx && ps[j] == len(d[j]) {
			if res == "" || res > d[j] {
				res = d[j]
			}
		}
	}

	return res
}
