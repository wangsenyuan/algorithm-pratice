package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		n, a, b := 0, 0, 0
		fmt.Scanf("%d %d %d\n", &n, &a, &b)
		c := dragnxor(n, a, b)
		fmt.Println(c)
		t--
	}
}

func dragnxor(n, a, b int) int {
	acnt := bits(n, a)
	bcnt := bits(n, b)

	res := (1 << uint(n)) - 1
	overlap := 0
	if acnt+bcnt <= n {
		overlap = n - acnt - bcnt
	} else {
		overlap = acnt + bcnt - n
	}
	//fmt.Printf("acnt, bcnt, overlap: %d, %d, %d\n", acnt, bcnt, overlap)

	return (res >> uint(overlap)) << uint(overlap)
}

func bits(n, x int) int {
	cnt := 0
	for i := 0; i < n; i++ {
		//x = (x >> uint(i))
		//fmt.Println(x)
		if (x>>uint(i))&1 == 1 {
			cnt++
		}
	}
	return cnt
}
