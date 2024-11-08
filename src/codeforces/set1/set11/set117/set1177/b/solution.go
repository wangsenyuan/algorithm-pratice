package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	fmt.Scan(&k)

	res := solve(k)

	fmt.Println(res)
}

func solve(k int) int {
	if k < 10 {
		return k
	}
	count := func(num int) int {
		// 包括num, 有多少个digit
		var cnt int
		if num < 10 {
			cnt += num
			return cnt
		}
		cnt += 9
		// num >= 10
		base := 100
		prev := 9
		d := 2
		for num >= base {
			cnt += (base - 1 - prev) * d
			prev = base - 1
			base *= 10
			d++
		}
		// base > num
		if num*10 != base {
			cnt += (num - prev) * d
		} else {
			// 如果num = 10， 那么这个是没有计算的
			cnt += d
		}
		return cnt
	}

	num := sort.Search(k, func(i int) bool {
		return count(i) >= k
	})

	prev := count(num - 1)
	rem := k - prev
	s := fmt.Sprintf("%d", num)
	return int(s[rem-1] - '0')
}
