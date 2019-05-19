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

	n := readNum(scanner)
	S := make([]string, n)

	for i := 0; i < n; i++ {
		for len(S[i]) == 0 {
			scanner.Scan()
			S[i] = compress(scanner.Bytes())
		}
	}
	//fmt.Fprintf(os.Stderr, "%v\n", S)
	fmt.Println(solve(S))
}

func compress(bs []byte) string {
	var i = 0
	var j = 0
	for j < len(bs) {
		if bs[j] == 'T' || bs[j] == 'F' {
			bs[i] = bs[j]
			i++
		}
		j++
	}

	return string(bs[:i])
}

func solve(S []string) int {
	n := len(S)
	cnt := make(map[string]int)

	for i := 0; i < n; i++ {
		cnt[S[i]]++
	}

	var res int

	for k, v := range cnt {
		if v <= res {
			continue
		}
		var x int
		for i := 0; i < len(k); i++ {
			if k[i] == 'T' {
				x++
			}
		}
		if x == v {
			res = max(res, v)
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
