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
		a := readNNums(reader, n)
		res := solve(a)

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

const X = 100010

var factors [X][]int
var cnt [X]int
var idx [X][]int

func init() {
	for i := 1; i < X; i++ {
		factors[i] = getFactors(i)
	}
}

func getFactors(num int) []int {
	var res []int
	for i := 1; i <= num/i; i++ {
		if num%i == 0 {
			res = append(res, i)
			if num != i*i {
				res = append(res, num/i)
			}
		}
	}
	sort.Ints(res)
	return res
}

func solve(a []int) int {
	sort.Ints(a)
	n := len(a)

	for i := 0; i < n; i++ {
		for _, x := range factors[a[i]] {
			if len(idx[x]) == 0 {
				idx[x] = make([]int, 0, 1)
			}
			idx[x] = append(idx[x], i)
		}
	}

	m := a[n-1]
	var res int
	for x := m; x > 0; x-- {
		var sum int
		for i := len(idx[x]) - 1; i >= 0; i-- {
			cnt[x] += sum
			sum += n - idx[x][i] - 1
		}
		for y := 2 * x; y <= m; y += x {
			cnt[x] -= cnt[y]
		}
		res += x * cnt[x]
	}

	for x := m; x > 0; x-- {
		cnt[x] = 0
		idx[x] = idx[x][:0]
	}

	return res
}
