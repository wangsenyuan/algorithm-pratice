package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k, d := readThreeNums(reader)
	a := readNNums(reader, n)
	res := solve(k, d, a)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(k int, d int, a []int) bool {
	sort.Ints(a)

	if a[k-1]-a[0] > d {
		return false
	}

	n := len(a)

	dp := make([]int, n+1)
	var cnt int
	dp[0] = 1
	for i, j := k, 0; i <= n; i++ {
		cnt += dp[i-k]
		for a[i-1]-a[j] > d {
			cnt -= dp[j]
			j++
		}
		if cnt > 0 {
			dp[i] = 1
		}
	}

	return dp[n] == 1
}

func solve1(k int, d int, a []int) bool {
	sort.Ints(a)
	n := len(a)

	dp := make(fenwick, n+10)
	for i := k; i <= n; i++ {
		j := sort.SearchInts(a, a[i-1]-d)
		// a[j] >= a[i-1] - d
		if j == 0 {
			dp.update(i, 1)
		} else if i-j >= k {
			// j > 0
			cnt := dp.query(j, i-k)
			if cnt > 0 {
				dp.update(i, 1)
			}
		}
	}

	return dp.query(n, n) == 1
}

type fenwick []int

func (f fenwick) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}
