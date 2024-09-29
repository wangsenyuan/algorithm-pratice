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
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, n)
	}

	k := readNum(reader)
	roads := make([][]int, k)
	for i := 0; i < k; i++ {
		roads[i] = readNNums(reader, 3)
	}

	res := solve(mat, roads)

	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d ", cur))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const inf = 1 << 60

func solve(mat [][]int, roads [][]int) []int {
	n := len(mat)

	ans := make([]int, len(roads))

	for i, cur := range roads {
		a, b, c := cur[0], cur[1], cur[2]
		a--
		b--
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				mat[u][v] = min(mat[u][v], mat[u][a]+mat[b][v]+c, mat[u][b]+mat[a][v]+c)
			}
		}
		for u := 0; u < n; u++ {
			for v := u + 1; v < n; v++ {
				ans[i] += mat[u][v]
			}
		}
	}

	return ans
}
