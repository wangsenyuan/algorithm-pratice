package main

import (
	"bufio"
	"fmt"
	"os"
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
		scanner.Scan()

		fmt.Println(solve(scanner.Bytes()))

		t--
	}
}

func solve(s []byte) int {
	n := len(s)
	right := make([]int, n)

	right[n-1] = n - 1

	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			right[i] = right[i+1]
		} else {
			right[i] = i
		}
	}

	var ans int

	for i := 0; i < n; {
		j := right[i]

		if i < j {
			// same character from i to j, any sub array of [i:j] is good (length more than 1)
			m := j - i + 1
			ans += (m - 1) * m / 2
		}
		if i > 0 && j < n-1 && s[i-1] == s[j+1] {
			ans++
		}
		i = j + 1
	}

	return ans
}
