package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(v int) string {
		fmt.Printf("? %d\n", v)
		s, _ := reader.ReadString('\n')
		return s[:1]
	}

	n := readNum(reader)

	res := solve(n, ask)

	var buf bytes.Buffer
	buf.WriteString("!\n")
	for i := 0; i < n; i++ {
		buf.WriteString(res[i])
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

func solve(n int, ask func(int) string) []string {
	if n == 1 {
		return []string{"0"}
	}

	Q := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		Q[i] = make([]int, 0, n+1)
	}
	for j := 1; j <= n; j++ {
		res := ask(j)
		if res == "B" {
			Q[j] = append(Q[j], 1)
		} else {
			Q[j] = append(Q[j], 0)
		}
	}
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			res := ask(j)
			if res == "B" {
				Q[j] = append(Q[j], 1)
			} else {
				Q[j] = append(Q[j], 0)
			}
		}
	}

	grid := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		grid[i] = make([]int, n+1)
	}

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			grid[i][j] = Q[i][j-1] ^ Q[i][j+1]
			grid[j][i] = grid[i][j]
		}
	}

	res := make([]string, n)
	for i := 1; i <= n; i++ {
		buf := make([]byte, n)
		for j := 1; j <= n; j++ {
			buf[j-1] = byte('0' + grid[i][j])
		}
		res[i-1] = string(buf)
	}

	return res
}
