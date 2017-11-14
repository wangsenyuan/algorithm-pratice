package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n, d int
		fmt.Scanf("%d %d", &n, &d)
		var a, b string
		fmt.Scanf("%s", &a)
		fmt.Scanf("%s", &b)
		res := solve(n, d, a, b)
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func solve(n int, d int, a string, b string) bool {
	cnt := make([]int, 26)
	for i := 0; i < d; i++ {
		for j := i; j < len(a); j += d {
			cnt[a[j]-'a']++
			cnt[b[j]-'a']--
		}
		for k := 0; k < len(cnt); k++ {
			if cnt[k] != 0 {
				return false
			}
		}
	}
	return true
}
