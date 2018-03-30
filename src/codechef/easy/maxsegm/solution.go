package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
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

func readOneNum(scanner *bufio.Scanner) (a int) {
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

func toIntArray(s []byte) []int {
	n := len(s)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(s[i] - '0')
	}
	return res
}

func intsArrayToString(s []int) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		buf.WriteByte(byte(s[i] + '0'))
	}

	return buf.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readOneNum(scanner)

	for t > 0 {
		n := readOneNum(scanner)
		C := readNNums(scanner, n)
		W := readNNums(scanner, n)
		fmt.Println(solve(n, W, C))
		t--
	}
}

func solve(n int, W, C []int) int64 {
	elems := make([]int, n)

	var ans int64
	var sum int64
	for i, j := 0, 0; i < n; i++ {
		elems[C[i]]++
		sum += int64(W[i])
		if elems[C[i]] > 1 {
			// duplicate
			for j < i && elems[C[i]] > 1 {
				elems[C[j]]--
				sum -= int64(W[j])
				j++
			}
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}
