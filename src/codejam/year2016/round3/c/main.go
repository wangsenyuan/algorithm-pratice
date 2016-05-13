package main

import "fmt"

func main() {
	var t = 0
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		var a, b, c, d int
		fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
		k := min(c, d)
		fmt.Printf("Case #%d: %d\n", i, a*b*k)
		result := play(a, b, c, d)
		for _, o := range result {
			fmt.Println(o)
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type outfit struct {
	_1, _2, _3 int
}

func (o outfit) String() string {
	return fmt.Sprintf("%d %d %d", o._1, o._2, o._3)
}

func play(a, b, c, d int) []outfit {
	k := min(c, d)
	result := make([]outfit, 0, a*b*k)
	for i := 1; i <= a; i++ {
		now := i
		for j := 0; j < b*k; j++ {
			result = append(result, outfit{i, j/k + 1, now})
			now++
			if now == c+1 {
				now = 1
			}
		}
	}
	return result
}
