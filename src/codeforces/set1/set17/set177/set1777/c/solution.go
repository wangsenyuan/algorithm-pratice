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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, m)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const N = 100000 + 10

var factors [N][]int

func init() {

	for i := 2; i < N; i++ {
		factors[i] = getFactors(i)
	}
}

func getFactors(num int) []int {
	var res []int
	res = append(res, 1, num)

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			res = append(res, i)
			if i*i != num {
				res = append(res, num/i)
			}
		}
	}
	sort.Ints(res)
	return res
}

func solve(a []int, m int) int {
	if m == 1 {
		return 0
	}
	// 如果 a[i] % a[j] = 0, 那么就没有必要使用a[j]
	// 只需要检查m的后半部分就可以了
	//  比如 x和y >= m / 2, 如果存在某个数 a[i] % lcm(x, y) = 0 的话，
	// 不是要最少的，而是要最接近的
	// x 和 x-1最好的选择是 x和x-1
	n := len(a)
	sort.Ints(a)
	reverse(a)

	// 首先找到一个区间，能够cover所有的m
	freq := make([]int, m+1)
	var cnt int

	add := func(num int) {
		for i := 1; i < len(factors[num]); i++ {
			x := factors[num][i]
			if x > m {
				break
			}
			freq[x]++
			if freq[x] == 1 {
				cnt++
			}
		}
	}

	rem := func(num int) {
		for i := 1; i < len(factors[num]); i++ {
			x := factors[num][i]
			if x > m {
				break
			}
			freq[x]--
			if freq[x] == 0 {
				cnt--
			}
		}
	}

	best := 1 << 30
	for l, r := 0, 0; r < n; r++ {
		add(a[r])

		for cnt == m-1 {
			rem(a[l])
			if cnt < m-1 {
				add(a[l])
				break
			}
			l++
		}
		if cnt == m-1 {
			best = min(best, a[l]-a[r])
		}
	}

	if cnt < m-1 {
		return -1
	}

	return best
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
