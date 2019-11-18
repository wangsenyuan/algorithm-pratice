package main

import (
	"bufio"
	"fmt"
	"math/cmplx"
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
		line := readNNums(scanner, 6)
		x1, y1, r1, x2, y2, r2 := line[0], line[1], line[2], line[3], line[4], line[5]
		res := solve(x1, y1, r1, x2, y2, r2)
		if len(res) == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			fmt.Printf("%.7f %.7f\n", res[0], res[1])
		}
	}
}

func solve(x1, y1, r1, x2, y2, r2 int) []float64 {
	if x1 == x2 && y1 == y2 {
		if r2 >= r1 {
			return nil
		}
		// r2 < r1
		return []float64{float64(x1 + r1), float64(y1)}
	}

	a := complex(float64(x1), float64(y1))
	b := complex(float64(x2), float64(y2))

	line := b - a

	_, theta := cmplx.Polar(line)

	// may or may not
	theta += 1.0
	rr := cmplx.Rect(float64(r1), theta)
	dd := a + rr

	// if dd in the cycle 2, then no point
	pp := dd - b

	pr, _ := cmplx.Polar(pp)

	R2 := float64(r2)

	if pr < R2 || (pr-R2) < 1e-7 {
		return nil
	}

	x := real(dd)
	y := imag(dd)

	return []float64{x, y}
}
