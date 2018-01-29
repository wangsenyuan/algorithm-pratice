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

	for i := 0; i < t; i++ {
		n := readNum(scanner)
		row := readNNums(scanner, n)
		fmt.Println(solve(n, row))
	}
}

func solve(n int, row []int) int {
	var prev, ans int

	for i, j := 0, 0; i < n; i++ {
		if row[i] == 1 {
			//a girl
			if i-j <= prev && prev != 0 {
				ans = prev + 1
			} else {
				ans = i - j
			}
			j++
			prev = ans
		}
	}
	return ans
}
