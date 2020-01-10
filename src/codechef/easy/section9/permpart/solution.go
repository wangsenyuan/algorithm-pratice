package main

import (
	"bufio"
	"fmt"
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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

func solve(n int, A []int) int64 {
	var res int64
	f := make([]int64, 2*n+1)
	n = 0
	for i := 0; i < len(A); i++ {
		if A[i] > 2*len(A) {
			res++
			continue
		}
		A[n] = A[i]
		f[A[n]]++
		n++
	}

	dp := make([][]int64, 2*len(A)+1)
	dp[0] = make([]int64, 1+2*len(A))

	for cur := 1; cur <= 2*len(A); cur++ {
		mx := 2 * n / cur
		dp[cur] = make([]int64, mx+1)

		for st := 0; st <= mx; st++ {
			dp[cur][st] = dp[cur-1][st] + abs(f[cur]-int64(st))
		}

		for st := mx - 1; st >= 0; st-- {
			dp[cur][st] = minInt64(dp[cur][st], dp[cur][st+1])
		}
	}

	res += minInt64(dp[2*len(A)][0], dp[2*len(A)][1])

	return res
}

func minInt64(a, b int64) int64 {
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

func solve1(n int, A []int) int {
	sort.Ints(A)

	items := make([]*Item, 0, n)

	var res int

	for j, i := 0, 1; i <= n; i++ {
		if i < n && A[i] > 2*n {
			res = n - i
			if j < i {
				item := new(Item)
				item.num = A[i-1]
				item.cnt = i - j
				items = append(items, item)
			}
			break
		}

		if i == n || A[i] > A[i-1] {
			item := new(Item)
			item.num = A[i-1]
			item.cnt = i - j
			j = i
			items = append(items, item)
		}
	}

	m := len(items)

	for m > 0 {
		best := m
		var prev int
		var missing int
		for i := 0; i < m; i++ {
			item := items[i]
			missing += item.num - prev - 1
			best = min(best, missing+m-i-1)
			prev = item.num
			item.cnt--
		}

		res += best

		var j int
		for i := 0; i < m; i++ {
			if items[i].cnt == 0 {
				continue
			}
			items[j] = items[i]
			j++
		}
		m = j
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	num int
	cnt int
}
