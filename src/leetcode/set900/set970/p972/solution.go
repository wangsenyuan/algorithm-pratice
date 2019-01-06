package p972

import (
	"regexp"
	"strconv"
)

func isRationalEqual(S string, T string) bool {
	a := ParseAsFrac(S)
	b := ParseAsFrac(T)
	return a.n == b.n && a.d == b.d
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Frac struct {
	n int64
	d int64
}

func NewFrac(n, d int64) Frac {
	g := gcd(n, d)
	return Frac{n / g, d / g}
}

func (this Frac) Add(other Frac) Frac {
	numerator := this.n*other.d + this.d*other.n
	denominator := this.d * other.d
	g := gcd(numerator, denominator)
	return Frac{numerator / g, denominator / g}
}

func ParseAsFrac(S string) Frac {
	var state int

	regex := regexp.MustCompile("[.()]")
	ss := regex.Split(S, -1)

	ans := NewFrac(0, 1)
	var decimalSize int

	for _, s := range ss {
		state++
		if len(s) == 0 {
			continue
		}
		x, _ := strconv.Atoi(s)
		sz := len(s)
		if state == 1 {
			ans = ans.Add(NewFrac(int64(x), 1))
		} else if state == 2 {
			ans = ans.Add(NewFrac(int64(x), pow(10, sz)))
			decimalSize = sz
		} else {
			deno := pow(10, decimalSize)
			deno *= pow(10, sz) - 1
			ans = ans.Add(NewFrac(int64(x), deno))
		}
	}

	return ans
}

func pow(x int64, n int) int64 {
	y := int64(1)
	for n > 0 {
		if n&1 == 1 {
			y = y * x
		}
		x *= x
		n >>= 1
	}
	return y
}
