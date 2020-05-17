package p1446

import "fmt"

func simplifiedFractions(n int) []string {
	res := make([]string, 0, 100)
	mem := make(map[Pair]bool)

	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			x := gcd(i, j)
			p := Pair{i / x, j / x}
			if !mem[p] {
				res = append(res, p.String())
				mem[p] = true
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Pair struct {
	first, second int
}

func (pair Pair) String() string {
	return fmt.Sprintf("%d/%d", pair.first, pair.second)
}
