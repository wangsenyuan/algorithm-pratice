package main

import (
	"bufio"
	"bytes"
	"container/list"
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
		sheets := make([]string, n)
		for i := 0; i < n; i++ {
			sheets[i] = readString(reader)
		}
		res := solve(n, sheets)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(res[i])
			buf.WriteByte('\n')
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, sheets []string) []string {
	// 已经有 win 的，不需要参与
	// 其他的人，有以下几种情况
	// www, wwi, wwn, wii, wnn, iii, nnn, inn, iin
	// wwn iin
	graph := make([][]*list.List, 3)
	for i := 0; i < 3; i++ {
		graph[i] = make([]*list.List, 3)
		for j := 0; j < 3; j++ {
			graph[i][j] = list.New()
		}
	}
	ii := map[byte]int{'w': 0, 'i': 1, 'n': 2}
	cnt := make([]int, 3)
	for i, cur := range sheets {
		for j := 0; j < 3; j++ {
			cnt[j] = 0
		}
		for j := 0; j < 3; j++ {
			cnt[ii[cur[j]]]++
		}
		if cnt[0] == 1 && cnt[1] == 1 {
			continue
		}
		for x := 0; x < 3; x++ {
			if cnt[x] > 1 {
				// cnt[x] = 2, or cnt[x] = 3
				for y := 0; y < 3; y++ {
					if cnt[y] == 0 {
						graph[x][y].PushBack(i)
					}
				}
			}
		}
	}
	var ans []string
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y {
				continue
			}
			for graph[x][y].Len() > 0 && graph[y][x].Len() > 0 {
				i := graph[x][y].Back()
				j := graph[y][x].Back()
				ans = append(ans, exchange(i.Value.(int), x, j.Value.(int), y))
				graph[x][y].Remove(i)
				graph[y][x].Remove(j)
			}
		}
	}
	for graph[0][1].Len() > 0 && graph[1][2].Len() > 0 && graph[2][0].Len() > 0 {
		a := graph[0][1].Back()
		b := graph[1][2].Back()
		c := graph[2][0].Back()

		ans = append(ans, exchange(a.Value.(int), 0, b.Value.(int), 1))
		ans = append(ans, exchange(b.Value.(int), 0, c.Value.(int), 2))
		graph[0][1].Remove(a)
		graph[1][2].Remove(b)
		graph[2][0].Remove(c)
	}

	for graph[1][0].Len() > 0 && graph[2][1].Len() > 0 && graph[0][2].Len() > 0 {
		a := graph[1][0].Back()
		b := graph[2][1].Back()
		c := graph[0][2].Back()

		ans = append(ans, exchange(a.Value.(int), 1, c.Value.(int), 0))
		ans = append(ans, exchange(b.Value.(int), 2, c.Value.(int), 1))
		graph[1][0].Remove(a)
		graph[2][1].Remove(b)
		graph[0][2].Remove(c)
	}
	return ans
}

const win = "win"

func exchange(i int, x int, j int, y int) string {
	return fmt.Sprintf("%d %c %d %c", i+1, win[x], j+1, win[y])
}
