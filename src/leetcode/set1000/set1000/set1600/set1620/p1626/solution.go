package p1626

import (
	"bytes"
	"fmt"
)

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)

	arr := make([]int, 2*n)

	for i := 0; i < n; i++ {
		arr[i] = int(s[i] - '0')
		arr[i+n] = arr[i]
	}

	best := make([]int, n)

	copy(best, arr[:n])

	g := gcd(n, b)

	tmp := make([]int, n)

	for i := 0; i < n; i += g {
		for j := 0; j < 10; j++ {
			var th int
			if g%2 == 1 {
				th = 9
			}
			for k := 0; k <= th; k++ {
				copy(tmp, arr[i:i+n])

				for t := 1; t < n; t += 2 {
					tmp[t] += a * j
					tmp[t] %= 10
				}
				for t := 0; t < n; t += 2 {
					tmp[t] += a * k
					tmp[t] %= 10
				}
				var p int
				for p < n && tmp[p] == best[p] {
					p++
				}
				if p < n && tmp[p] < best[p] {
					copy(best, tmp)
				}
			}
		}
	}

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d", best[i]))
	}

	return buf.String()
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
