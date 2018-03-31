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

	for t > 0 {
		n := readNum(scanner)
		arr := readNNums(scanner, n)
		fmt.Println(solve(n, arr))
		t--
	}
}

func solve(n int, arr []int) int {
	rem := make([]int, 4)

	for i := 0; i < n; i++ {
		rem[arr[i]%4]++
	}

	var ans int
	if rem[2] > 0 {
		ans += rem[2] / 2
		rem[2] = rem[2] % 2
	}

	a := min(rem[1], rem[3])
	ans += a
	rem[1] -= a
	rem[3] -= a

	b := rem[1]
	if rem[3] > 0 {
		b = rem[3]
	}

	ans += (b / 4) * 3

	b = b % 4

	if b == 1 || b == 3 {
		return -1
	}

	if b == 0 {
		if rem[2] > 0 {
			return -1
		}
		return ans
	}

	// b == 2
	if rem[2] == 0 {
		return -1
	}
	return ans + 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
