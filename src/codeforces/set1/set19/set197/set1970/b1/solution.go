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
	ok, houses, move := solve(n, a)

	if !ok {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer

	buf.WriteString("YES\n")

	for _, h := range houses {
		buf.WriteString(fmt.Sprintf("%d %d\n", h[0], h[1]))
	}
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", move[i]))
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

func solve(n int, a []int) (bool, [][]int, []int) {
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{a[i], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first > arr[j].first
	})
	// 从两头开始分配
	assign := make([]int, n)
	rev := make([]int, n)

	dispatch := func(i int, j int) {
		assign[arr[i].second] = j
		rev[j] = arr[i].second
	}

	l, r := 0, n-1
	for i := 0; i < n; {
		dispatch(i, l)
		i++
		l++
		if i < n {
			dispatch(i, r)
			r--
			i++
		}
	}

	houses := make([][]int, n)
	move_to := make([]int, n)
	for i := 0; i < n; i++ {
		j := assign[i]
		houses[i] = []int{j + 1, j + 1}
		d := a[i] / 2
		// (i - j) * 2 = x
		if j+d < n {
			move_to[i] = rev[j+d] + 1
		} else {
			move_to[i] = rev[j-d] + 1
		}
	}

	return true, houses, move_to
}

type Pair struct {
	first  int
	second int
}
