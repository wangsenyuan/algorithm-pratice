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
		n := readNum(reader)
		l := readNNums(reader, n)
		r := readNNums(reader, n)
		c := readNNums(reader, n)
		res := solve(l, r, c)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

func solve(l []int, r []int, c []int) int {
	n := len(l)

	evets := make([]Pair, 2*n)
	for i := 0; i < n; i++ {
		evets[i] = Pair{l[i], 0}
	}
	for i := 0; i < n; i++ {
		evets[i+n] = Pair{r[i], 1}
	}

	sort.Slice(evets, func(i, j int) bool {
		return evets[i].first < evets[j].first || evets[i].first == evets[j].first && evets[i].second > evets[j].second
	})

	stack := make([]int, n)
	var top int

	var arr []Pair

	for _, evt := range evets {
		if evt.second == 0 {
			stack[top] = evt.first
			top++
		} else {
			// = 1
			arr = append(arr, Pair{stack[top-1], evt.first})
			top--
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]
		return a.second-a.first < b.second-b.first
	})
	sort.Ints(c)
	reverse(c)

	var res int
	for i, cur := range arr {
		res += c[i] * (cur.second - cur.first)
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
