package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	H := readNNums(reader, n)
	V := readNNums(reader, n)
	res := solve1(H, V)

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

func solve(H []int, V []int) int64 {
	n := len(H)

	type Pair struct {
		first  int
		second int
	}

	arr := make([]Pair, n*2)
	for i := n; i < 2*n; i++ {
		arr[i] = Pair{H[i-n], i - n}
	}
	for i := n - 1; i > 0; i-- {
		if arr[2*i].first < arr[2*i+1].first {
			arr[i] = arr[2*i]
		} else {
			arr[i] = arr[2*i+1]
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		res := Pair{1 << 30, -1}
		for l < r {
			if l&1 == 1 {
				if arr[l].first < res.first {
					res = arr[l]
				}
				l++
			}
			if r&1 == 1 {
				r--
				if arr[r].first < res.first {
					res = arr[r]
				}
			}
			l >>= 1
			r >>= 1
		}
		return res.second
	}

	var dfs func(l int, r int, t1 bool, t2 bool) int64

	dfs = func(l int, r int, t1 bool, t2 bool) int64 {
		if r <= l {
			return 0
		}
		pos := get(l, r)
		a := dfs(l, pos, t1, true)
		b := dfs(pos+1, r, true, t2)
		res := a + b + int64(V[pos])
		if t1 || t2 {
			res = max(res, 0)
		}
		if t1 {
			res = max(res, max(b+int64(V[pos]), b))
		}
		if t2 {
			res = max(res, max(a+int64(V[pos]), a))
		}
		return res
	}

	return dfs(0, n, false, false)
}

const inf = 1e15

func solve1(H []int, V []int) int64 {
	n := len(H)

	arr := make([]int64, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = -inf
	}

	set := func(p int, v int64) {
		p += n
		arr[p] = max(arr[p], v)
		for p > 1 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int64 {
		l += n
		r += n
		var res int64 = -inf
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}
	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		for p > 0 && H[stack[p-1]] > H[i] {
			p--
		}

		if p == 0 {
			tmp := max(0, get(0, i)) + int64(V[i])
			set(i, tmp)
		} else {
			j := stack[p-1]
			tmp := get(j, i) + int64(V[i])
			tmp = max(tmp, get(j, j+1))
			set(i, tmp)
		}

		stack[p] = i
		p++
	}

	return get(n-1, n)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
