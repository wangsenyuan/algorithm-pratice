package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	for i := 0; i < t; i++ {
		n, sig := readTwoNums(scanner)
		res, can := solve(n, sig)
		if !can {
			fmt.Println("-1")
		} else {
			fmt.Printf("%.9f", res[0])
			for j := 1; j < n; j++ {
				fmt.Printf(" %.9f", res[j])
			}
			fmt.Println()
		}
	}
}

func solve(n int, sig int) ([]float64, bool) {
	if n == 1 {
		if sig == 0 {
			return []float64{0}, true
		}
		return nil, false
	}

	// n * sig ^ 2 = sum((ai - u)^2)
	// make all ai as zero except a0 & a1
	// and a0 = -a1, then u would equal to 0
	// then n * sig ^ 2 = a0 ^ 2 + a1 ^ 2 => n * sig ^ 2 = 2 * a0 ^ 2
	x := int64(n) * int64(sig) * int64(sig)
	y := float64(x) / 2.0
	y = math.Sqrt(y)
	res := make([]float64, n)
	if y != 0 {
		res[0] = y
		res[1] = -y
	}
	return res, true
}
