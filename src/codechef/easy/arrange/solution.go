package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t -= 1
		var k int
		var s string
		fmt.Scanf("%d %s\n", &k, &s)

		n := 1 << uint(k)
		bs := []byte(s)

		for i := 0; i < n; i++ {
			j := mirror(k, i)
			if j > i {
				bs[i], bs[j] = bs[j], bs[i]
			}
		}

		fmt.Println(string(bs))
	}
}

func mirror(k int, x int) int {
	y := 0
	for i := 0; i < k; i++ {
		if ((x >> uint(i)) & 1) == 1 {
			y |= 1 << uint(k - 1 - i)
		}
	}
	return y
}
