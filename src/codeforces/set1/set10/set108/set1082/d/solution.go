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

	n := readNum(reader)
	a := readNNums(reader, n)

	dia, res := solve(a)

	if dia < 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES %d\n%d\n", dia, len(res)))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(a []int) (int, [][]int) {
	n := len(a)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] > a[id[j]]
	})

	if a[id[0]] == 1 {
		return -1, nil
	}

	var res [][]int
	connect := func(i, j int) {
		res = append(res, []int{id[i] + 1, id[j] + 1})
	}
	var it int
	for it+1 < n && a[id[it+1]] > 1 {
		connect(it, it+1)
		it++
	}
	it++

	if it == n {
		return n - 1, res
	}

	if it == 1 {
		// 其他的必须都连接到id[0]
		if a[id[0]] < n-1 {
			return -1, nil
		}
		for i := 1; i < n; i++ {
			connect(0, i)
		}
		return 2, res
	}
	// it > 1
	//a[id[0]]--
	connect(0, it)

	if it+1 == n {
		return it, res
	}

	connect(it-1, it+1)
	//a[id[1]]--

	for i, j := it+2, 0; i < n; i++ {
		for j < it && a[id[j]] == 2 {
			j++
		}
		if j == it {
			return -1, nil
		}
		a[id[j]]--
		connect(j, i)
	}

	return it + 1, res
}
