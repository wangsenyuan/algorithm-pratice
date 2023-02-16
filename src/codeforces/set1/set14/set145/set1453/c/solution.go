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
		G := make([]string, n)

		for i := 0; i < n; i++ {
			G[i], _ = reader.ReadString('\n')
		}

		res := solve(n, G)

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, G []string) []int {
	D := make([][]int, 10)
	for i := 0; i < 10; i++ {
		D[i] = make([]int, 4)
		//top
		D[i][0] = n + 1
		//right
		D[i][1] = -1
		// bottom
		D[i][2] = -1
		// left
		D[i][3] = n + 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := int(G[i][j] - '0')
			D[x][0] = min(D[x][0], i)
			D[x][1] = max(D[x][1], j)
			D[x][2] = max(D[x][2], i)
			D[x][3] = min(D[x][3], j)
		}
	}

	res := make([]int, 10)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := int(G[i][j] - '0')
			// if put (i, j) as an endpoint, parallel to row, another endpoint is 0 or (n - 1)
			a := max(j, n-j-1)
			b := max(i-D[x][0], D[x][2]-i)
			res[x] = max(res[x], a*b)
			a = max(i, n-i-1)
			b = max(j-D[x][3], D[x][1]-j)
			res[x] = max(res[x], a*b)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
