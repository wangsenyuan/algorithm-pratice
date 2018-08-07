package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for x := 1; x <= tc; x++ {
		n := readNum(scanner)
		G := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			G[i] = scanner.Text()
		}
		fmt.Fprintf(os.Stderr, "[debug %d] %d\n", x, n)
		for i := 0; i < n; i++ {
			fmt.Fprintf(os.Stderr, "[debug %d] %s\n", x, G[i])
		}
		// os.Stderr.Sync()
		res := solve(n, G)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

func solve(n int, G []string) int {
	row := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] == '1' {
				row[i] = j + 1
			}
		}
	}

	var ans int

	for i := 0; i < n; i++ {
		j := i
		for j < n && row[j] > i+1 {
			j++
		}
		ans += j - i
		//row[j] <= i + 1
		tmp := row[j]
		for k := j; k > i; k-- {
			row[k] = row[k-1]
		}
		row[i] = tmp
	}

	return ans
}
