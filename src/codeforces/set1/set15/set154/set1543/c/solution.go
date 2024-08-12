package main

import (
	"bytes"
	"fmt"
)

func main() {

	var tc int
	fmt.Scan(&tc)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		var c, m, p, v float64
		fmt.Scan(&c, &m, &p, &v)

		res := solve(c, m, p, v)

		buf.WriteString(fmt.Sprintf("%.10f\n", res))
	}

	fmt.Print(buf.String())
}

const scale = 1e7
const eps = 1e-9

func solve(c float64, m float64, p float64, v float64) float64 {

	var play func(c int, m int, p int, v int) float64

	play = func(c int, m int, p int, v int) float64 {
		var ans = float64(p) / scale

		if c > 0 {
			if c > v {
				if m > 0 {
					// 可以分配过去
					ans += float64(c) / scale * (1 + play(c-v, m+v/2, p+v/2, v))
				} else {
					ans += float64(c) / scale * (1 + play(c-v, 0, p+v, v))
				}
			} else {
				if m > 0 {
					ans += float64(c) / scale * (1 + play(0, m+c/2, p+c/2, v))
				} else {
					ans += float64(c) / scale * (1 + play(0, 0, p+c, v))
				}
			}
		}

		if m > 0 {
			if m > v {
				if c > 0 {
					ans += float64(m) / scale * (1 + play(c+v/2, m-v, p+v/2, v))
				} else {
					ans += float64(m) / scale * (1 + play(0, m-v, p+v, v))
				}
			} else {
				if c > 0 {
					ans += float64(m) / scale * (1 + play(c+m/2, 0, p+m/2, v))
				} else {
					ans += float64(m) / scale * (1 + play(0, 0, p+m, v))
				}
			}
		}

		return ans
	}

	return play(int(c*scale), int(m*scale), int(p*scale), int(v*scale))
}
