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

	for tc > 0 {
		tc--
		n := readNum(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(grid []string) int {
	n := len(grid)

	var res int

	cnt := make([]int, 2)

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			// (i, j) => (u, v)

			for k := 0; k < 2; k++ {
				cnt[k] = 0
			}
			u, v := i, j
			cnt[int(grid[u][v]-'0')]++
			u, v = v, n-1-u
			cnt[int(grid[u][v]-'0')]++
			u, v = v, n-1-u
			cnt[int(grid[u][v]-'0')]++
			u, v = n-1-j, i
			cnt[int(grid[u][v]-'0')]++
			res += 4 - max(cnt[0], cnt[1])
		}
	}

	if n&1 == 1 {
		for r := 0; r < n/2; r++ {
			cnt[0] = 0
			cnt[1] = 0
			cnt[int(grid[r][n/2]-'0')]++
			cnt[int(grid[n-1-r][n/2]-'0')]++
			cnt[int(grid[n/2][r]-'0')]++
			cnt[int(grid[n/2][n-1-r]-'0')]++
			res += 4 - max(cnt[0], cnt[1])
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
