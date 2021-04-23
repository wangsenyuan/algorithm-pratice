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
		n := readNum(reader)
		G := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			G[i] = readNNums(reader, n-1)
		}
		R := solve(n, G)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				buf.WriteString(fmt.Sprintf("%d ", R[i][j]))
			}
			buf.WriteByte('\n')
		}
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

func solve(n int, G [][]int) [][]int {
	R := make([][]int, n)
	for i := 0; i < n; i++ {
		R[i] = make([]int, n)
	}

	pos := make([][]int, n*n+1)
	for i := 1; i <= n*n; i++ {
		pos[i] = make([]int, 0, 2)
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			pos[G[i][j]] = append(pos[G[i][j]], i*n+j)
		}
	}
	x := make([]int, 0, n*n)
	y := make([]int, 0, n*n)

	for i := 1; i <= n*n; i++ {
		if len(pos[i]) > 0 {
			x = append(x, i)
		} else {
			y = append(y, i)
		}
	}

	var it int

	place := func(num int) {
		var i, j int
		for _, cur := range pos[num] {
			i = max(i, cur/n)
			j = max(j, cur%n)
		}
		if R[i][j] != 0 {
			if len(pos[num]) == 1 {
				if R[i+1][j] == 0 {
					i++
				} else if R[i][j+1] == 0 {
					j++
				} else {
					i++
					j++
				}
			} else {
				tmp := pos[num][0]
				i = tmp/n + 1
				j = tmp%n + 1
			}
		}
		R[i][j] = num
		for _, p := range pos[num] {
			f, s := p/n, p%n
			for ii := f; ii < f+2; ii++ {
				for jj := s; jj < s+2; jj++ {
					if ii < n && jj < n && R[ii][jj] == 0 {
						R[ii][jj] = y[it]
						it++
					}
				}
			}
		}
	}

	for i := 0; i < len(x); i++ {
		place(x[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if R[i][j] == 0 {
				R[i][j] = y[it]
				it++
			}
		}
	}

	return R
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
