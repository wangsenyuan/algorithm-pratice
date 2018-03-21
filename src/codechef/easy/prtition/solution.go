package main

import (
	"bufio"
	"bytes"
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
		tmp := readNNums(scanner, 2)
		x, n := tmp[0], tmp[1]
		can, res := solve(n, x)
		if !can {
			fmt.Println("impossible")
		} else {
			var buf bytes.Buffer
			for i := 0; i < n; i++ {
				buf.WriteByte(byte(res[i] + '0'))
			}
			fmt.Printf("%s\n", buf.String())
		}
		t--
	}
}

func solve1(n int, x int) (bool, []int) {
	m := int64(n)
	sum := m*(m+1)/2 - int64(x)
	if sum%2 == 1 {
		return false, nil
	}

	var rec func(i int, sum0, sum1 int64) bool
	res := make([]int, n)
	rec = func(i int, sum0, sum1 int64) bool {
		if sum0+sum1 == sum {
			return sum0 == sum1
		}
		if sum0+sum1 > sum {
			return false
		}
		if i > n {
			return false
		}

		if i == x {
			return rec(i+1, sum0, sum1)
		}

		if sum0 <= sum1 {
			res[i-1] = 0
			tmp := rec(i+1, sum0+int64(i), sum1)
			if tmp {
				return true
			}
			res[i-1] = 1
			return rec(i+1, sum0, sum1+int64(i))
		}

		res[i-1] = 1
		tmp := rec(i+1, sum0, sum1+int64(i))
		if tmp {
			return true
		}
		res[i-1] = 0
		return rec(i+1, sum0+int64(i), sum1)
	}

	can := rec(1, 0, 0)
	if can {
		res[x-1] = 2
	}
	return can, res
}

func solve(n int, x int) (bool, []int) {
	m := int64(n)
	sum := m*(m+1)/2 - int64(x)
	if sum%2 == 1 {
		return false, nil
	}

	res := make([]int, n)
	sum /= 2
	for i := n; i >= 1 && sum > 0; i-- {
		if i == x {
			continue
		}

		if sum >= int64(i) && sum > int64(n+i) {
			res[i-1] = 1
			sum -= int64(i)
		} else if sum >= int64(i) && sum != int64(x+i) {
			res[i-1] = 1
			sum -= int64(i)
		}
	}

	if sum == 0 {
		res[x-1] = 2
		return true, res
	}
	return false, nil
}
