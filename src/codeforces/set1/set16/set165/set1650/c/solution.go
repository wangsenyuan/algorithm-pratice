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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readString(reader)
		n, m := readTwoNums(reader)
		X, W := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			X[i], W[i] = readTwoNums(reader)
		}
		sum, res := solve(m, X, W, n)
		buf.WriteString(fmt.Sprintf("%d\n", sum))
		for _, op := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(m int, X []int, W []int, n int) (int64, [][]int) {
	// 2 * n <= m,
	// 感觉不需要管是否满足 li < lj < rj < ri的条件，
	// 始终选择最小weight 2 * n个点即可
	type Point struct {
		x  int
		w  int
		id int
	}

	ps := make([]Point, m)

	for i := 0; i < m; i++ {
		ps[i] = Point{X[i], W[i], i + 1}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].w < ps[j].w
	})
	qs := ps[:2*n]

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].x < qs[j].x
	})

	var sum int64
	res := make([][]int, n)

	for i, j := 0, 2*n-1; i < j; i, j = i+1, j-1 {
		sum += int64(qs[i].w + qs[j].w)
		res[i] = []int{qs[i].id, qs[j].id}
	}

	return sum, res
}
