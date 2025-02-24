package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	// pref[i] = a[1] + a[2] + ... + a[i]
	// flip(l..r) = pref[l-1] + r - (l - 1) - (pref[r] - pref[l-1])
	//            = 2 * pref[l-1] + r - (l - 1) - pref[r]
	//            = 2 * pref[l-1] - (l - 1) + r - pref[r]
	var suf int
	for i := n - 1; i >= 0; i-- {
		suf += a[i]
	}

	best := 0

	var pl, pr int

	for r, x := range a {
		suf -= x
		pr += x
		tmp := r + 1 - pr + pl + suf
		best = max(best, tmp)
		pl = max(pl, 2*pr-(r+1))
	}

	return best
}
