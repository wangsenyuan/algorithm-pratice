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

	ask := func(i int) int {
		fmt.Printf("? %d\n", i)
		return readNum(scanner)
	}

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		N, _ := readTwoNums(scanner)

		res := solve(N, ask)

		fmt.Printf("! %d\n", res)

		scanner.Scan()
	}
}

const MAX_N = 1000007

var factors []int

func init() {
	factors = make([]int, MAX_N)

	set := make([]bool, MAX_N)

	for x := 2; x < MAX_N; x++ {
		if !set[x] {
			for y := x; y < MAX_N; y += x {
				set[y] = true
				factors[y] = x
			}
		}
	}
}

func solve(n int, ask func(int) int) int {
	fs := findPrimeFactors(n)

	check := func(x int) bool {
		if x == 0 {
			return false
		}
		if x == n {
			return true
		}
		return x+ask(x) == n
	}

	prod := 1

	for i := 0; i < len(fs); i++ {
		cur := fs[i]
		f, s := cur.first, cur.second

		full := n / pow(f, s)

		var low int
		high := s + 1

		for low < high {
			mid := (low + high) / 2

			x := full * pow(f, mid)

			if check(x) {
				high = mid
			} else {
				low = mid + 1
			}
		}

		prod *= pow(f, high)
	}

	return n / prod
}

func findPrimeFactors(n int) []Pair {
	res := make([]int, 0, 10)

	for n > 1 {
		res = append(res, factors[n])
		n /= factors[n]
	}

	sort.Ints(res)

	ans := make([]Pair, 0, len(res))

	var j int

	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			ans = append(ans, Pair{res[i-1], i - j})
			j = i
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}
