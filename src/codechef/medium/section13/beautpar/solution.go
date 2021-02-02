package main

import (
	"bufio"
	"bytes"
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--

		n, k := readTwoNums(reader)

		A := readNNums(reader, n)

		res := solve(n, k, A)

		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func solve(n, k int, A []int) []int {

	sum := make([]int64, n+1)

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(A[i])
	}

	K := int64(k)

	if sum[n] < K {
		return nil
	}

	stack := make([]int, n+1)
	prev := make([]int, n+1)

	var p int

	for i := 1; i <= n; i++ {
		if sum[i] <= 0 {
			continue
		}
		j := search(p, func(j int) bool {
			return sum[stack[j]] >= sum[i]
		})
		if j > 0 {
			prev[i] = stack[j-1]
		}
		stack[j] = i

		if j == p {
			p++
		}
	}

	pivots := make([]int, 0, k)

	for i := n; i > 0; i = prev[i] {
		pivots = append(pivots, i)
		if len(pivots) == k {
			break
		}
	}
	if len(pivots) < k {
		return nil
	}

	reverse(pivots)

	ans := make([]int, k)
	var cur int
	for i := 0; i < k; i++ {
		ans[i] = pivots[i] - cur
		cur = pivots[i]
	}
	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
