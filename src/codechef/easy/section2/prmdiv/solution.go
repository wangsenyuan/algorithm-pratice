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

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

const MAX_N = 1000000

var good [][]int

func init() {
	S := make([]int, MAX_N+1)
	for i := 2; i <= MAX_N; i++ {
		if S[i] == 0 {
			for j := i; j <= MAX_N; j += i {
				S[j] += i
			}
		}
	}

	good = make([][]int, MAX_N+1)
	for i := 0; i <= MAX_N; i++ {
		good[i] = make([]int, 0, 10)
	}
	for i := 2; i <= MAX_N; i++ {
		if S[i] == 0 {
			continue
		}
		for j := i; j <= MAX_N; j += i {
			// j divides i
			if S[j]%S[i] == 0 {
				good[j] = append(good[j], i)
			}
		}
	}
}

func solve(n int, A []int) int64 {
	freq := make([]int64, MAX_N+1)
	for i := 0; i < n; i++ {
		freq[A[i]]++
	}
	var ans int64
	for i := 2; i <= MAX_N; i++ {
		for _, j := range good[i] {
			ans += freq[j] * freq[i]
		}
	}
	ans -= int64(n)
	return ans
}
