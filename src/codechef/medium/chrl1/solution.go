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
		t--
		n, K := readTwoNums(scanner)
		C := make([]int, n)
		W := make([]int, n)
		for i := 0; i < n; i++ {
			C[i], W[i] = readTwoNums(scanner)
		}
		fmt.Println(solve(n, K, W, C))
	}
}

func solve(n int, K int, W []int, C []int) int64 {
	N := 1 << uint(n)
	var best int64
	for i := 1; i < N; i++ {
		var cost, weight int64
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) > 0 {
				cost += int64(C[j])
				weight += int64(W[j])
			}
		}
		if cost <= int64(K) && weight > best {
			best = weight
		}
	}

	return best
}
