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
		scanner.Scan()
		s := scanner.Text()
		res := solve(s)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

func solve(s string) int64 {
	n := len(s)
	//base has to be at lease len(bs)
	xs := make(map[byte]int64)
	// to get minimim, first has to be 1
	xs[s[0]] = 1
	var num int64
	for i := 1; i < n; i++ {
		if _, found := xs[s[i]]; !found {
			xs[s[i]] = num
			num++
			if num == 1 {
				num++
			}
		}
	}
	if num < 2 {
		num = 2
	}
	//base has to be at lease num
	var res int64
	for i := 0; i < n; i++ {
		x := xs[s[i]]
		res = res*num + x
	}
	return res
}
