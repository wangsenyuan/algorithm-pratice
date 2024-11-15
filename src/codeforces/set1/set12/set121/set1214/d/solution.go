package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	mat := make([]string, n)
	for i := range n {
		mat[i] = readString(reader)[:m]
	}
	return solve(mat)
}
func solve(mat []string) int {
	n := len(mat)
	m := len(mat[0])

	que := make([]int, n*m)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}

	var head, tail int
	vis[n-1][m-1] = true
	que[head] = n*m - 1
	head++

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		if r > 0 && mat[r-1][c] == '.' && !vis[r-1][c] {
			vis[r-1][c] = true
			que[head] = (r-1)*m + c
			head++
		}
		if c > 0 && mat[r][c-1] == '.' && !vis[r][c-1] {
			vis[r][c-1] = true
			que[head] = r*m + c - 1
			head++
		}
	}

	if !vis[0][0] {
		return 0
	}

	head = 0
	tail = 0
	que[head] = 0
	head++

	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}

	marked[0][0] = true

	for tail < head {
		if tail+1 == head && que[tail] != 0 && que[tail] != n*m-1 {
			return 1
		}
		// move to next level
		pos := head
		for tail < pos {
			r, c := que[tail]/m, que[tail]%m
			tail++
			if c+1 < m && vis[r][c+1] && !marked[r][c+1] {
				que[head] = r*m + c + 1
				head++
				marked[r][c+1] = true
			}
			if r+1 < n && vis[r+1][c] && !marked[r+1][c] {
				que[head] = (r+1)*m + c
				head++
				marked[r+1][c] = true
			}
		}
	}

	return 2
}
