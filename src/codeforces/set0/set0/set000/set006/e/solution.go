package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	H := readNNums(reader, n)

	a, b := solve(n, k, H)
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d %d\n", a, len(b)))

	for i := 0; i < len(b); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", b[i][0], b[i][1]))
	}

	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const N_INF = -(10000000)

func solve(n, k int, H []int) (a int, b [][]int) {
	get := func(l, r int, arr []int) int {
		res := N_INF

		l += n
		r += n

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

	x := make([]int, 2*n)
	y := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		x[i] = H[i-n]
		y[i] = -H[i-n]
	}

	for i := n - 1; i > 0; i-- {
		x[i] = max(x[2*i], x[2*i+1])
		y[i] = max(y[2*i], y[2*i+1])
	}

	getDiff := func(i, j int) int {
		xx := get(i, j, x)
		yy := -get(i, j, y)
		return xx - yy
	}

	for i, j := 0, 1; j <= n; j++ {
		for i < j && getDiff(i, j) > k {
			i++
		}
		if j-i > a {
			a = j - i
		}
	}

	for i := 0; i+a <= n; i++ {
		d := getDiff(i, i+a)
		if d <= k {
			b = append(b, []int{i + 1, i + a})
		}
	}
	return
}

func solve1(n, k int, H []int) (a int, b [][]int) {
	get := func(l, r int, arr []int) int {
		res := N_INF

		l += n
		r += n

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

	x := make([]int, 2*n)
	y := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		x[i] = H[i-n]
		y[i] = -H[i-n]
	}

	for i := n - 1; i > 0; i-- {
		x[i] = max(x[2*i], x[2*i+1])
		y[i] = max(y[2*i], y[2*i+1])
	}

	getDiff := func(i, j int) int {
		xx := get(i, j, x)
		yy := -get(i, j, y)
		return xx - yy
	}

	check := func(l int) bool {
		for i := 0; i+l <= n; i++ {
			d := getDiff(i, i+l)
			if d <= k {
				return true
			}
		}
		return false
	}

	var left, right = 0, n + 1

	for left < right {
		mid := (left + right) / 2
		if !check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	a = right - 1

	for i := 0; i+a <= n; i++ {
		d := getDiff(i, i+a)
		if d <= k {
			b = append(b, []int{i + 1, i + a})
		}
	}
	return
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
