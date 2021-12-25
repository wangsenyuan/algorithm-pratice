package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
		res := solve(scanner.Text())
		fmt.Println(res)
	}
}

var cnt []int

func init() {
	cnt = make([]int, 26)
}

func solve(S string) int {
	for i := 0; i < 26; i++ {
		cnt[i] = 0
	}

	var num int

	for i := 0; i < len(S); i++ {
		x := int(S[i] - 'A')
		cnt[x]++
		if cnt[x] == 1 {
			num++
		}
	}
	sort.Ints(cnt)

	check := func(k int) int {
		avg := len(S) / k
		var res int
		for i := 25; i >= 0 && k > 0; i, k = i-1, k-1 {
			if cnt[i] < avg {
				res += avg - cnt[i]
			}
		}
		return res
	}

	res := math.MaxInt32
	for i := 1; i <= num; i++ {
		if len(S)%i == 0 {
			res = min(res, check(i))
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
