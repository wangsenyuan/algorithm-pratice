package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	T := readNNums(reader, n)

	res := solve(n, m, E, T)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(strconv.Itoa(res[i]))
		buf.WriteByte(' ')
	}
	fmt.Println(buf.String())
}

func solve(n, m int, E [][]int, T []int) []int {
	X := make([]Item, n)

	for i := 0; i < n; i++ {
		X[i] = Item{i, T[i]}
	}

	sort.Sort(Items(X))

	degree := make([]int, n)

	for _, e := range E {
		u, v := e[0], e[1]
		degree[u-1]++
		degree[v-1]++
	}

	conn := make([][]int, n)

	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
	}

	return solve1(n, X, conn)
}

func solve2(n int, X []Item, conn [][]int) []int {
	ans := make([]int, n)
	last := make([]int, n)
	res := make([]int, n)
	for i, x := range X {
		u := x.first
		for _, v := range conn[u] {
			last[ans[v]] = u
		}
		ans[u] = 1

		for last[ans[u]] == u {
			ans[u]++
		}
		if ans[u] != i {
			return nil
		}
		res[i] = u + 1
	}
	return res
}

func solve1(n int, X []Item, conn [][]int) []int {
	assign := make([]int, n)

	for i := 0; i < n; i++ {
		assign[i] = -1
	}

	mem := make(map[int]int)

	for _, item := range X {
		// cur == item.second
		u := item.first
		for _, v := range conn[u] {
			if assign[v] < 0 {
				continue
			}
			if assign[v] == item.second {
				return nil
			}
			if assign[v] < item.second {
				mem[assign[v]]++
			}
		}

		if len(mem) != item.second-1 {
			return nil
		}

		for _, v := range conn[u] {
			if assign[v] < 0 {
				continue
			}
			mem[assign[v]]--
			if mem[assign[v]] == 0 {
				delete(mem, assign[v])
			}
		}

		assign[u] = item.second
	}

	for i := 0; i < n; i++ {
		if assign[i] < 0 {
			return nil
		}
	}

	res := make([]int, n)

	for i, item := range X {
		res[i] = item.first + 1
	}

	return res
}

type Item struct {
	first, second int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].second < items[j].second
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
