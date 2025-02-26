package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	qs := make([]int, m)
	for i := range m {
		qs[i] = readNum(reader)
	}
	return solve(a, qs)
}

const H = 30

func solve(a []int, qs []int) []int {
	n := len(a)
	prev := make([][]int, n+1)
	for i := range n + 1 {
		prev[i] = make([]int, H)
		for j := 0; j < H; j++ {
			prev[i][j] = -1
		}
	}

	for i, x := range a {
		if i > 0 {
			copy(prev[i], prev[i-1])
		}
		for j := range H {
			if (x>>j)&1 == 1 {
				prev[i][j] = i
			}
		}
	}

	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = a[i] ^ sum[i+1]
	}

	find := func(x int) int {
		l := 0
		pos := n
		for d := H - 1; d >= 0 && l < pos; d-- {
			j := prev[pos-1][d]
			if (x>>d)&1 == 1 {
				// d是x的最高位
				if j < l {
					// 可以到达位置l
					break
				}
				// j > l
				y := x ^ sum[j+1] ^ sum[pos]
				if y < a[j] {
					// 因为后面的d位都为0，都小于x
					l = j + 1
					break
				}
				// y >= a[j]
				pos = j
				x = y ^ a[j]
				if pos > 0 {
					l = max(l, prev[pos-1][d]+1)
				}
			} else {
				l = max(l, j+1)
			}
		}
		return n - l
	}

	ans := make([]int, len(qs))

	for i, x := range qs {
		ans[i] = find(x)
	}

	return ans
}
