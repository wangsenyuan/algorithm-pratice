package p3149

import "math/bits"

func findProductsOfElements(queries [][]int64) []int {
	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r, m := cur[0], cur[1], cur[2]
		el := calc(int(l))
		er := calc(int(r + 1))
		ans[i] = pow(2, er-el, int(m))
	}

	return ans
}

func pow(a, b int, mod int) int {
	r := 1 % mod
	for b > 0 {
		if b&1 == 1 {
			r = r * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return r
}

func calc(k int) (res int) {
	var n, cnt, sum int
	for i := bits.Len(uint(k+1)) - 1; i > 0; i-- {
		// 前面1个数量，重复 1 << i次
		// 后面1的个数, i * (1 << (i - 1))
		c := cnt<<i + i<<(i-1)
		if c <= k {
			// 这一位可以放置1
			k -= c

			res += sum<<i + i*(i-1)/2<<(i-1)

			sum += i
			cnt++

			n |= 1 << i
		}
	}

	if cnt <= k {
		k -= cnt
		res += sum
		n++
	}

	for k > 0 {
		res += bits.TrailingZeros(uint(n))
		n &= n - 1
		k--
	}

	return
}
