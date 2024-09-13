package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	res := solve(n, m, k)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))

	for _, x := range res {
		buf.WriteString(x)
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, m int, k int) []string {
	tot := 4*n*m - 2*n - 2*m
	if tot < k {
		return nil
	}
	var res []string

	if m == 1 {
		res = append(res, fmt.Sprintf("%d %s\n", min(n-1, k), "D"))
		k -= n - 1
		if k > 0 {
			res = append(res, fmt.Sprintf("%d %s\n", min(n-1, k), "U"))
			k -= n - 1
		}
	}

	if k > 0 {
		res = append(res, fmt.Sprintf("%d %s\n", min(m-1, k), "R"))
		k -= m - 1
	}
	if k > 0 {
		res = append(res, fmt.Sprintf("%d %s\n", min(m-1, k), "L"))
		k -= m - 1
	}
	// 第一行全部被处理了
	if k <= 0 {
		return res
	}

	res = append(res, fmt.Sprintf("%d %s\n", 1, "D"))
	k--

	for i := 1; i < n && k > 0; i++ {
		// 在当前行的最左端
		res = append(res, fmt.Sprintf("%d %s\n", min(m-1, k), "R"))
		k -= m - 1
		if k <= 0 {
			break
		}
		// 现在在当前行的最右边
		a := min(k/3, m-1)
		j := m - 1
		if a > 0 {
			res = append(res, fmt.Sprintf("%d %s\n", a, "UDL"))
			k -= a * 3
			j -= a
		}
		if k == 0 {
			break
		}
		if j > 0 {
			// 可以不用到第一列
			if k == 1 {
				res = append(res, "1 U\n")
			} else {
				// k == 2
				res = append(res, "1 UD\n")
			}
			break
		}
		// j == 0
		if i+1 < n {
			res = append(res, "1 D\n")
			k--
			continue
		}

		res = append(res, fmt.Sprintf("%d %s\n", min(n-1, k), "U"))
		k -= n - 1
	}

	return res
}
