package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

var cs [26]int

func firstUniqChar(s string) int {
	for i := 0; i < len(cs); i++ {
		cs[i] = -1
	}

	for j, x := range s {
		i := x - 'a'
		if cs[i] < 0 {
			cs[i] = j
		} else if cs[i] < len(s) {
			cs[i] += len(s)
		}
	}

	j := -1

	for i := range cs {
		if cs[i] < 0 || cs[i] >= len(s) {
			continue
		}
		if j == -1 || cs[i] < j {
			j = cs[i]
		}
	}
	return j
}
