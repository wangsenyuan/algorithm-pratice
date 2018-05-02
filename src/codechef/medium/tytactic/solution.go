package main

import (
	"bufio"
	"os"
	"bytes"
	"strconv"
	"fmt"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNum(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNum(scanner)
	skills := readNNums(scanner, n)

	edges := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(scanner, 2)
	}

	queries := make([][]byte, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		queries[i] = scanner.Bytes()
	}

	res := solve(n, m, skills, edges, queries)
	var buf bytes.Buffer
	for _, num := range res {
		buf.WriteString(strconv.Itoa(num))
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func solve(n int, m int, skills []int, edges [][]int, queries [][]byte) []int {
	neigbhors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		neigbhors[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		neigbhors[a][b] = true
		neigbhors[b][a] = true
	}

	begin := make([]int, n)
	end := make([]int, n)

	var dfs func(p, u int, open int) int
	dfs = func(p, u int, open int) int {
		begin[u] = open
		close := open
		for v := range neigbhors[u] {
			if p == v {
				continue
			}
			close = dfs(u, v, close+1)
		}
		end[u] = close
		return end[u]
	}

	dfs(-1, 0, 0)

	bit := BIT(n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = skills[i]
	}

	res := make([]int, m)
	m = 0
	bit(func(update func(int, int), query func(int) int) {
		for i := 0; i < n; i++ {
			update(begin[i], skills[i])
		}

		for _, que := range queries {
			var id, x int
			j := toNum(que, &id, 2)
			id--
			if que[0] == 'Q' {
				res[m] = query(end[id]) - query(begin[id]-1)
				m++
			} else {
				toNum(que, &x, j+1)
				update(begin[id], x-arr[id])
				arr[id] = x
			}
		}

	})

	return res[:m]
}

func BIT(n int) func(func(func(int, int), func(int) int)) {
	arr := make([]int, n+1)

	update := func(i int, x int) {
		i++
		for i <= n {
			arr[i] += x
			i += i & (-i)
		}
	}

	query := func(i int) int {
		i++
		var sum int
		for i > 0 {
			sum += arr[i]
			i -= i & (-i)
		}
		return sum
	}

	g := func(f func(func(int, int), func(int) int)) {
		f(update, query)
	}

	return g
}

func toNum(bs []byte, num *int, from int) int {
	var res int
	i := from
	for i < len(bs) && bs[i] != ' ' {
		res = res*10 + int(bs[i]-'0')
		i++
	}
	*num = res
	return i
}
