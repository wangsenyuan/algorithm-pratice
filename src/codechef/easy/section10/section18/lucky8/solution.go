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
	for tc > 0 {
		tc--
		L, _ := reader.ReadString(' ')
		R, _ := reader.ReadString('\n')
		fmt.Println(solve(L[:len(L)-1], R[:len(R)-1]))
	}
}

func solve(L, R string) int {
	L = leftPadZero(L, len(R)-len(L))

	n := len(R)

	var n4, n7 int

	var i int

	for i < n && L[i] == R[i] {
		if L[i] == '4' {
			n4++
		} else if L[i] == '7' {
			n7++
		}
		i++
	}

	var best int

	if i < n {
		// L[i] < R[i]

		best = max(best, check(int(L[i]-'0')+1, int(R[i]-'0')-1, n4, n7, n-i-1))

		a, b := n4, n7

		for j := i; j < n; j++ {
			if L[j] == '4' {
				a++
			} else if L[j] == '7' {
				b++
			}
			if j+1 < n {
				best = max(best, check(int(L[j+1]-'0')+1, 9, a, b, n-(j+2)))
			}
		}
		best = max(best, a*b)

		a, b = n4, n7

		for j := i; j < n; j++ {
			if R[j] == '4' {
				a++
			} else if R[j] == '7' {
				b++
			}
			if j+1 < n {
				best = max(best, check(0, int(R[j+1]-'0')-1, a, b, n-(j+2)))
			}
		}

		best = max(best, a*b)
	} else {
		best = n4 * n7
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func leftPadZero(s string, n int) string {
	var buf bytes.Buffer

	for n > 0 {
		buf.WriteByte('0')
		n--
	}
	buf.WriteString(s)
	return buf.String()
}

func check(l, r int, cnt4, cnt7, n int) int {
	if l > r {
		return 0
	}

	ans := count(cnt4, cnt7, n)
	if l <= 4 && 4 <= r {
		ans = max(ans, count(cnt4+1, cnt7, n))
	}
	if l <= 7 && 7 <= r {
		ans = max(ans, count(cnt4, cnt7+1, n))
	}

	return ans
}

func count(cnt4, cnt7 int, n int) int {
	k := (cnt7 + n - cnt4) / 2
	if k >= 0 && k <= n {
		return (cnt4 + k) * (cnt7 + n - k)
	}

	k++
	if k >= 0 && k <= n {
		return (cnt4 + k) * (cnt7 + n - k)
	}

	return max((cnt4+n)*cnt7, cnt4*(cnt7+n))
}
