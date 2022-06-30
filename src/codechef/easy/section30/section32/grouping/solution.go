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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		C := readNNums(reader, n)
		P := make([][]int, m)
		for i := 0; i < m; i++ {
			P[i] = readNNums(reader, 2)
		}
		res := solve(n, k, P, C)
		for _, cur := range res {
			for i := 0; i < len(cur); i++ {
				buf.WriteString(fmt.Sprintf("%d ", cur[i]))
			}
			buf.WriteByte('\n')
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(n int, k int, pairs [][]int, C []int) [][]int {

	G := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		G[i] = make(map[int]bool)
	}

	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if C[a] > C[b] {
			G[a][b] = true
		} else {
			G[b][a] = true
		}
	}

	employees := make([]Employee, n)

	for i := 0; i < n; i++ {
		employees[i] = Employee{i, C[i]}
	}

	sort.Sort(Employees(employees))

	var dfs func(i int, p int)

	buf := make([]int, n)

	var res [][]int

	addResult := func(arr []int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		sort.Ints(tmp)
		res = append(res, tmp)
	}

	degree := make([]int, n)

	dfs = func(i int, p int) {
		for i < n {
			u := employees[i].index

			if degree[u] == p {
				// good to add it
				buf[p] = u

				for v := range G[u] {
					degree[v]++
				}

				dfs(i+1, p+1)

				if len(res) == k {
					return
				}
				for v := range G[u] {
					degree[v]--
				}
			}
			i++
		}
		if len(res) < k {
			addResult(buf[:p])
		}
	}

	dfs(0, 0)

	return res
}

type Employee struct {
	index int
	contr int
}

type Employees []Employee

func (this Employees) Len() int {
	return len(this)
}

func (this Employees) Less(i, j int) bool {
	return this[i].contr > this[j].contr
}

func (this Employees) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
