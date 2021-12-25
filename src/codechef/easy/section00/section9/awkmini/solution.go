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
		scanner.Scan()
		s := scanner.Text()
		r := solve(len(s), s)
		fmt.Println(r)
	}
}

func solve(n int, s string) int64 {
	var boy, girl int

	for i := 0; i < n; i++ {
		if s[i] == 'b' {
			boy++
		} else {
			girl++
		}
	}

	if boy < girl {
		boy, girl = girl, boy
	}

	buf := make([]byte, n)

	var i int

	for i < (boy-girl)/2 {
		buf[i] = 'b'
		i++
	}

	for j := 0; j < girl; j++ {
		buf[i] = 'b'
		i++
		buf[i] = 'g'
		i++
	}

	for i < n {
		buf[i] = 'b'
		i++
	}

	var ans int64

	var bc, gc int
	var bi, gi int64

	for i := n - 1; i >= 0; i-- {
		if buf[i] == 'g' {
			gc++
			gi += int64(i)
			ans += abs(bi - int64(bc)*int64(i))
		} else {
			bc++
			bi += int64(i)
			ans += abs(gi - int64(gc)*int64(i))
		}
	}

	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
