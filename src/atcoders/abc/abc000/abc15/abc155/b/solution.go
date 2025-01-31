package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readNInt64s(reader, 2)
	n := int(first[0])
	k := first[1]
	a := readNInt64s(reader, n)
	res := solve(k, a)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

const inf = 1 << 61

func solve(k int64, a []int64) int64 {
	n := len(a)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	p0 := -1
	for i := 0; i < n; i++ {
		if a[i] >= 0 {
			p0 = i
			break
		}
	}

	count := func(target int64) int64 {
		var cnt int64
		var r3 int
		r1 := n - 1
		r2 := n - 1
		for i := 0; i < n; i++ {
			x := a[i]
			// x * y <= target
			if target < 0 {
				if x >= 0 {
					break
				}
				// x < 0, find all > 0
				//cnt += int64(n - pos[1])
				// 绝对值越大的负数 * 绝对值越大的正数 => 越小
				// r3有问题，是因为随着绝对值变小，r3其实应该是要减少的
				// -3, -2,  1, 2
				// 对于负数来说
				for r3 < 0 && a[r3] <= 0 {
					r3++
				}
				// a, b, c, d
				// 如果 a * c > target => b * c > target
				for r3 < n && x*a[r3] > target {
					r3++
				}

				cnt += int64(n - r3)
			} else if target == 0 {
				if x > 0 {
					break
				}
				if x == 0 {
					cnt += int64(n - (i + 1))
				} else {
					// x < 0, find all >= 0
					cnt += int64(n - max(i+1, p0))
				}
			} else {
				// target > 0
				if x == 0 {
					cnt += int64(n - (i + 1))
				} else if x < 0 {
					// those will always be true
					for r1 > i && a[r1] >= 0 {
						r1--
					}
					for r1 > i && x*a[r1] <= target {
						r1--
					}
					cnt += int64(n - max(r1+1, i+1))
				} else {
					// a, b, c, d
					// a * d <= target => b * d <= target it not true
					// 但是 a * d >= target => b * d > target 是成立的
					for r2 > i && x*a[r2] > target {
						r2--
					}
					cnt += int64(max(0, r2-i))
				}
			}
		}

		return cnt
	}

	var l, r int64 = -inf, inf
	var ans int64
	for r >= l {
		mid := (l + r) / 2
		if count(mid) >= k {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
