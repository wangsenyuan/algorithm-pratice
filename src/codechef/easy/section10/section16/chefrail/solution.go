package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		X := readNNums(scanner, n)
		Y := readNNums(scanner, m)
		fmt.Println(solve(n, m, X, Y))
	}
}

const MAX_N = 100007

var factors []int

var c [][]int
var d []int

func init() {
	factors = make([]int, MAX_N)

	for x := 2; x < MAX_N; x++ {
		if factors[x] == 0 {
			for y := 0; y < MAX_N; y += x {
				factors[y] = x
			}
		}
	}

	c = make([][]int, 2)

	for i := 0; i < 2; i++ {
		c[i] = make([]int, MAX_N)
	}

	d = make([]int, MAX_N)
}

func solve(n, m int, X []int, Y []int) int64 {

	var zero int64

	for i, j := 0, 0; i < len(X); i++ {
		if X[i] == 0 {
			zero = 1
			n--
			continue
		}
		X[j] = X[i]
		j++
	}

	for i, j := 0, 0; i < len(Y); i++ {
		if Y[i] == 0 {
			zero = 1
			m--
			continue
		}
		Y[j] = Y[i]
		j++
	}

	X = X[:n]
	Y = Y[:m]

	res := zero * int64(m) * int64(n)

	res += solve3(X, Y)
	res += solve3(Y, X)

	return res
}

func solve3(X, Y []int) int64 {
	updateFreq(X, Y)

	var ans int64

	for a := int64(1); a*a < MAX_N; a++ {
		for b := int64(1); b*b < MAX_N; b++ {
			g := gcd(a, b)
			if g != 1 {
				continue
			}
			for x := int64(1); x*a*a < MAX_N && x*b*b < MAX_N; x++ {
				y := int(x * a * b)

				ans += int64(d[y]) * int64(c[0][int(x*a*a)]) * int64(c[1][int(x*b*b)])
			}
		}
	}
	return ans
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func updateFreq(X, Y []int) {
	for i := 0; i < 2; i++ {
		for j := 0; j < MAX_N; j++ {
			c[i][j] = 0
		}
	}

	for i := 0; i < len(Y); i++ {
		y := Y[i]
		if y < 0 {
			c[1][-y]++
		} else {
			c[0][y]++
		}
	}

	for i := 0; i < MAX_N; i++ {
		d[i] = 0
	}

	for i := 0; i < len(X); i++ {
		d[abs(X[i])]++
	}

}

func solve2(X []int, Y []int) int64 {
	updateFreq(X, Y)
	var ans int64

	for i := 0; i < len(X); i++ {
		x := X[i]
		x2 := abs(x)

		var ps []Pair

		for x2 > 1 {
			p := factors[x2]
			var cnt int
			for x2%p == 0 {
				cnt++
				x2 /= p
			}
			ps = append(ps, Pair{p, cnt})
		}

		sort.Sort(Pairs(ps))

		f := getDivisors(ps)

		for _, j := range f {
			if x*x/j < MAX_N {
				ans += int64(c[0][j]) * int64(c[1][x*x/j])
			}
		}
	}

	return ans
}

func getDivisors(ps []Pair) []int {
	var res []int

	var dfs func(i int, cur int64)

	dfs = func(i int, cur int64) {
		if i == len(ps) || cur*int64(ps[i].first) > MAX_N {
			res = append(res, int(cur))
			return
		}

		for x := 0; x <= 2*ps[i].second && cur < MAX_N; x++ {
			dfs(i+1, cur)
			cur *= int64(ps[i].first)
		}
	}

	dfs(0, 1)

	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
