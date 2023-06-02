package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res, path := solve(n, Q)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i], len(path[i])))
		for j := 0; j < len(path[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", path[i][j]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n int, Q [][]int) ([]int, [][]int) {
	ans := make([]int, len(Q))
	path := make([][]int, len(Q))

	h := 1

	for h < n {
		h = h*2 + 1
	}

	for i, cur := range Q {
		u, v := cur[0], cur[1]
		if n <= 10 {
			ans[i], path[i] = querySmall(n, u, v)
		} else {
			ans[i], path[i] = queryLarge(h, n, u, v)
		}
	}
	return ans, path
}

func queryLarge(h int, n int, u int, v int) (int, []int) {
	if u^v == h {
		return h, []int{u, v}
	}
	w := h ^ u ^ v
	if w <= n && w != u && w != v {
		return h, []int{u, w, v}
	}

	for x := 1; x <= n; x++ {
		y := w ^ x
		if set_size(x, y, u, v) == 4 && y <= n && y >= 1 {
			return h, []int{u, x, y, v}
		}
	}
	for z := 1; z <= 2; z++ {
		for x := 1; x <= n; x++ {
			y := w ^ x ^ z
			if y >= 1 && y <= n && set_size(u, v, x, y, z) == 5 {
				return h, []int{u, x, y, z, v}
			}
		}
	}

	ans := n - 2
	path := []int{u, (n - 2) ^ 1, v}
	return ans, path
}

func set_size(arr ...int) int {
	set := make(map[int]bool)
	for _, num := range arr {
		set[num] = true
	}
	return len(set)
}

func querySmall(n int, u int, v int) (int, []int) {
	// n <= 10
	u--
	v--

	N := 1 << n
	var best int
	var state int
	for mask := 1; mask < N; mask++ {
		if (mask>>u)&1 == 0 || (mask>>v)&1 == 0 {
			continue
		}
		var tmp int
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				tmp ^= (i + 1)
			}
		}
		if tmp > best {
			best = tmp
			state = mask
		}
	}

	var res []int
	res = append(res, u+1)
	for i := 0; i < n; i++ {
		if i == u || i == v {
			continue
		}
		if (state>>i)&1 == 1 {
			res = append(res, i+1)
		}
	}
	res = append(res, v+1)
	return best, res
}
