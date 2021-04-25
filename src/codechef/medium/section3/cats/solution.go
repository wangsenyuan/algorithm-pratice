package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	P := readNNums(reader, n)
	C := readNNums(reader, n)
	res := solve(n, P, C)
	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const H = 20

func solve(n int, P []int, C []int) []int {
	pp := make([][]int, n)

	for i := 0; i < n; i++ {
		pp[i] = make([]int, H)
		for j := 0; j < H; j++ {
			pp[i][j] = -1
		}
		pp[i][0] = P[i] - 1
	}

	for j := 1; j < H; j++ {
		for i := 0; i < n; i++ {
			if pp[i][j-1] >= 0 {
				pp[i][j] = pp[pp[i][j-1]][j-1]
			}
		}
	}

	vis := make([]bool, n)

	res := make([]int, n)

	for i, c := range C {
		c--
		if vis[c] {
			res[i] = 0
			continue
		}
		var cnt = 1
		for j := H - 1; j >= 0; j-- {
			if pp[c][j] >= 0 && !vis[pp[c][j]] {
				cnt += 1 << uint(j)
				c = pp[c][j]
			}
		}
		vis[c] = true
		res[i] = cnt
	}

	return res
}
