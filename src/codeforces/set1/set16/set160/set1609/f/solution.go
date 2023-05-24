package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		H := readNInt64s(reader, n)
		res := solve(H)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

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

func solve(H []int64) int64 {
	n := len(H)
	h := make([]int64, n+1)
	copy(h[1:], H)
	L := make([]int, n+1)
	stack := make([]int, n+2)
	top := 1
	for i := 1; i <= n; i++ {
		for top > 1 && h[stack[top]] >= h[i] {
			top--
		}
		L[i] = stack[top]
		top++
		stack[top] = i
	}

	top = 1
	R := make([]int, n+2)
	R[n+1] = n + 1
	stack[1] = n + 1
	for i := n; i > 0; i-- {
		for top > 1 && h[stack[top]] >= h[i] {
			top--
		}
		R[i] = stack[top]
		top++
		stack[top] = i
	}

	var ans int64

	var sum int64

	for i := 1; i <= n; i++ {
		l, r := i, i+1
		var x, y int64
		for l > 0 {
			x = max(x, h[l]*int64(i-L[l]))
			l = L[l]
		}
		for r <= n {
			y = max(y, h[r]*int64(R[r]-i-1))
			r = R[r]
		}
		sum = max(sum, x)
		ans = max(ans, sum+y)
	}

	for i := 1; i <= n; i++ {
		l, r := L[i], R[i]
		id := i
		w := r - l - 1
		for {
			ans = max(ans, h[id]*int64(r-l-1)+(h[i]-h[id])*int64(w))
			if l <= 0 && r > n {
				break
			}
			if r > n || l > 0 && h[l] > h[r] {
				id = l
				l = L[l]
			} else {
				id = r
				r = R[r]
			}
		}
	}

	stack[0] = 0

	for i := 1; i <= n; i++ {
		l, r := i, i
		top = 0
		for l > 0 {
			top++
			stack[top] = l
			l = L[l]
		}

		for {
			for top > 0 && h[r]+h[stack[top]] <= h[i] {
				top--
			}
			l = stack[top]
			ans = max(ans, h[l]*int64(R[i]-L[l]-1)+(h[i]-h[l])*int64(r-L[i]-1))
			ans = max(ans, h[L[l]]*int64(R[i]-L[L[l]]-1)+h[r]*int64(R[r]-L[i]-1))
			ans = max(ans, (h[i]-h[r])*int64(R[i]-L[l]-1)+h[r]*int64(R[r]-L[i]-1))
			if R[r] > n {
				break
			}
			r = R[r]
		}
	}

	return ans
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func min(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}
