package main

import "fmt"

func main() {
	fmt.Println(countSegments("Hello,  my name is  John"))
	fmt.Println(countSegments(" Hello,    my name is  John      "))
}

func countSegments(s string) int {
	res := 0

	last := 0 //1 means letter, 0 others
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			last = 1
			continue
		}

		if last == 1 {
			last = 0
			res++
		}
	}
	if last == 1 {
		res++
	}

	return res
}
