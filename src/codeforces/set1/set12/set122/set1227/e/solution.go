package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, t, res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", t))
	for _, row := range res {
		buf.WriteString(row)
		buf.WriteByte('\n')
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

func process(reader *bufio.Reader) ([]string, int, []string) {
	n, m := readTwoNums(reader)
	region := make([]string, n)
	for i := range n {
		region[i] = readString(reader)[:m]
	}

	t, res := solve(region)

	return region, t, res
}

func solve(region []string) (int, []string) {

	n := len(region)
	m := len(region[0])

	cnt := make([][]int, n+1)
	sum := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		sum[i] = make([]int, m+1)
		cnt[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
			if region[i][j] == 'X' {
				sum[i][j]++
			}
		}
	}

	onFire := func(top int, left int, sz int) bool {
		if top+sz > n || left+sz > m {
			// can't burn outside
			return false
		}
		num := sum[top+sz-1][left+sz-1]
		if top > 0 {
			num -= sum[top-1][left+sz-1]
		}
		if left > 0 {
			num -= sum[top+sz-1][left-1]
		}
		if top > 0 && left > 0 {
			num += sum[top-1][left-1]
		}
		return num == sz*sz
	}

	check := func(T int) bool {
		if T == 0 {
			return true
		}
		// 是否能够用T来覆盖
		sz := 2*T + 1

		for i := 0; i < n; i++ {
			clear(cnt[i])
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if onFire(i, j, sz) {
					cnt[i][j]++
					cnt[i+sz][j]--
					cnt[i][j+sz]--
					cnt[i+sz][j+sz]++
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i > 0 {
					cnt[i][j] += cnt[i-1][j]
				}
				if j > 0 {
					cnt[i][j] += cnt[i][j-1]
				}
				if i > 0 && j > 0 {
					cnt[i][j] -= cnt[i-1][j-1]
				}
				if (region[i][j] == 'X') != (cnt[i][j] > 0) {
					return false
				}
			}
		}
		return true
	}

	l, r := 0, min(n+1, m+1)/2+1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	t := r - 1

	if t == 0 {
		return t, region
	}

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			res[i][j] = '.'
		}
	}

	sz := 2*t + 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 不管这里cnt的情况
			if onFire(i, j, sz) {
				res[i+t][j+t] = 'X'
			}
		}
	}

	return t, toStringArray(res)
}

func toStringArray(arr [][]byte) []string {
	res := make([]string, len(arr))
	for i := range len(arr) {
		res[i] = string(arr[i])
	}
	return res
}
