package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)

	fmt.Println(solve(n, A))
}

const INF = math.MaxInt32 >> 1
const N_INF = -INF

func solve(n int, A []int) int {
	var ans int64
	for mx := int64(0); mx < 31; mx++ {
		var cur int64
		var best int64

		for i := 0; i < n; i++ {
			v := int64(A[i])
			if v > mx {
				v = N_INF
			}
			cur += v
			if cur > best {
				best = cur
			}
			if cur < 0 {
				cur = 0
			}
		}

		if best-mx > ans {
			ans = best - mx
		}
	}

	return int(ans)
}

func solve2(n int, A []int) int {
	var ans int64
	for mx := int64(0); mx < 31; mx++ {
		var cur int64
		var best int64

		for i := 0; i < n; i++ {
			v := int64(A[i])
			if v > mx {
				v = N_INF
			}
			cur += v
			if cur < best {
				best = cur
			}
			if cur-best-mx > ans {
				ans = cur - best - mx
			}
		}
	}

	return int(ans)
}

func solve1(n int, A []int) int {
	N := n + 1
	arr := make([]int, 2*N)

	reset := func() {
		for i := 0; i < len(arr); i++ {
			arr[i] = INF
		}
	}

	reset()

	update := func(p int, v int) {
		p += N
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		var res = INF

		l += N
		r += N

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

	stack := make([]int, n)
	var p int

	left := make([]int, n)
	var sum int
	update(0, sum)
	for i := 0; i < n; i++ {
		for p > 0 && A[stack[p-1]] <= A[i] {
			p--
		}
		var j = 0
		if p > 0 {
			j = stack[p-1] + 1
		}
		sum += A[i]
		update(i+1, sum)
		left[i] = sum - get(j, i+1)
		stack[p] = i
		p++
	}

	reset()

	sum = 0
	p = 0
	var best = 0
	update(0, sum)
	// i + j == n
	for i := n - 1; i >= 0; i-- {
		for p > 0 && A[stack[p-1]] <= A[i] {
			p--
		}

		sum += A[i]
		update(n-i, sum)

		var j = 0
		if p > 0 {
			j = n - stack[p-1] + 1
		}
		tmp := sum - get(j, n-i)

		if tmp+left[i]-2*A[i] > best {
			best = tmp + left[i] - 2*A[i]
		}
		stack[p] = i
		p++
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
