package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		var n uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &n)
		res := solve(int64(n))
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const INF = math.MaxInt64
const X = 100010

var P []int64

func init() {
	P = make([]int64, 0, 20*X)
	for i := 2; i < X; i++ {
		I := int64(i)
		J := I * I * I * I
		for J < INF && J > 0 {
			P = append(P, J)
			J *= I
		}
	}
	sort.Sort(Int64s(P))
}

type Int64s []int64

func (this Int64s) Len() int {
	return len(this)
}

func (this Int64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve(n int64) int64 {
	i := sort.Search(len(P), func(i int) bool {
		return P[i] >= n
	})
	if P[i] == n {
		return P[i]
	}
	res := P[i]
	if i-1 >= 0 && P[i]-n >= n-P[i-1] {
		res = P[i-1]
	}

	x := int64(math.Sqrt(float64(n)))

	if x > 0 {
		xx := x * x
		if abs(n-res) >= abs(n-xx) {
			res = xx
		}
		xx = (x + 1) * (x + 1)
		if abs(n-res) > abs(n-xx) {
			res = xx
		}
	}

	x = int64(math.Pow(float64(n), 1.0/3))

	if x > 0 {
		xx := x * x * x
		if abs(n-res) >= abs(n-xx) {
			res = xx
		}
		xx = (x + 1) * (x + 1) * (x + 1)
		if abs(n-res) > abs(n-xx) {
			res = xx
		}
	}

	return res
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
