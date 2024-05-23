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
		s := fmt.Sprintf("%v", res)

		buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
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

func solve(a []int) []int {
	n := len(a)

	pref := make([]int, n)

	// L[i] 是左边和a[i]不相同的第一个下标
	L := make([]int, n)
	L[0] = -1
	for i := 0; i < n; i++ {
		pref[i] = a[i]
		if i > 0 {
			pref[i] += pref[i-1]
			if a[i] != a[i-1] {
				L[i] = i - 1
			} else {
				L[i] = L[i-1]
			}
		}
	}
	R := make([]int, n)
	R[n-1] = n
	for i := n - 2; i >= 0; i-- {
		if a[i] != a[i+1] {
			R[i] = i + 1
		} else {
			R[i] = R[i+1]
		}
	}

	getFront := func(i int) int {
		if i == 0 {
			// 前面的都一样
			return n
		}
		if a[i-1] > a[i] {
			return 1
		}
		if L[i-1] == -1 {
			return n
		}
		tmp := pref[i-1]
		if tmp <= a[i] {
			// 前面全部加起来也无法吃掉i
			return n
		}
		k := search(0, L[i-1], func(k int) bool {
			return tmp-pref[k] <= a[i]
		})
		// tmp - pref[k] > a[i]
		return i - k
	}

	getBack := func(i int) int {
		if i == n-1 {
			return n
		}
		if a[i+1] > a[i] {
			return 1
		}
		if R[i+1] == n {
			return n
		}
		if pref[n-1]-pref[i] <= a[i] {
			return n
		}
		k := search(R[i+1], n-1, func(k int) bool {
			return pref[k]-pref[i] > a[i]
		})
		return k - i
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = -1
		a := getFront(i)
		b := getBack(i)
		a = min(a, b)
		if a < n {
			res[i] = a
		}
	}

	return res
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
