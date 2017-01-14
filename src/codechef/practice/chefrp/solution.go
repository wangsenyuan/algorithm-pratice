package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int
		fmt.Scanf("%d", &n)

		var N, minFreq, freq int

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &freq)
			if i == 0 || freq < minFreq {
				minFreq = freq
			}
			N += freq
		}

		if minFreq == 1 {
			fmt.Println(-1)
		} else {
			fmt.Println(N - minFreq + 2)
		}

		t--
	}
}
