package main

import "bytes"

func minInteger(num string, k int) string {
	n := len(num)
	cnt := make([]int, 10)
	for i := 0; i < n; i++ {
		x := int(num[i] - '0')
		cnt[x]++
	}
	pos := make([][]int, 10)
	for i := 0; i < 10; i++ {
		pos[i] = make([]int, 0, cnt[i])
	}

	for i := 0; i < n; i++ {
		x := int(num[i] - '0')
		pos[x] = append(pos[x], i)
	}

	arr := make([]int, n+1)

	update := func(p int, v int) {
		p++
		for p <= n {
			arr[p] += v
			p += p & -p
		}
	}

	get := func(p int) int {
		var res int
		p++
		for p > 0 {
			res += arr[p]
			p -= p & -p
		}
		return res
	}

	for i := 0; i < n; i++ {
		update(i, 1)
	}

	var buf bytes.Buffer
	var cur int
	ii := make([]int, 10)

	for cur < n {

		for i := 0; i < 10; i++ {
			if ii[i] == len(pos[i]) {
				continue
			}
			x := pos[i][ii[i]]
			y := get(x - 1)
			if y <= k {
				ii[i]++
				k -= y
				update(x, -1)
				buf.WriteByte(byte(i + '0'))
				break
			}
		}

		cur++
	}

	return buf.String()
}
