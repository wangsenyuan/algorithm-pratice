package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		var a, b, c, d int
		fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

		res := float64(a) / float64(a + b)

		fmt.Printf("%f\n", res)

		t--
	}
}
