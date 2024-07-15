package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(p []int) int {
	// 先要找到，最长的那个山谷
	valleies := make([]Valley, 3)

	n := len(p)

	update := func(cur Valley) {
		if cur.Distance() >= valleies[0].Distance() {
			valleies[2] = valleies[1]
			valleies[1] = valleies[0]
			valleies[0] = cur
		} else if cur.Distance() >= valleies[1].Distance() {
			valleies[2] = valleies[1]
			valleies[1] = cur
		} else if cur.Distance() >= valleies[2].Distance() {
			valleies[2] = cur
		}
	}

	for j, i := 0, 1; i <= n; i++ {
		if i == n || p[i] > p[i-1] {
			cur := Valley{i - 1, j, i - 1}
			update(cur)
			j = i
		}
	}

	for i, j := n-2, n-1; i >= -1; i-- {
		if i < 0 || p[i] > p[i+1] {
			cur := Valley{i + 1, i + 1, j}
			update(cur)
			j = i
		}
	}

	var res int
	// L[i] 表示从i出发，可以进行多少次操作
	L := make([]int, n)
	for i := 0; i < n; i++ {
		// 可以进行一次操作
		L[i] = 1
		if i > 0 && p[i-1] < p[i] {
			L[i] = L[i-1] + 1
		}
	}

	var r int
	for i := n - 1; i >= 0; i-- {
		if i+1 < n && p[i] > p[i+1] {
			r++
		} else {
			r = 1
		}
		u, v := L[i], r

		if min(u, v) == 1 {
			continue
		}

		// 如果i是alice选的第一个位置
		ok := true

		for j := 0; j < len(valleies); j++ {
			y := valleies[j].id

			if y < i {
				if valleies[j].r < i && valleies[j].Distance() >= max(u, v) {
					ok = false
					break
				}
				if valleies[j].r == i && u%2 == 0 && u >= v {
					// 如果 u < v, alice可以选择另外一边
					ok = false
					break
				}
			} else {
				// y 在i的后面
				if valleies[j].l > i && valleies[j].Distance() >= max(u, v) {
					ok = false
					break
				}
				if valleies[j].l == i && v%2 == 0 && u <= v {
					ok = false
					break
				}
			}
		}

		if ok && max(u, v)/2*2 < min(u, v) {
			// bob 会选择长边种的，且从距离u偶数的位置开始
			// 这样子如果alice也选择这条边，他肯定会输
			res++
		}
	}

	return res
}

type Valley struct {
	id int
	l  int
	r  int
}

func (this Valley) Distance() int {
	return this.r - this.l + 1
}
