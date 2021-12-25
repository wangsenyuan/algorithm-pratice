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
		N := readNum(scanner)
		res := solve(N)
		fmt.Println(res)
	}
}

func solve(N int) int {
	if N == 0 {
		return 3
	}

	if N == 1 {
		return 2
	}

	ones := countOnes(N)

	if ones == 1 {
		return 1
	}

	if ones == 2 {
		return 0
	}
	var x, indexSecond, indexFirst int
	for i := 0; i < 31; i++ {
		if N&(1<<uint(i)) > 0 {
			x++
			if x+1 == ones {
				indexSecond = x
			} else if x == ones {
				indexFirst = x
			}
		}
	}
	a := N & (1<<uint(indexSecond) - 1)

	shift := 1 << uint(indexSecond+1)
	b := shift - N&(shift-1)
	if indexFirst == indexSecond+1 {
		b++
	}

	return min(a, b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countOnes(num int) int {
	var cnt int
	for num > 0 {
		cnt += num & 1
		num >>= 1
	}
	return cnt
}
