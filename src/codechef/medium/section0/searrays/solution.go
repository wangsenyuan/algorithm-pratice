package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

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

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	for t > 0 {
		n, k := readTwoNums(scanner)
		f.WriteString(fmt.Sprintf("%d\n", solve(n, k)))
		t--
	}
}

func solve(n, k int) int64 {
	if n < k {
		return 1
	}
	if n == k {
		return 2
	}
	f := make([]int64, n+1)

	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		if i >= k {
			f[i] = (f[i] + f[i-k]) % MOD
		}
	}
	return f[n]
}
