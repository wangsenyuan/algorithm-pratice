package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		nums := readNInt64s(reader, 3)
		res := solve(nums[0], nums[1], nums[2])
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(A int64, B int64, k int64) int64 {
	ca := count(A - 1)
	cb := count(B)
	cnt := cb - ca
	if cnt < k {
		return -1
	}
	// cnt >= k

	l, r := A, B
	for l < r {
		mid := (l + r) / 2
		tmp := count(mid)
		if tmp >= ca+k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

const LCM = 63

func count(num int64) int64 {
	var arr []int
	for num > 0 {
		arr = append(arr, int(num%10))
		num /= 10
	}

	reverse(arr)

	n := len(arr)

	dp := make([][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, LCM)
		for j := 0; j < LCM; j++ {
			dp[i][j] = make([]int64, 16)
			for k := 0; k < 16; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	var dfs func(i int, eq int, rem int, mask int, prev5 bool) int64

	dfs = func(i int, eq int, rem int, mask int, prev5 bool) int64 {
		if i == n {
			for d := 3; d <= 9; d += 2 {
				if d != 5 && mask&(1<<(((d-1)/2)-1)) > 0 && rem%d == 0 {
					return 0
				}
			}
			if prev5 {
				return 0
			}
			return 1
		}
		tmp := dp[i][rem][mask]
		if eq == 0 && tmp >= 0 {
			return tmp
		}
		tmp = 0
		if mask == 0 {
			tmp += dfs(i+1, 0, 0, 0, false)
		}
		till := 9
		if eq > 0 {
			till = arr[i]
		}
		for x := 3; x <= till; x += 2 {
			var next_eq int
			if eq == 1 && x == till {
				next_eq = 1
			}
			next_rem := (rem*10 + x) % LCM
			next_mask := mask | (1 << (((x - 1) / 2) - 1))
			tmp += dfs(i+1, next_eq, next_rem, next_mask, x == 5)
		}
		dp[i][rem][mask] = tmp
		return tmp
	}

	return dfs(0, 1, 0, 0, false)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
