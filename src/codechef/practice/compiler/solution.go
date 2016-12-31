package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		var s string
		fmt.Scanf("%s", &s)

		left, res := 0, 0

		for i := 0; i < len(s); i++ {
			if left < 0 {
				break
			}

			if s[i] == '>' && left == 1 {
				res = i + 1
			}

			if s[i] == '>' {
				left--
			} else {
				left++
			}
		}

		fmt.Println(res)
	}
}
