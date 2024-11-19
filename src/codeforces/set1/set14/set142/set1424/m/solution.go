package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
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

func process(reader *bufio.Reader) (string, []string, int) {
	n, k := readTwoNums(reader)
	pages := make([]string, n*(k+1))
	for i := range len(pages) {
		pages[i] = readString(reader)
	}
	return solve(pages, k), pages, k
}

func solve(input []string, k int) string {
	book := readAsBook(input, k)

	ord := make([][]int, 26)
	for i := 0; i < 26; i++ {
		ord[i] = make([]int, 26)
	}

	var dfs func(lines []string, l int, r int) bool

	dfs = func(lines []string, l int, r int) bool {
		if l == r {
			return true
		}
		mid := (l + r) / 2
		// 左边的都要小于右边的
		if !dfs(lines, l, mid) {
			return false
		}
		if !dfs(lines, mid+1, r) {
			return false
		}
		a := lines[mid]
		b := lines[mid+1]
		var u int
		for u < len(a) && u < len(b) && a[u] == b[u] {
			u++
		}
		if u == len(a) {
			return true
		}
		if u == len(b) {
			return false
		}
		x := int(a[u] - 'a')
		y := int(b[u] - 'a')
		if ord[x][y] == 0 {
			ord[x][y] = -1
			ord[y][x] = 1
			return true
		}
		return ord[x][y] == -1
	}
	arr := make([]string, len(book)*k)
	for i, page := range book {
		copy(arr[i*k:(i+1)*k], page.lines)
	}

	if !dfs(arr, 0, len(arr)-1) {
		return "IMPOSSIBLE"
	}

	cnt := make([]int, 26)
	for _, s := range arr {
		for _, x := range []byte(s) {
			cnt[int(x-'a')]++
		}
	}

	// now it a dag, just topo sort
	deg := make([]int, 26)

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if ord[i][j] < 0 {
				// i is before j
				deg[j]++
			}
		}
	}

	que := make([]int, 26)
	var head, tail int

	var buf []byte

	for i := 0; i < 26; i++ {
		if cnt[i] > 0 && deg[i] == 0 {
			que[head] = i
			head++
			buf = append(buf, byte(i+'a'))
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for v := 0; v < 26; v++ {
			if ord[u][v] < 0 {
				deg[v]--
				if deg[v] == 0 {
					que[head] = v
					head++
					buf = append(buf, byte(v+'a'))
				}
			}
		}
	}

	for i := 0; i < 26; i++ {
		if cnt[i] > 0 && deg[i] > 0 {
			return "IMPOSSIBLE"
		}
	}

	return string(buf)
}

func readAsBook(input []string, k int) []Page {
	var book []Page

	for i := 0; i < len(input); i += k + 1 {
		id, _ := strconv.Atoi(input[i])
		lines := input[i+1 : i+k+1]
		book = append(book, Page{id, lines})
	}
	slices.SortFunc(book, func(a, b Page) int {
		return a.id - b.id
	})
	return book
}

type Page struct {
	id    int
	lines []string
}
