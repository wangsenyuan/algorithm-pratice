package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
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
		A := readNNums(scanner, 2*n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007

const MAX_N = 100111

func pow(a, b int) int64 {
	A := int64(a)
	res := int64(1)

	for b > 0 {
		if b&1 == 1 {
			res *= A
			res %= MOD
		}

		A *= A
		A %= MOD
		b >>= 1
	}

	return res
}

func inverse(a int) int64 {
	return pow(a, MOD-2)
}

var fact []int64
var invFacts []int64

func init() {
	fact = make([]int64, MAX_N)

	fact[0] = 1

	for i := 1; i < MAX_N; i++ {
		fact[i] = int64(i) * fact[i-1]
		fact[i] %= MOD
	}

	invFacts = make([]int64, MAX_N)

	for i := MAX_N - 1; i >= 0; i-- {
		invFacts[i] = inverse(int(fact[i]))
	}
}

func solve(n int, X []int) int64 {
	Y := make([]int, len(X)+2)
	copy(Y, X)
	Y[len(X)] = 0
	Y[len(X)+1] = 0
	sort.Ints(Y)

	ans := fact[n-1]
	cnt := make(map[Pair]int)

	skip0s := 2

	coef := int64(1)

	sum := Y[0] + Y[len(X)+1]

	for i, j := 0, len(X)+1; i < j; i, j = i+1, j-1 {
		num1, num2 := Y[i], Y[j]
		if num1+num2 != sum {
			return 0
		}
		if num1 == 0 || num2 == 0 {
			if skip0s > 0 {
				skip0s--
				continue
			}
		}

		p := Pair{num1, num2}
		cnt[p]++

		if num1 != num2 {
			coef *= 2
			coef %= MOD
		}
	}

	if skip0s > 0 {
		return 0
	}

	for _, c := range cnt {
		ans *= invFacts[c]
		ans %= MOD
	}

	ans *= coef
	ans %= MOD

	return ans
}

type Pair struct {
	first  int
	second int
}
