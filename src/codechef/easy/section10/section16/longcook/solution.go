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

	n := readNum(scanner)

	for n > 0 {
		n--
		first := readNNums(scanner, 2)
		second := readNNums(scanner, 2)
		fmt.Println(solve(first, second))
	}

}

var c []int64

func init() {
	c = make([]int64, 4801)
	s := 6

	for y := 0; y < 400; y++ {
		for m := 1; m <= 12; m++ {
			var d int

			if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
				d = 31
			} else if m == 4 || m == 6 || m == 9 || m == 11 {
				d = 30
			} else {
				if y == 0 || y%4 == 0 && y%100 != 0 {
					d = 29
				} else {
					d = 28
				}
			}

			d1, d2 := 0, d-1
			for (s+d1)%7 != 5 {
				d1++
			}
			for (s+d2)%7 != 0 {
				d2--
			}
			d2 -= 7

			c[y*12+m] = c[y*12+m-1]
			if d1+10 > d2 {
				c[y*12+m]++
			}
			s = (s + d) % 7
		}
	}
}

func solve(first, second []int) int64 {

	check := func(m, y int) int64 {
		r := c[4800] * int64(y/400)
		r += c[(y%400)*12+m-1]
		return r
	}

	second[0]++
	if second[0] > 12 {
		second[0] = 1
		second[1]++
	}

	return check(second[0], second[1]) - check(first[0], first[1])
}

// sunday, monday, tusday, wendsday, thursday, friday, saturday
