package main

import "fmt"

func main() {
	store := make([]int, 1<<26)
	//fmt.Println(1 << 26)
	var q int
	var s, a, b uint32
	fmt.Scanf("%d %d %d %d", &q, &s, &a, &b)
	var ans uint32
	for q > 0 {
		x := s / 2
		//fmt.Println(x)
		//fmt.Println(x / 32)
		if s&1 == 1 {
			store[x/32] |= 1 << ( x % 32)
			ans += x
		} else if store[x/32]&(1<<(x%32)) > 0 {
			store[x/32] ^= 1 << x % 32
			ans -= x
		}
		s = s*a + b
		q--
	}

	println(ans)
}
