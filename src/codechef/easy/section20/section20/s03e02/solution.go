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
		S, _ := reader.ReadString('\n')
		res := solve(n, m, S)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(n, m int, S string) int {
	// if keep color x, how many moves made?
	if m == 0 {
		return solve0(n, S)
	}

	check := func(x int) int {
		// try to keep x only

		var a int
		for a < n && int(S[a]-'A') == x {
			a++
		}

		p := n
		for p > 0 && int(S[p-1]-'A') == x {
			p--
		}

		if p == 0 {
			return n
		}
		pq := make([]int, 0, n)

		res := a + n - p

		for i := a; i < p; i++ {
			y := int(S[i] - 'A')
			if y != x {
				continue
			}
			// y == x
			j := i
			for i < n && S[j] == S[i] {
				i++
			}

			pq = append(pq, i-j)
			i--
		}
		sort.Ints(pq)
		reverse(pq)
		if len(pq) > m-1 {
			pq = pq[:m-1]
		}
		for _, num := range pq {
			res += num
		}
		return res
	}

	var res int

	for i := 0; i < 26; i++ {
		res = max(res, check(i))
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func solve0(n int, S string) int {
	for i := 1; i < n; i++ {
		if S[i] != S[0] {
			return -1
		}
	}
	return n
}


func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
