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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
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

func solve(a []int) int {
	n := len(a)
	var res = 1 << 60

	if n%2 == 0 {
		res = 0
		for i := 0; i < n; i += 2 {
			res = max(res, a[i+1]-a[i])
		}
	} else {
		pref := make([]int, n)
		pref[0] = 1
		// 前面偶数个，后面偶数个， 前后缀
		for i := 1; i < n; i += 2 {
			pref[i/2] = a[i] - a[i-1]
			if i/2 > 0 {
				pref[i/2] = max(pref[i/2], pref[i/2-1])
			}
		}
		suf := 1
		for i := n - 1; i > 0; i -= 2 {
			res = min(res, max(suf, pref[(i-1)/2]))
			suf = max(suf, a[i]-a[i-1])
		}
		res = min(res, suf)
	}

	return res

}
