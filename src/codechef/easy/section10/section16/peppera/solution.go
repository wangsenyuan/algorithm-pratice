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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		G := make([]string, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			G[i] = scanner.Text()
		}
		res := solve(n, G)

		for i := 0; i < n; i++ {
			fmt.Println(res[i])
		}
	}
}

// 75 * 150 < 15000
const MAX_D = 15000
const MAX_N = 151

var can [][]bool
var diff []int
var good []bool
var pick []bool

func init() {
	can = make([][]bool, MAX_N)

	for i := 0; i < MAX_N; i++ {
		can[i] = make([]bool, MAX_D)
	}
	diff = make([]int, MAX_N)
	good = make([]bool, MAX_N)
	pick = make([]bool, MAX_N)
}

func reset() {
	for i := 0; i < MAX_N; i++ {
		for j := 0; j < MAX_D; j++ {
			can[i][j] = false
		}
	}

	for i := 0; i < MAX_N; i++ {
		diff[i] = 0
		good[i] = false
		pick[i] = false
	}
}

func solve(n int, G []string) []string {
	reset()

	for i := 0; i < n; i++ {
		var first, second int

		for j := 0; j < n/2; j++ {
			if G[i][j] == '1' {
				first++
			}
		}

		for j := n / 2; j < n; j++ {
			if G[i][j] == '1' {
				second++
			}
		}

		diff[i] = abs(first - second)
		good[i] = first >= second
	}

	can[0][0] = true
	var sum int
	for i := 1; i <= n; i++ {
		d := diff[i-1]

		sum += d

		for x := 0; x <= sum; x++ {
			can[i][x] = can[i-1][x]
			if x >= d && !can[i][x] {
				can[i][x] = can[i-1][x-d]
			}
		}
	}

	best := sum / 2

	for !can[n][best] {
		best--
	}

	for best > 0 {
		for i := 1; i <= n; i++ {
			if can[i][best] {
				pick[i-1] = true
				best -= diff[i-1]
				break
			}
		}
	}

	res := make([]string, n)

	for i := 0; i < n; i++ {
		if pick[i] != good[i] {
			res[i] = reverse(G[i])
		} else {
			res[i] = G[i]
		}
	}

	return res
}

func reverse(s string) string {
	bs := []byte(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
