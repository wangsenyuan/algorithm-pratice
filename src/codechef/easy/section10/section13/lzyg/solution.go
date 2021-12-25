package main

import (
	"fmt"
	"math"
)

func main() {
	var tc int

	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var m, n, k, l float64
		fmt.Scanf("%f %f %f %f\n", &m, &n, &k, &l)
		// fmt.Fprintf(os.Stderr, "%f %f %f %f", m, n, k, l)
		ans := math.Ceil((n * l) / (k + l))

		if ans > m {
			ans = m
		}

		fmt.Printf("%d\n", int(ans))
	}
}
