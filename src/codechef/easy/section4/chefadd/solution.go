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

		var a, b, c uint64
		scanner.Scan()
		pos := readUint64(scanner.Bytes(), 0, &a)
		pos = readUint64(scanner.Bytes(), pos+1, &b)
		readUint64(scanner.Bytes(), pos+1, &c)

		fmt.Println(solve(uint(a), uint(b), uint(c)))
	}
}

const B = 32

var dp [B][B][B][2]int64

func solve(a, b, c uint) int64 {
	for i := 0; i < B; i++ {
		for j := 0; j < B; j++ {
			for k := 0; k < B; k++ {
				for u := 0; u < 2; u++ {
					dp[i][j][k][u] = -1
				}
			}
		}
	}

	var fn func(bit, a, b, cf int, c uint) int64

	fn = func(bit, a, b, cf int, c uint) int64 {
		if a < 0 || b < 0 {
			return 0
		}
		if bit == 32 {
			if a == 0 && b == 0 && cf == 0 {
				return 1
			}
			return 0
		}

		if dp[bit][a][b][cf] < 0 {
			var x int
			if c&uint(1<<uint(bit)) > 0 {
				x = 1
			}
			var ans int64
			if x == cf {
				ans = fn(bit+1, a, b, 0, c) + fn(bit+1, a-1, b-1, 1, c)
			}
			if cf%2 != x {
				ans += fn(bit+1, a-1, b, (cf+1)/2, c)
				ans += fn(bit+1, a, b-1, (cf+1)/2, c)
			}
			dp[bit][a][b][cf] = ans
		}

		return dp[bit][a][b][cf]
	}

	return fn(0, bitCount(a), bitCount(b), 0, c)
}

func bitCount(num uint) int {
	var ans int
	for num != 0 {
		ans++
		num = num & (num - 1)
	}
	return ans
}
