package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)

	A := readNNums(reader, n)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(n, A, Q)

	var buf bytes.Buffer
	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d\n", num))
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
		if s[i] == '\n' {
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

func solve1(n int, A []int, Q [][]int) []int64 {
	res := make([]int64, len(Q))

	for i, cur := range Q {
		x, l, r := cur[0], cur[1], cur[2]

		for j := l - 1; j < r; j++ {
			if x%A[j] == 0 {
				res[i] += int64(x / A[j])
			} else {
				res[i] += int64(x/A[j]) + 1
			}
		}
	}

	return res
}

const LI = 100010

func solve(n int, A []int, Q [][]int) []int64 {
	bit := make([]int64, LI+1)

	add := func(p int, v int64) {
		for p < len(bit) {
			bit[p] += v
			p += p & -p
		}
	}

	get := func(p int) int64 {
		var res int64
		for p > 0 {
			res += bit[p]
			p -= p & -p
		}
		return res
	}

	calc := func(x int) int64 {
		var ans int64
		small := int(math.Ceil(math.Sqrt(float64(x))))

		for i := 1; ; i++ {
			val := (x + i - 1) / i
			if val <= small {
				break
			}
			cnt := get(i) - get(i-1)
			ans += int64(val) * int64(cnt)
		}

		for i := 1; i <= small; i++ {
			start := (x + i - 1) / i
			var end int
			if i == 1 {
				end = LI
			} else {
				end = x / (i - 1)
				if x%(i-1) == 0 {
					end--
				}
			}
			cnt := get(end) - get(start-1)
			ans += int64(cnt) * int64(i)
		}

		return ans
	}

	pp := make([][]Pair, n+1)
	comp := make([]int64, len(Q))

	for i := 0; i < len(comp); i++ {
		comp[i] = -1
	}

	app := func(i int, p Pair) {
		if pp[i] == nil {
			pp[i] = make([]Pair, 0, 2)
		}
		pp[i] = append(pp[i], p)
	}

	for i, cur := range Q {
		x, l, r := cur[0], cur[1], cur[2]
		if pp[r] == nil {
			pp[r] = make([]Pair, 0, 2)
		}
		app(r, Pair{x, i})
		if l == 1 {
			comp[i] = 0
		} else {
			app(l-1, Pair{x, i})
		}
	}

	for i := 1; i <= n; i++ {
		add(A[i-1], 1)
		for _, cur := range pp[i] {
			ans := calc(cur.first)
			j := cur.second
			if comp[j] < 0 {
				comp[j] = ans
			} else {
				comp[j] = ans - comp[j]
			}
		}
	}
	return comp
}

type Pair struct {
	first  int
	second int
}
