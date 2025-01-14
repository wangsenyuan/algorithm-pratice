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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d %d %d %d\n", res[0], res[1], res[2], res[3]))
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

type pair struct {
	first  int
	second int
}

func min_of(a, b pair) pair {
	if a.first*b.second <= a.second*b.first {
		return a
	}
	return b
}

const inf = 1 << 40

func solve(a []int) []int {

	sort.Ints(a)
	n := len(a)
	var arr []pair

	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		if i-j >= 2 {
			arr = append(arr, pair{a[j], i - j})
		}
	}

	// 对于给定的h
	// (w + w + h + h)**2 / (wh)
	// y =  4 * (w + h)**2 / (wh)  最小
	// y = (w * w + 2 * w * h + h * h) / (wh)
	// y = w / h + 2 + h / w 最小
	// 从上面看，感觉 w 和 h越接近越好

	ans := pair{-1, -1}

	res := make([]int, 4)

	for i, cur := range arr {
		if cur.second >= 4 {
			return []int{cur.first, cur.first, cur.first, cur.first}
		}
		if i > 0 {
			prev := arr[i-1]
			p := 2 * (prev.first + cur.first)
			p *= p
			s := prev.first * cur.first
			c := gcd(p, s)
			p /= c
			s /= c
			tmp := pair{p, s}
			if ans.first < 0 || min_of(tmp, ans) == tmp {
				ans = tmp
				res[0] = prev.first
				res[1] = prev.first
				res[2] = cur.first
				res[3] = cur.first
			}
		}
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
