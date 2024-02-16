package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		s := readString(reader)
		i := strings.IndexByte(s, ' ')
		l := s[:i]
		r := s[i+1:]
		res := solve(l, r)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Println(buf.String())
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

var lcms []int

var ind []int
var next [][]int

func init() {
	var arr []int
	arr = append(arr, 1)
	for state := 1; state < (1 << 8); state++ {
		var nums []int
		for i := 0; i < 8; i++ {
			if (state>>i)&1 == 1 {
				nums = append(nums, i+2)
			}
		}
		tmp := nums[0]
		for i := 1; i < len(nums); i++ {
			tmp = lcm(tmp, nums[i])
		}
		arr = append(arr, tmp)
	}

	sort.Ints(arr)

	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}

	lcms = arr[:n]

	mx := arr[n-1]

	ind = make([]int, mx+1)

	for i := 0; i < n; i++ {
		ind[arr[i]] = i
	}

	next = make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 10)
		next[i][0] = i
		next[i][1] = i
		for j := 2; j < 10; j++ {
			next[i][j] = ind[lcm(arr[i], j)]
		}
	}
}

func solve(l string, r string) int {
	n := len(r)

	if len(l) < n {
		l = strings.Repeat("0", n-len(l)) + l
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, len(lcms))
		for j := 0; j < len(lcms); j++ {
			dp[i][j] = make([]int, 2520)
			for k := 0; k < 2520; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i int, j int, rem int, lo bool, hi bool) (res int)

	dfs = func(i int, j int, rem int, lo bool, hi bool) (res int) {
		if i == n {
			if rem%lcms[j] == 0 {
				return 1
			}
			return 0
		}
		if !lo && !hi {
			p := &dp[i][j][rem]

			if *p >= 0 {
				return *p
			}

			defer func() {
				*p = res
			}()
		}

		x, y := 0, 9
		if lo {
			x = int(l[i] - '0')
		}
		if hi {
			y = int(r[i] - '0')
		}

		for d := x; d <= y; d++ {
			res += dfs(i+1, next[j][d], (rem*10+d)%2520, lo && d == int(l[i]-'0'), hi && d == int(r[i]-'0'))
		}
		return
	}

	return dfs(0, 0, 0, true, true)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
