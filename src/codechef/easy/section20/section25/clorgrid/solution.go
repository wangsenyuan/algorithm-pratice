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
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)
		}
		res := solve(n, m, grid)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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
func solve(n int, m int, grid []string) bool {
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
	}

	canColor := func(i, j int) bool {
		if i+2 > n || j+2 > m {
			// no room
			return false
		}
		for x := 0; x < 2; x++ {
			for y := 0; y < 2; y++ {
				if grid[i+x][j+y] == '#' {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' && canColor(i, j) {
				cnt[i][j]++
				if i+2 < n {
					cnt[i+2][j]--
				}
				if j+2 < m {
					cnt[i][j+2]--
				}
				if i+2 < n && j+2 < m {
					cnt[i+2][j+2]++
				}
			}
		}
	}
	row := make([]int, m)

	for i := 0; i < n; i++ {
		var cur int
		for j := 0; j < m; j++ {
			cur += cnt[i][j] + row[j]
			if grid[i][j] == '.' && cur == 0 {
				return false
			}
			row[j] += cnt[i][j]
		}
	}

	return true
}
