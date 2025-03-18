package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		for _, cur := range res {
			fmt.Println(cur[0], cur[1])
		}
	}
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

func process(reader *bufio.Reader) [][]int {
	n, m, k := readThreeNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)[:m]
	}
	return solve(k, a)
}

func solve(k int, a []string) [][]int {
	n := len(a)
	m := len(a[0])

	check := func(i, j, x int) bool {
		if a[i][j] != '*' {
			return false
		}
		if a[i-x][j] != '*' {
			return false
		}
		if a[i+x][j] != '*' {
			return false
		}
		if a[i][j-x] != '*' {
			return false
		}
		if a[i][j+x] != '*' {
			return false
		}
		return true
	}

	for x := 1; x <= min(n/2, m/2); x++ {
		for i := x; i+x < n; i++ {
			for j := x; j+x < m; j++ {
				if check(i, j, x) {
					k--
					if k == 0 {
						return [][]int{{i + 1, j + 1}, {i - x + 1, j + 1}, {i + x + 1, j + 1}, {i + 1, j - x + 1}, {i + 1, j + x + 1}}
					}
				}
			}
		}
	}
	return nil
}
