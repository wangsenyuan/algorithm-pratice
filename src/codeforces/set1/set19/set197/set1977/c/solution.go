package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)

		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

type Pair struct {
	first  int
	second int
}

func solve(a []int) int {
	sort.Ints(a)
	n := len(a)
	var arr []Pair
	var j int
	for i := 1; i <= n; i++ {
		if i == n || a[i] > a[i-1] {
			arr = append(arr, Pair{a[i-1], i - j})
			j = i
		}
	}

	var best int

	// 这里有点难搞了
	// 不考虑复杂性，应该尝试所有的组合，看它们的lcm是否存在
	// 但这样子的复杂性是 2 ** m 的
	// cnt[i]表示可以被arr[i].first 整除的数的个数
	// 现在假设 arr[i].first 是这个set中的一个
	// 那么 lcm 必然是它的一个倍数
	// 然后考察这个倍数，如果这个倍数在后面不存在，且可以被计算出来，那么这个set就是可以的
	m := len(arr)
	if m == 1 {
		return 0
	}
	// 先检查这个arr

	last := arr[m-1].first

	for i := 0; i < m; i++ {
		if last%arr[i].first != 0 {
			// 那么lcm 肯定会超过 last
			return n
		}
	}

	find := func(expect int) (ok bool, cnt int) {
		l := 1
		for i := 0; i < m; i++ {
			if expect == arr[i].first {
				// lcm 出现了
				return false, 0
			}
			if expect%arr[i].first == 0 {
				l = lcm(l, arr[i].first)
				cnt += arr[i].second
			}
		}

		ok = l == expect

		return
	}

	// lcm肯定是last的一个因子
	for x := 1; x <= last/x; x++ {
		if last%x == 0 {
			ok, cnt := find(x)
			if ok {
				best = max(best, cnt)
			}
			if x*x == last {
				continue
			}
			ok, cnt = find(last / x)
			if ok {
				best = max(best, cnt)
			}
		}
	}

	return best
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
