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

	n, m := readTwoNum(scanner)

	mtr := make([][]byte, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		mtr[i] = scanner.Bytes()
	}
	fmt.Println(solve(n, m, mtr))
}

func solve(N, M int, MTR [][]byte) int64 {
	cols := make([]int64, M)
	var ans int64

	for i := N - 1; i >= 0; i-- {
		for j := 0; j < M; j++ {
			if MTR[i][j] == '1' {
				cols[j]++
			} else {
				cols[j] = 0
			}
		}
		var last int64
		for j := 0; j < M; j++ {
			last = min(last+1, cols[j])
			ans += last
		}
	}

	return ans
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
