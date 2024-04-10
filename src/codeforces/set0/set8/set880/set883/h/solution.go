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

	readNum(reader)
	s := readString(reader)
	res := solve(s)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i, x := range res {
		buf.WriteString(x)
		if i+1 == len(res) {
			buf.WriteByte('\n')
		} else {
			buf.WriteByte(' ')
		}
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string) []string {
	n := len(s)

	freq := make(map[byte]int)

	for i := 0; i < n; i++ {
		freq[s[i]]++
	}

	var odd int

	for _, v := range freq {
		odd += v & 1
	}

	if odd <= 1 {
		return prepare(freq, n, 1)
	}

	// n % odd != 0

	for x := odd; x < n; x++ {
		diff := x - odd
		// 单个串必须是奇数长
		if n%x == 0 && diff%2 == 0 && (n/x)%2 == 1 {
			return prepare(freq, n, x)
		}
	}

	return prepare(freq, n, n)
}

type Byte struct {
	val byte
	cnt int
}

func prepare(freq map[byte]int, n, k int) []string {
	m := n / k
	// 每个字符中有m长

	var arr []*Byte
	for x, v := range freq {
		arr = append(arr, &Byte{x, v})
	}

	// 把奇数排在前面
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].cnt%2 > arr[j].cnt%2 || arr[i].cnt%2 == arr[j].cnt%2 && arr[i].cnt > arr[j].cnt
	})
	var j int
	res := make([][]byte, k)
	for i := 0; i < k; i++ {
		res[i] = make([]byte, m)
		if m%2 == 1 {
			for arr[j].cnt == 0 {
				j = (j + 1) % len(arr)
			}
			if arr[j].cnt%2 == 1 {
				res[i][m/2] = arr[j].val
				arr[j].cnt--
				j = (j + 1) % len(arr)
			} else {
				res[i][m/2] = arr[j].val
				arr[j].cnt--
			}
		}
	}

	ans := make([]string, k)
	// 现在全部是偶数了
	j = 0
	for i := 0; i < k; i++ {
		l, r := m/2, m/2
		if m%2 == 1 {
			l--
			r++
		} else {
			l--
		}

		for l >= 0 {
			for arr[j].cnt == 0 {
				j++
			}
			res[i][l] = arr[j].val
			res[i][r] = arr[j].val
			arr[j].cnt -= 2
			l--
			r++
		}
		ans[i] = string(res[i])
	}

	return ans
}
