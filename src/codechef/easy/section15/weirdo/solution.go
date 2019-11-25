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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		ss := make([]string, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			ss[i] = scanner.Text()
		}
		res := solve(n, ss)
		if res > INF {
			fmt.Println("Infinity")
		} else {
			fmt.Printf("%.7f\n", res)
		}
	}
}

const INF = 10000001

func solve(n int, ss []string) float64 {
	A := make([]string, 0, n)
	B := make([]string, 0, n)

	for _, s := range ss {
		if check(s) {
			A = append(A, s)
		} else {
			B = append(B, s)
		}
	}

	if len(A) == 0 {
		return 0
	}
	if len(B) == 0 {
		return INF
	}

	return calculate(A) / calculate(B)
}

func calculate(ss []string) float64 {
	xc := make([]int, 26)
	fxc := make([]int, 26)

	seen := make([]bool, 26)

	for _, s := range ss {
		for i := 0; i < 26; i++ {
			seen[i] = false
		}
		for i := 0; i < len(s); i++ {
			x := int(s[i] - 'a')
			fxc[x]++
			if !seen[x] {
				seen[x] = true
				xc[x]++
			}
		}
	}

	res := 1.0

	for i := 0; i < 26; i++ {
		if fxc[i] > 0 {
			res *= float64(xc[i]) / float64(pow(fxc[i], len(ss)))
		}
	}

	return res
}

func check(s string) bool {
	for i := 1; i < len(s); i++ {
		if !isVowel(s[i]) && (!isVowel(s[i-1]) || (i >= 2 && !isVowel(s[i-2]))) {
			return false
		}
	}
	return true
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'o' || x == 'e' || x == 'i' || x == 'u'
}

func pow(a, b int) int64 {
	res := int64(1)
	A := int64(a)

	for b > 0 {
		if b&1 == 1 {
			res *= A
		}
		A *= A
		b >>= 1
	}
	return res
}
