package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	s := readNNums(reader, k)

	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}

	res := solve(mat, s)
	var buf bytes.Buffer

	for i := 0; i < k; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(mat []string, s []int) []int {
	p := len(s)
	n := len(mat)
	m := len(mat[0])

	// . for empty, # for block, number for start position
	que := make([][]Item, 1)

	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
		for j := 0; j < m; j++ {
			vis[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == '.' || mat[i][j] == '#' {
				if mat[i][j] == '#' {
					vis[i][j] = 10
				}
				continue
			}
			x := int(mat[i][j] - '1')
			it := Item{x, i, j}
			que[0] = append(que[0], it)
			vis[i][j] = x
		}
	}

	sort.Slice(que[0], func(i, j int) bool {
		return que[0][i].from < que[0][j].from
	})

	pq := make([]int, n*m)
	dist := make([]int, n*m)

	for i := 0; i < len(dist); i++ {
		dist[i] = -1
	}

	expand := func(x int, active []Item) []Item {
		if len(active) == 0 {
			return nil
		}
		var front, tail int
		for _, it := range active {
			u, v := it.x, it.y
			pq[front] = u*m + v
			front++
			dist[u*m+v] = 0
		}

		var res []Item

		for tail < front {
			pos := pq[tail]
			u, v := pos/m, pos%m

			if tail >= len(active) {
				// 这部分是新增加的
				res = append(res, Item{x, u, v})
			}

			tail++

			if dist[pos] == s[x] {
				continue
			}

			for i := 0; i < 4; i++ {
				a, b := u+dd[i], v+dd[i+1]
				if a >= 0 && a < n && b >= 0 && b < m && dist[a*m+b] < 0 && vis[a][b] < 0 {
					vis[a][b] = x
					dist[a*m+b] = dist[pos] + 1
					pq[front] = a*m + b
					front++
				}
			}
		}
		return res
	}

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]

		var tmp []Item
		for i, x := 0, 0; x < p && i < len(cur); x++ {
			j := i
			for i < len(cur) && cur[i].from == x {
				i++
			}
			tmp = append(tmp, expand(x, cur[j:i])...)
		}
		if len(tmp) > 0 {
			que = append(que, tmp)
		}
	}

	res := make([]int, p)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] >= 0 && vis[i][j] < p {
				res[vis[i][j]]++
			}
		}
	}

	return res
}

type Item struct {
	from int
	x    int
	y    int
}
