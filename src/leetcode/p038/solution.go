package main

import "fmt"

func countAndSay(n int) string {
	s := make([]byte, 0, 100)
	s = append(s, '1')

	for i := 1; i < n; i++ {
		t := make([]byte, 0, len(s))
		for j := 0; j < len(s); {
			k := j
			for ; k < len(s) && s[k] == s[j]; k++ {

			}
			t = append(t, byte('0'+(k-j)))
			t = append(t, s[j])
			j = k
		}

		s = t
	}

	return string(s)
}

func main() {
	fmt.Printf("%dth count and say is %s", 5, countAndSay(5))
}
