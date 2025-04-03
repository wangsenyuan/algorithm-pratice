package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a)
	v := make([][]int, n)
	for i := range n {
		v[i] = make([]int, n)
		for j := range n {
			v[i][j] = int(a[i][j] - '0')
		}
	}

	row := make([]int, n)
	col := make([]int, n)

	var ans int
	clear(row)
	clear(col)

	for r := 0; r < n-1; r++ {
		for c := n - 1; c > r; c-- {
			cur := row[r] ^ col[c]
			if cur != v[r][c] {
				ans++
				row[r] ^= 1
				col[c] ^= 1
			}
		}
	}

	// 中轴线的需要特殊处理
	for r := 0; r < n; r++ {
		cur := row[r] ^ col[r]
		v[r][r] ^= cur
	}

	clear(row)
	clear(col)

	for r := n - 1; r >= 0; r-- {
		for c := 0; c <= r; c++ {
			cur := row[r] ^ col[c]
			if cur != v[r][c] {
				ans++
				row[r] ^= 1
				col[c] ^= 1
			}
		}
	}
	return ans
}
