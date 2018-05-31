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

func readTwoNum(scanner *bufio.Scanner) (a int, b int) {
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
		A := readNNums(scanner, n)
		fmt.Println(solveTask2(n, A))
		t--
	}
}

func solveTask1(N int, A []int) int64 {
	sort.Ints(A[:N])
	// cnt[A[N-1]] = 1

	var ans int64
	for i := N - 2; i >= 0; i-- {
		for j := N - 1; j > i; j-- {
			if A[i]|A[j] == A[j] {
				ans++
			}
		}
	}
	return ans
}

var MAX_A = 1 << 22
var cnt = make([]int64, MAX_A)
var dp = make([]int64, MAX_A)
var t = make([]int64, MAX_A)

func solveTask2(N int, A []int) int64 {

	for i := 0; i < MAX_A; i++ {
		cnt[i] = 0
		t[i] = 0
	}

	for i := 0; i < N; i++ {
		cnt[A[i]]++
		t[A[i]]++
	}

	for j := 0; j < 22; j++ {
		for i := 0; i < MAX_A; i++ {
			if i&(1<<uint(j)) > 0 {
				cnt[i] += cnt[i^(1<<uint(j))]
			}
		}
	}

	var ans int64

	for i := 0; i < N; i++ {
		ans += cnt[A[i]] - t[A[i]]
	}

	for i := 0; i < MAX_A; i++ {
		ans += t[i] * (t[i] - 1) / 2
	}

	return ans
}
