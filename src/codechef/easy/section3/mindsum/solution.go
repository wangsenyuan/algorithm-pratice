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
		var n, d uint64
		scanner.Scan()
		pos := readUint64(scanner.Bytes(), 0, &n)
		readUint64(scanner.Bytes(), pos+1, &d)
		ans, cnt := solve(int64(n), int64(d))
		fmt.Printf("%d %d\n", ans, cnt)
	}
}

func sumOfDigit(x int64) int64 {
	r := x % 9
	if r == 0 {
		return 9
	}
	return r
}

func solve(n, d int64) (int, int) {
	ans := n
	x := n

	if n > 9 {
		ans = sumOfDigit(n)
		x = ans
	}

	for i := 0; i < 12; i++ {
		x += d
		if x > 9 {
			x = sumOfDigit(x)
			if x < ans {
				ans = x
			}
		}
	}
	return int(ans), count(n, d, ans, 0)
}

func count(n, d, min int64, c int) int {
	if c > 12 {
		return c
	}
	if n == min {
		return c
	}

	k1 := count(n+d, d, min, c+1)
	k2 := count(sum(n), d, min, c+1)
	if k1 < k2 {
		return k1
	}
	return k2
}

func sum(n int64) int64 {
	var res int64
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}
