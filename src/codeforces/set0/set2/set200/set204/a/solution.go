package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	res := solve(l, r)
	fmt.Println(res)
}

func solve(l int, r int) int {

	get := func(num int) int {
		if num < 10 {
			return num
		}
		res := 9 + num/10
		x := num % 10

		for num >= 10 {
			num /= 10
		}

		if num > x {
			res--
		}

		return res
	}

	return get(r) - get(l-1)
}

func solve1(l int, r int) int {
	sr := fmt.Sprintf("%d", r)
	n := len(sr)

	pw := make([]int, n+1)

	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = pw[i-1] * 10
	}

	// 只需要首位和末尾相同
	f := make([]int, n+1)
	// 只有1位的时候
	f[1] = 10
	for i := 2; i <= n; i++ {
		// 只需要确定前半部分
		f[i] = 9*pw[i-2] + f[i-1]
	}

	// 找出比 <= num 的good的数目
	get := func(num int) int {
		if num < 10 {
			// 0, 1, ... num
			return num + 1
		}
		s := fmt.Sprintf("%d", num)
		m := len(s)
		res := f[m-1]

		x := int(s[0] - '0')

		res += (x - 1) * pw[m-2]

		for i := 1; i < m-1; i++ {
			x := int(s[i] - '0')
			// 最后一位是确定的，当前位取比x小的数(0, 1, .. x- 1)
			res += x * pw[m-2-i]
		}

		if s[0] <= s[m-1] {
			// num 也符合条件
			res++
		}

		return res
	}

	return get(r) - get(l-1)
}
