package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		Q := readNum(scanner)
		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			scanner.Scan()
			bs := scanner.Bytes()
			var tp int
			pos := readInt(bs, 0, &tp)
			if tp == 1 {
				queries[i] = make([]int, 3)
				queries[i][0] = 1
				for j := 1; j < 3; j++ {
					pos = readInt(bs, pos+1, &queries[i][j])
				}
			} else {
				queries[i] = make([]int, 4)
				queries[i][0] = 2
				for j := 1; j < 4; j++ {
					pos = readInt(bs, pos+1, &queries[i][j])
				}
			}
		}
		res := solve(n, A, Q, queries)
		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

const MAX_NUM = 100000

var factors [][]int
var records []int
var zeros []int
var cnt [][]int

func init() {
	primes := make([]int, MAX_NUM+1)
	for i := 2; i <= MAX_NUM; i++ {
		primes[i] = i
	}
	for i := 2; i <= MAX_NUM; i++ {
		if primes[i] == i {
			for j := i + i; j <= MAX_NUM; j += i {
				primes[j] = i
			}
		}
	}
	factors = make([][]int, MAX_NUM+1)
	factors[0] = make([]int, 0)
	factors[1] = make([]int, 0)
	hsh := make([]bool, MAX_NUM+1)
	for x := 2; x <= MAX_NUM; x++ {
		factors[x] = factorize(x, primes, hsh)
	}

	records = make([]int, MAX_NUM+1)
	zeros = make([]int, MAX_NUM+1)
	cnt = make([][]int, 250)
	for i := 0; i < 250; i++ {
		cnt[i] = make([]int, MAX_NUM+1)
	}
}

func factorize(num int, primes []int, hsh []bool) []int {
	res := make([]int, 0, 10)
	for num > 1 {
		p := primes[num]
		num /= p
		if !hsh[p] {
			res = append(res, p)
		}
		hsh[p] = true
	}

	for _, ans := range res {
		hsh[ans] = false
	}

	return res
}

func solve(n int, A []int, Q int, queries [][]int) []int {
	copy(records, zeros)

	for i := 0; i < 250; i++ {
		copy(cnt[i], zeros)
	}

	sq := int(math.Sqrt(float64(n)))

	add := func(num int, idx int, val int) {
		fs := factors[num]
		sz := len(fs)

		for i := 1; i < (1 << uint(sz)); i++ {
			mul := 1
			for j := 0; j < sz; j++ {
				if i&(1<<uint(j)) > 0 {
					mul *= fs[j]
				}
			}
			cnt[idx][mul] += val
		}
	}

	count := func(index int, g int) int {
		fs := factors[g]
		sz := len(fs)

		var ans int

		for i := 1; i < (1 << uint(sz)); i++ {
			ct, mul := 0, 1
			for j := 0; j < sz; j++ {
				if i&(1<<uint(j)) > 0 {
					ct++
					mul *= fs[j]
				}
			}
			if ct%2 == 0 {
				ans -= cnt[index][mul]
			} else {
				ans += cnt[index][mul]
			}
		}
		return ans
	}

	find := func(index int, g int) int {
		if index < 0 {
			return 0
		}
		times := index / sq
		var res int
		for i := 0; i < times; i++ {
			res += count(i, g)
		}
		for i := index; i/sq == times && i > 0; i-- {
			if gcd(g, A[i]) > 1 {
				res++
			}
		}
		return res
	}

	for i := 0; i < n; i++ {
		add(A[i], i/sq, 1)
	}

	res := make([]int, Q)
	var q int

	for i := 0; i < Q; i++ {
		query := queries[i]
		if query[0] == 1 {
			x, y := query[1]-1, query[2]
			add(A[x], x/sq, -1)
			A[x] = y
			add(A[x], x/sq, 1)
		} else {
			l, r, g := query[1]-1, query[2]-1, query[3]
			res[q] = r - l + 1 - (find(r, g) - find(l-1, g))
			q++
		}
	}
	return res[:q]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
