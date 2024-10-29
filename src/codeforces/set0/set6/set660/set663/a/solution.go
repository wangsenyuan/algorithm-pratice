package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	ok, res := solve(s)

	if !ok {
		fmt.Println("Impossible")
	} else {
		fmt.Println("Possible")
		fmt.Println(res)
	}
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

func getLastN(s string) int {

	var n int
	ll := strings.LastIndex(s, " ") + 1
	readInt([]byte(s), ll, &n)
	return n
}
func solve(s string) (bool, string) {
	var pos []int
	var neg []int
	// 第一个数肯定是正数
	pos = append(pos, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			pos = append(pos, len(pos)+len(neg))
		} else if s[i] == '-' {
			neg = append(neg, len(pos)+len(neg))
		}
	}
	n := getLastN(s)

	mx := n*len(pos) - len(neg)
	mn := len(pos) - n*len(neg)
	if mx < n || mn > n {
		return false, ""
	}
	arr := make([]int, len(pos)+len(neg))
	// 所有的都放置位-1/1
	for _, i := range pos {
		arr[i] = 1
	}

	for _, i := range neg {
		arr[i] = -1
	}

	sum := len(pos) - len(neg)
	i, j := len(pos), len(neg)
	for sum != n {
		if sum > n {
			diff := sum - n
			if diff+1 <= n {
				// 可以加在neg上
				arr[neg[j-1]] -= diff
				sum -= diff
				break
			}
			arr[neg[j-1]] = -n
			sum -= n - 1
			j--
		} else {
			// sum < n
			diff := n - sum
			if diff+1 <= n {
				arr[pos[i-1]] += diff
				sum += diff
				break
			}
			arr[pos[i-1]] = n
			sum += n - 1
			i--
		}
	}

	return true, construct(arr, n)
}

func construct(arr []int, n int) string {
	var buf bytes.Buffer

	for i, x := range arr {
		if x > 0 {
			if i == 0 {
				buf.WriteString(fmt.Sprintf("%d", x))
			} else {
				buf.WriteString(fmt.Sprintf(" + %d", x))
			}
		} else {
			buf.WriteString(fmt.Sprintf(" - %d", -x))
		}
	}
	buf.WriteString(fmt.Sprintf(" = %d", n))
	return buf.String()
}
