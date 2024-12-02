package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func process(reader *bufio.Reader) []int {
	line := readNNums(reader, 4)
	n, L, R, K := line[0], line[1], line[2], line[3]
	a := readNNums(reader, n)
	return solve(a, L, R, K)
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

type pair struct {
	fi int
	sd int
}

func solve(a []int, L int, R int, K int) []int {
	n := len(a)

	if n == 1 {
		return []int{L}
	}

	suf := make([]pair, n)
	suf[n-1] = pair{L, R}

	for i := n - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			// 可以选择一个足够小的数，但是必须保证高位有空间可以用
			suf[i] = pair{max(suf[i+1].fi-K, L), suf[i+1].sd - 1}
		} else if a[i] > a[i+1] {
			suf[i] = pair{suf[i+1].fi + 1, min(suf[i+1].sd+K, R)}
		} else {
			suf[i] = suf[i+1]
		}
		if suf[i].fi > suf[i].sd {
			return nil
		}
	}
	ans := make([]int, n)
	ans[0] = suf[0].fi

	for j := 1; j < n; j++ {
		if a[j] == a[j-1] {
			ans[j] = ans[j-1]
			continue
		}
		l, r := ans[j-1], ans[j-1]
		if a[j] > a[j-1] {
			l++
			r = min(r+K, R)
		} else {
			l = max(l-K, L)
			r--
		}
		l = max(l, suf[j].fi)
		r = min(r, suf[j].sd)
		if l > r {
			return nil
		}
		ans[j] = l
	}
	return ans
}

func abs(num int) int {
	return max(num, -num)
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}
