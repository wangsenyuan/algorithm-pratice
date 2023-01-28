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
	A := readNNums(reader, n)
	ok, res := solve(A)
	if !ok {
		fmt.Println("-1")
	} else {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, sw := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", sw[0], sw[1]))
		}
		fmt.Print(buf.String())
	}
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(A []int) (bool, [][]int) {
	if sort.IntsAreSorted(A) {
		return true, nil
	}
	n := len(A)
	// not sorted
	// 假设 需要交换将num放到position i, 且num 不是luck
	// 可以借助某个luck num， 先将position i 的数字和某个luck交换，然后再交换num
	luckAt := -1
	for i := 0; i < n; i++ {
		if checkLuck(A[i]) {
			luckAt = i
			break
		}
	}
	if luckAt < 0 {
		return false, nil
	}
	type Pair struct {
		first  int
		second int
	}

	B := make([]Pair, n)

	pos := make([]int, n)
	rpos := make([]int, n)

	for i := 0; i < n; i++ {
		B[i] = Pair{A[i], i}
		pos[i] = i
		rpos[i] = i
	}

	sort.Slice(B, func(i, j int) bool {
		return B[i].first < B[j].first || B[i].first == B[j].first && B[i].second < B[j].second
	})

	var res [][]int

	swap := func(a, b int) {
		x := rpos[a]
		y := rpos[b]
		if x == y {
			return
		}
		res = append(res, []int{x + 1, y + 1})
		rpos[pos[x]], rpos[pos[y]] = rpos[pos[y]], rpos[pos[x]]
		pos[x], pos[y] = pos[y], pos[x]
	}

	for i := 0; i < n; i++ {
		j := B[i].second
		if luckAt == j {
			continue
		}
		swap(pos[i], luckAt)
		swap(luckAt, j)
	}

	return true, res
}

func checkLuck(num int) bool {
	for num > 0 {
		r := num % 10
		num /= 10
		if r != 4 && r != 7 {
			return false
		}
	}
	return true
}
