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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nums := make([]int, MAX_N)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		fillNNums(scanner, n, nums)
		res := solve(n, nums[:n])
		fmt.Println(res)
	}

}

const MAX_N = 1e5 + 10
const MOD = 1e9 + 7

func solve(n int, A []int) int {
	sort.Ints(A)
	var ans int64 = 1
	var carry bool
	var j int
	for i := 1; i <= n; i++ {
		if i < n && A[i] == A[j] {
			i++
			continue
		}
		k := int64(i - j)
		if carry {
			ans = (ans * k) % MOD
			k--
			carry = false
		}
		if k&1 == 1 {
			ans = (ans * k) % MOD
			k--
			carry = true
		}
		for k > 0 {
			ans = (ans * (k - 1)) % MOD
			k -= 2
		}
		j = i
	}

	return int(ans)
}
