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
		x, n, m := readThreeNums(reader)
		H := readNNums(reader, n)
		A := readNNums(reader, m)
		res := solve(x, H, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(x int, H []int, A []int) int {
	for i := 0; i < len(H); i++ {
		H[i] = x - H[i]
	}
	sort.Ints(H)
	reverse(H)

	n := len(H)

	bit := make([]int, n+1)

	get := func(p int) int {
		var res int
		for ; p > 0; p -= p & -p {
			res += bit[p]
		}
		return res
	}

	set := func(p int, v int) {
		for ; p <= n; p += p & -p {
			bit[p] += v
		}
	}

	findLowerBound := func(v int) int {
		l, r := 1, n
		for l <= r {
			mid := (l + r) / 2
			tmp := H[mid-1] + get(mid)
			if tmp <= v {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return l
	}

	findUpperBound := func(v int) int {
		l, r := 1, n
		for l <= r {
			mid := (l + r) / 2
			tmp := H[mid-1] + get(mid)
			if tmp >= v {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return r
	}

	for i := 0; i < len(A); i++ {
		cur := A[i]
		ith := H[cur-1] + get(cur)
		if ith <= 0 {
			return i
		}

		lb := findLowerBound(ith)
		up := findUpperBound(ith)

		set(1, -1)
		set(lb, 1)

		cur -= lb

		set(up-cur, -1)
		set(up+1, 1)
	}

	return len(A)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
