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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		S := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			S[i] = scanner.Text()
		}
		fmt.Println(solve(n, S))
	}
}

func solve(n int, S []string) int {
	if n == 0 {
		return 0
	}
	flag := make([]bool, 26)
	for j := 0; j < len(S[0]); j++ {
		flag[int(S[0][j]-'a')] = true
	}

	tmp := make([]bool, 26)

	for i := 1; i < n; i++ {
		for j := 0; j < 26; j++ {
			tmp[j] = false
		}
		for j := 0; j < len(S[i]); j++ {
			tmp[int(S[i][j]-'a')] = true
		}
		for j := 0; j < 26; j++ {
			flag[j] = flag[j] && tmp[j]
		}
	}

	var res int

	for i := 0; i < 26; i++ {
		if flag[i] {
			res++
		}
	}

	return res
}
