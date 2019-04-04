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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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
		X := make([]int, n)
		Y := make([]int, n)
		for i := 0; i < n; i++ {
			X[i], Y[i] = readTwoNums(scanner)
		}
		a, b := solve(n, X, Y)
		if a > b {
			fmt.Println("invalid")
		} else {
			fmt.Printf("%d %d\n", a, b)
		}
	}
}

var table [][]bool

func init() {
	table = make([][]bool, 10)
	table[0] = []bool{true, true, true, true, true, true, false}
	table[1] = []bool{false, true, true, false, false, false, false}
	table[2] = []bool{true, true, false, true, true, false, true}
	table[3] = []bool{true, true, true, true, false, false, true}
	table[4] = []bool{false, true, true, false, false, true, true}
	table[5] = []bool{true, false, true, true, false, true, true}
	table[6] = []bool{true, false, true, true, true, true, true}
	table[7] = []bool{true, true, true, false, false, false, false}
	table[8] = []bool{true, true, true, true, true, true, true}
	table[9] = []bool{true, true, true, true, false, true, true}
}

func solve(n int, X []int, Y []int) (int, int) {
	bad := make([]bool, 7)
	isValid := func(mask int) bool {
		for i := 0; i < 7; i++ {
			bad[i] = mask&(1<<uint(i)) > 0
		}

		for i := 0; i < n; i++ {
			x := X[i]
			y := Y[i]
			var cnt = 0
			for j := 0; j < 7; j++ {
				if table[x][j] {
					// position j need to turn on
					if !bad[j] {
						// it is not bad
						cnt++
					}
				}

			}
			if cnt != y {
				return false
			}
		}
		return true
	}

	N := 1 << 7
	a := 7
	b := 0
	for mask := 0; mask < N; mask++ {
		if isValid(mask) {
			cnt := digitCount(mask)
			if cnt < a {
				a = cnt
			}
			if cnt > b {
				b = cnt
			}
		}
	}
	return a, b
}

func digitCount(mask int) int {
	var res = 0
	for mask > 0 {
		res += mask & 1
		mask >>= 1
	}
	return res
}
