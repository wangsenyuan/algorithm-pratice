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
		A := readNNums(scanner, n)

		fmt.Println(solve(n, A))

		t--
	}
}

func solve(n int, A []int) int {
	nums := make([]int, 2001)

	for i := 0; i < n; i++ {
		nums[1000+A[i]]++
	}

	var res int
	for num, cnt := range nums {
		if cnt > 0 {
			res += cnt * (cnt - 1) / 2
			for left, right := num-1, num+1; left >= 0 && right <= 2000; left, right = left-1, right+1 {
				res += nums[left] * nums[right]
			}
		}
	}
	return res
}
