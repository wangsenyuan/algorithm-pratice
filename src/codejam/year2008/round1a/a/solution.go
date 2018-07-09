package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// f, err := os.Open("./A-small-practice.in")
	f, err := os.Open("./A-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		n := readNum(scanner)

		A := readNNums(scanner, n)
		B := readNNums(scanner, n)
		fmt.Printf("Case #%d: %d\n", i, solve(A, B))
	}
}

func solve(A, B []int) int64 {
	sort.Ints(A)
	sort.Ints(B)

	n := len(A)
	var res int64
	for i := 0; i < n; i++ {
		res += int64(A[i]) * int64(B[n-1-i])
	}

	return res
}

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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}
