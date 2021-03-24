package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	A := readNNums(reader, n)

	res := solve1(n, k, A)

	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve1(n int, k int, A []int) int {
	B := make([]int, n)

	check := func(cand int) bool {
		for i := 0; i < n; i++ {
			if A[i] >= cand {
				B[i] = 1
			} else {
				B[i] = -1
			}
		}

		for i := 1; i < n; i++ {
			B[i] += B[i-1]
		}

		mx := B[k-1]
		mn := 0

		for i := k; i < n; i++ {
			mn = min(mn, B[i-k])
			mx = max(mx, B[i]-mn)
		}

		return mx > 0
	}

	var left, right = INF, 1

	for i := 0; i < n; i++ {
		if left > A[i] {
			left = A[i]
		}
		if right < A[i] {
			right = A[i]
		}
	}
	if left == right {
		return right
	}
	right++
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

const INF = 1 << 29

func solve(n int, k int, A []int) int {
	//如果存在某个区间[l...r] (r - l + 1 >= k), 它的median <= x
	//那么设ge[l...r]为大于等于x的数量, lt[l....r] 为 < x的数量
	//则ge >= lt 当 sz(S) 为偶数 ge+1>=lt when odd
	// ge(R) - ge(L) >= lt(R) - lt(L)
	// let ge(R) = x, then lt(R) = R + 1 - x
	// let ge(L) = y, then lt(L) = L + 1 - y
	// x - y >= R - L - (x - y)
	// 2x - R >= 2 * y - L

	arr := make([]int, 2*n)

	update := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		l += n
		r += n
		res := INF

		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	check := func(cand int) bool {
		for i := 0; i < len(arr); i++ {
			arr[i] = INF
		}
		var x int
		for i := 0; i < n; i++ {
			if A[i] >= cand {
				x++
			}
			cur := 2*x - (i + 1)
			if cur > 0 && i+1 >= k {
				return true
			}
			if i+1 > k {
				tmp := get(0, i+1-k)
				// [1, 1, 1, 2, 2, 2] suppose cand = 2, and k = 3, and i = 4, then tmp should get from [0...2)
				if cur > tmp {
					return true
				}
			}
			update(i, cur)
		}
		return false
	}

	var left, right = INF, 1

	for i := 0; i < n; i++ {
		if left > A[i] {
			left = A[i]
		}
		if right < A[i] {
			right = A[i]
		}
	}
	if left == right {
		return right
	}
	right++
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
