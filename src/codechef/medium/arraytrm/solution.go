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
		n, k := readTwoNums(scanner)
		arr := readNNums(scanner, n)

		res := solve(n, arr, k)

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}

func solve(n int, arr []int, k int) bool {
	rem := make(map[int]int)

	for i := 0; i < n; i++ {
		rem[arr[i]%(k+1)]++
	}

	if len(rem) == 1 {
		return true
	}

	if len(rem) > 2 {
		return false
	}

	for _, v := range rem {
		if v == n-1 || v == n {
			return true
		}
	}

	return false
}
