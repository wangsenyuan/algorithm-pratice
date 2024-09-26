package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	logs := make([]string, n)

	for i := 0; i < n; i++ {
		logs[i] = readString(reader)
	}

	res := solve(logs)

	fmt.Println(res)
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

func solve(logs []string) int {
	var res int

	n := len(logs)

	for i := 0; i < n; {
		cur := parse(logs[i])
		i++
		res++
		cnt := 1
		for i < n {
			tmp := parse(logs[i])
			if cur > tmp {
				break
			}
			if cur == tmp {
				cnt++
				if cnt > 10 {
					break
				}
			} else {
				// cur < tmp
				cnt = 1
			}
			cur = tmp
			i++
		}
	}
	return res
}

const pattern = "[05:00 a.m.]"

func parse(log string) int {
	i := len(pattern)
	s := log[1 : i-1]
	hour := int(s[0]-'0')*10 + int(s[1]-'0')
	hour %= 12
	min := int(s[3]-'0')*10 + int(s[4]-'0')
	time := hour*60 + min
	if s[6:] == "p.m." {
		time += 12 * 60
	}
	return time
}
