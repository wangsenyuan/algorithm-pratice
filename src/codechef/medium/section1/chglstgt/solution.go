package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		scanner.Scan()
		S := scanner.Bytes()
		res := solve(n, S)
		fmt.Println(res)
		t--
	}
}

func solve(n int, S []byte) int {
	pp := make([][]bool, n)

	// for i := 0; i < n; i++ {
	// 	pp[i] = make([]bool, n)
	// }

	// for i := 0; i < n; i++ {
	// 	l, r := i, i
	// 	for l >= 0 && r < n && S[l] == S[r] {
	// 		pp[l][r] = true
	// 		l--
	// 		r++
	// 	}
	// 	l = i - 1
	// 	r = i
	// 	for l >= 0 && r < n && S[l] == S[r] {
	// 		pp[l][r] = true
	// 		l--
	// 		r++
	// 	}
	// }

	for i := 0; i < n; i++ {
		pp[i] = make([]bool, n)
		pp[i][i] = true
		if i+1 < n && S[i] == S[i+1] {
			pp[i][i+1] = true
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			pp[i][j] = pp[i+1][j-1] && S[i] == S[j]
		}
	}

	// f[i] = mimimum splites ending at i
	f := make([]int, n)

	for i := 0; i < n; i++ {
		if pp[0][i] {
			f[i] = 1
			continue
		}
		best := i + 1
		for j := i; j > 0; j-- {
			if pp[j][i] {
				best = min(best, f[j-1]+1)
			}
		}
		f[i] = best
	}

	return f[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
