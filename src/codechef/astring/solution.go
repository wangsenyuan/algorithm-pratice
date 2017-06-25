package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var str string
		fmt.Scanf("%s", &str)
		var k int
		fmt.Scanf("%d", &k)
		res := solve(str, k)
		fmt.Println(res)
	}
}
func solve(str string, k int) string {
	res := make([]byte, k)
	tmp := str
	i := 0
	for i < k {
		if k-i == len(tmp) {
			for j := 0; j < len(tmp); j++ {
				res[i] = tmp[j]
				i++
				j++
			}
		} else {
			x := -1
			for j := 0; j+k-i <= len(tmp); j++ {
				if x < 0 || tmp[j] < tmp[x] {
					x = j
				}
			}
			res[i] = tmp[x]
			i++
			tmp = tmp[x+1:]
		}

	}

	return string(res)
}
