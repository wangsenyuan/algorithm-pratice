package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	fmt.Println(res)
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

func solve(a []int) string {
	n := len(a)

	l, r := 0, n-1

	for l+1 < n && a[l+1] > a[l] {
		l++
	}

	for r-1 >= 0 && a[r-1] > a[r] {
		r--
	}

	// pos是当前可以处理的位置
	var dfs func(players []string, pos []int, dest []int) string

	dfs = func(players []string, pos []int, dest []int) string {
		if pos[0] == dest[0]+1 || pos[1] == dest[1]-1 {
			// 只能处理另外一段
			ln := (pos[1] - dest[1] + 1) | (dest[0] - pos[0] + 1)

			return players[1-(ln&1)]
		}
		if a[pos[0]] <= a[pos[1]] {
			ln := pos[1] - dest[1] + 1
			if ln&1 == 1 {
				// 从右端开始取，当前玩家可以获胜
				return players[0]
			}
			// 只能取最左端的
			pos[0]++
			return dfs(reverse(players), pos, dest)
		}
		// a[pos[0]] >= a[pos[1]]
		ln := dest[0] - pos[0] + 1
		if ln&1 == 1 {
			return players[0]
		}
		pos[1]--
		return dfs(reverse(players), pos, dest)
	}

	return dfs([]string{"Alice", "Bob"}, []int{0, n - 1}, []int{l, r})
}

func reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
