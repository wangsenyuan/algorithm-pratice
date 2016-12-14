package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		t--
		var n int
		fmt.Scanf("%d\n", &n)
		var line string
		fmt.Scanf("%s\n", &line)
		ones, res := 0, int64(0)
		for i := 0; i < n; i++ {
			if line[i] == '1' {
				res += int64(ones)
				res++
				ones++
			}
		}
		fmt.Println(res)
	}
}
