package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(k, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const INF = 1 << 60

func solve(k int, A []int) int64 {
	n := len(A)
	left := make([]int, n)
	right := make([]int, n)
	size := make([]int, n)
	var build func(l, r int) int

	build = func(l, r int) int {
		if l == r {
			left[l] = -1
			right[l] = -1
			size[l]++
			return l
		}
		var pivot = l
		for i := l; i <= r; i++ {
			if A[i] > A[pivot] {
				pivot = i
			}
		}
		if pivot == l {
			left[l] = -1
			right[l] = build(l+1, r)
			size[l] = size[right[l]] + 1
			return l
		}
		if pivot == r {
			left[r] = build(l, r-1)
			right[r] = -1
			size[r] = size[left[r]] + 1
			return r
		}

		left[pivot] = build(l, pivot-1)
		right[pivot] = build(pivot+1, r)
		size[pivot] = size[left[pivot]] + size[right[pivot]] + 1
		return pivot
	}

	root := build(0, n-1)

	best := make([][]int64, n)

	for i := 0; i < n; i++ {
		best[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			best[i][j] = INF
		}
	}

	var compute func(u int)

	compute = func(u int) {
		if left[u] != -1 {
			compute(left[u])
		}
		if right[u] != -1 {
			compute(right[u])
		}

		if left[u] == -1 {
			if right[u] == -1 {
				best[u][0] = 0
				best[u][1] = int64(A[u])
			} else {
				for x := 0; x <= size[right[u]]; x++ {
					best[u][x] = min(best[u][x], best[right[u]][x])
					if x+1 <= n {
						best[u][x+1] = min(best[u][x+1], best[right[u]][x]+int64(x+1)*int64(A[u]))
					}
				}
			}
		} else {
			if right[u] == -1 {
				for x := 0; x <= size[left[u]]; x++ {
					best[u][x] = min(best[u][x], best[left[u]][x])
					if x+1 <= n {
						best[u][x+1] = min(best[u][x+1], best[left[u]][x]+int64(x+1)*int64(A[u]))
					}
				}
			} else {
				for x := 0; x <= size[left[u]]; x++ {
					for y := 0; y <= size[right[u]]; y++ {
						best[u][x+y] = min(best[u][x+y], best[left[u]][x]+best[right[u]][y]+int64(x)*int64(y)*int64(A[u]))
						best[u][x+y+1] = min(best[u][x+y+1], best[left[u]][x]+best[right[u]][y]+int64(1+x)*int64(1+y)*int64(A[u]))
					}
				}
			}
		}
	}

	compute(root)

	return best[root][k]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
