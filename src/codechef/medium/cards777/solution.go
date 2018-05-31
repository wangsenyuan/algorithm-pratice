package main

import (
	"bufio"
	"fmt"
	"os"
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

	for t > 0 {
		tmp := readNNums(scanner, 3)
		res := solve(tmp[0], tmp[1], tmp[2])
		fmt.Printf("%.7f\n", res)
		t--
	}
}

func solve(r, b, p int) float64 {
	x := float64(r)
	y := float64(b)
	q := r + b - p

	var res float64
	for i := 0; i < r+b; i++ {
		if p == 0 {
			res += y
			break
		}
		if q == 0 {
			res += x
			break
		}
		w := x / (x + y)
		res += w
		x -= w
		y -= 1 - w
		p--
	}

	return res
}

func solve1(r, b, p int) float64 {
	return float64(p)*(float64(r)/float64(r+b)) + float64(r+b-p)*(float64(b)/float64(r+b))
}
