package main

import (
	"bufio"
	"os"
	"fmt"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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
		fmt.Println(solve(n))
		t--
	}
}

var PC []int64
var MAX_N = 2500

func init() {
	HALF := (MAX_N + 1) / 2
	MAX := HALF * HALF

	PC = make([]int64, MAX+1)
	for i := 1; i <= MAX; i++ {
		for j := i; j <= MAX; j += i {
			PC[j]++
		}
	}

	for i := 2; i <= MAX; i++ {
		PC[i] += PC[i-1]
	}
	//fmt.Printf("%d\n", PC[MAX])
}

func solve(n int) int64 {
	var ans int64

	for a := 1; a < n; a++ {
		b := n - a
		m := a * b
		ans += PC[m-1]
	}

	return ans
}
