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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N := readNum(scanner)
		if solve(N) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

const MAX_N = 110

var flag [220]bool

func init() {
	set := make([]bool, MAX_N)
	primes := make([]int, MAX_N)
	var j int

	for num := 2; num < MAX_N; num++ {
		if !set[num] {
			primes[j] = num
			j++
			for x := num * num; x < MAX_N; x += num {
				set[x] = true
			}
		}
	}

	primes = primes[:j]

	semi := make([]int, 220)
	var p int

	for i := 0; i < j; i++ {
		for k := i + 1; k < j; k++ {
			num := primes[i] * primes[k]
			if num < 220 {
				semi[p] = num
				p++
			}
		}
	}

	semi = semi[:p]

	sort.Ints(semi)

	p = 0
	for i := 1; i <= len(semi); i++ {
		if i < len(semi) && semi[i] == semi[i-1] {
			continue
		}
		semi[p] = semi[i-1]
		p++
	}
	semi = semi[:p]

	for i := 0; i < p; i++ {
		for j := i; j < p; j++ {
			num := semi[i] + semi[j]
			if num < 220 {
				flag[num] = true
			}
		}
	}
}

func solve(N int) bool {
	return flag[N]
}
