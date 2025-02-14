package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	res := process(reader)
	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\r\n", cur[0], cur[1]))
	}
	fmt.Fprint(w, buf.String())
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
	cmds := make([]string, k)
	for i := range k {
		cmds[i] = readString(reader)
	}
	return solve(n, m, cmds)
}

func solve(n int, m int, cmds []string) [][]int {

	warehouse := make([][]int, n)
	for i := 0; i < n; i++ {
		warehouse[i] = make([]int, m)
	}

	var ans [][]int
	var drinks []string

	put := func(x int, y int, id int) int {
		for x < n {
			if warehouse[x][y] == 0 {
				warehouse[x][y] = id
				return x*m + y
			}
			y++
			if y == m {
				y = 0
				x++
			}
		}
		return -1
	}

	where := make(map[string]int)

	find := func(id string) []int {
		if v, ok := where[id]; !ok || v < 0 {
			return []int{-1, -1}
		} else {
			i, j := v/m, v%m
			warehouse[i][j] = 0
			delete(where, id)
			return []int{i + 1, j + 1}
		}
	}

	for _, cmd := range cmds {
		if cmd[0] == '+' {
			// +1 x y id
			var x, y, pos int
			pos = readInt([]byte(cmd), 3, &x) + 1
			pos = readInt([]byte(cmd), pos, &y) + 1
			id := cmd[pos:]
			// 不管怎么样都记下来
			drinks = append(drinks, id)
			where[id] = put(x-1, y-1, len(drinks))
		} else {
			id := cmd[3:]
			ans = append(ans, find(id))
		}
	}
	return ans
}
