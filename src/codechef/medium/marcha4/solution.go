package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNum(scanner *bufio.Scanner) (a int, b int) {
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

	for t > 0 {
		n, k := readTwoNum(scanner)
		begin, end := solve(n, k)
		fmt.Println(formatBegin(begin, k), formatEnd(end, k))
		t--
	}
}

func formatBegin(x, k int) string {
	s := strconv.Itoa(x)
	if len(s) == k {
		return s
	}
	var buf bytes.Buffer
	buf.WriteString(s)
	for i := len(s); i < k; i++ {
		buf.WriteByte('0')
	}
	return buf.String()
}

func formatEnd(x, k int) string {
	s := strconv.Itoa(x)
	if len(s) == k {
		return s
	}
	var buf bytes.Buffer
	for i := len(s); i < k; i++ {
		buf.WriteByte('0')
	}
	buf.WriteString(s)

	return buf.String()
}

func solve(n int, k int) (int, int) {
	m := float64(n)
	begin := getFirstK(m, k)
	end := getLastK(n, k)
	return begin, end
}

func getFirstK(m float64, k int) int {
	pf := m * math.Log10(m)

	df := pf - math.Floor(pf)

	de := math.Pow(10, df)

	digits := math.Pow10(k - 1)

	return int(digits * de)
}

func getLastK(n int, k int) int {
	c := int64(fastPow(10, k))

	res := int64(1)

	x := int64(n) % c

	y := x

	for y > 0 {
		if y%2 == 1 {
			res = (res * x) % c
		}
		y = y >> 1
		x = (x * x) % c
	}
	return int(res)
}

func fastPow(a, b int) int {
	if b == 0 {
		return 1
	}
	c := fastPow(a, b/2)
	if b%2 == 1 {
		return c * c * a
	}
	return c * c
}
