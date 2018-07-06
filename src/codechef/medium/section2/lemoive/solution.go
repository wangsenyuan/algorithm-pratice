package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n, k := readTwoNums(scanner)
		ps := readNNums(scanner, n)
		fmt.Println(solve(n, k, ps))
		tc--
	}
}

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

const MOD = 1000000007

func solve(n int, K int, ps []int) int64 {
	sort.Ints(ps)

	cnts := make([]int, n)
	var j, k int
	for i := 1; i <= n; i++ {
		if i < n && ps[i] == ps[i-1] {
			continue
		}
		cnts[j] = i - k
		j++
		k = i
	}

	numPerm := make([]int64, K+2)
	numPerm[0] = 1
	update := func(m int, n int) int {
		var t int64 = 1
		for i := m + 1; i < m+n; i++ {
			t = (t * int64(i)) % MOD
		}
		a := (t * int64(n)) % MOD
		b := (t * int64(m)) % MOD
		for r := K; r >= 0; r-- {
			numPerm[r+1] += (numPerm[r] * a) % MOD
			if numPerm[r+1] >= MOD {
				numPerm[r+1] -= MOD
			}
			numPerm[r] = (numPerm[r] * b) % MOD
		}
		return m + n
	}
	var m int
	for k := j - 1; k >= 0; k-- {
		m = update(m, cnts[k])
	}
	var res int64

	for k := 1; k <= K; k++ {
		res += numPerm[k]
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}
