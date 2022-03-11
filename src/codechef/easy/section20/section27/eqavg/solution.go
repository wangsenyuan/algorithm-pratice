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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(k, A)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(k int, A []int) []int {
	cnt := make(map[int]int)

	for _, num := range A {
		cnt[num]++
	}

	if len(cnt) > k {
		return nil
	}

	n := len(A)

	ans := make([]int, k)

	if n%k == 0 {
		var it int
		for c, f := range cnt {
			if f%(n/k) != 0 {
				return nil
			}
			for f > 0 {
				ans[it] = c
				it++
				f -= n / k
			}
		}
		// all f % k == 0
	} else {
		p1, p2 := 0, n%k
		reserve := make(map[int]int)
		for v, f := range cnt {
			x := decompose(f, n/k)
			if x == nil {
				return nil
			}
			for x[0] > 0 && p1 < n%k {
				ans[p1] = v
				p1++
				x[0]--
			}
			for x[2] > 0 && p2 < k {
				ans[p2] = v
				p2++
				x[2]--
			}
			if x[0] > 0 || x[2] > 0 {
				return nil
			}
			reserve[v] = x[1]
		}

		for v, f := range reserve {
			rem := n%k - p1
			for f > 0 && rem >= n/k {
				rem -= n / k
				f--
				for i := 0; i < n/k; i++ {
					ans[p1] = v
					p1++
				}
			}
			rem = k - p2
			for f > 0 && rem >= n/k+1 {
				rem -= n/k + 1
				f--
				for i := 0; i < n/k+1; i++ {
					ans[p2] = v
					p2++
				}
			}
		}
	}

	B := make([]int, n)
	copy(B, ans)

	for i := k; i < n; i++ {
		B[i] = B[i-k]
	}

	return B
}

func decompose(freq int, a int) []int {
	A := int64(a)
	F := int64(freq)
	top := int(F / (A * (A + 1)))
	x := make([]int, 3)
	for x[1] = top; x[1] >= 0; x[1]-- {
		rem := F - int64(x[1])*A*(A+1)
		for x[0] = 0; int64(x[0])*(A+1) <= rem; x[0]++ {
			tmp := rem - int64(x[0])*(A+1)
			if tmp%A == 0 {
				x[2] = int(tmp / A)
				return x
			}
		}
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
