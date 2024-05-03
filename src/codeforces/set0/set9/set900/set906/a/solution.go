package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	records := make([]string, n)

	for i := 0; i < n; i++ {
		records[i] = readString(reader)
	}

	res := solve(records)

	fmt.Println(res)
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

func solve(records []string) int {
	var res int
	active := make(map[int]bool)
	for i := 0; i < 26; i++ {
		active[i] = true
	}
	for _, record := range records {
		if len(active) == 1 {
			if record[0] != '.' {
				res++
			}
			continue
		}
		if record[0] == '!' {
			// keep the same ones
			var tmp int
			for i := 2; i < len(record); i++ {
				x := int(record[i] - 'a')
				if active[x] {
					tmp |= 1 << x
				}
			}
			// refresh the active
			active = make(map[int]bool)
			for i := 0; i < 26; i++ {
				if (tmp>>i)&1 == 1 {
					active[i] = true
				}
			}
		} else {
			// 这个里面肯定没有答案
			for i := 2; i < len(record); i++ {
				x := int(record[i] - 'a')
				delete(active, x)
			}
		}
	}

	return max(res-1, 0)
}
