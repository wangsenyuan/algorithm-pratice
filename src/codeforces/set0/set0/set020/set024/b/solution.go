package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0])
	fmt.Println(res[1])
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

func process(reader *bufio.Reader) []string {
	tc := readNum(reader)
	races := make([][]string, tc)
	for i := range tc {
		n := readNum(reader)
		races[i] = make([]string, n)
		for j := range n {
			races[i][j] = readString(reader)
		}
	}
	return solve(races)
}

var scores = []int{25, 18, 15, 12, 10, 8, 6, 4, 2, 1}

func solve(races [][]string) []string {
	points := make(map[string]int)
	wins := make(map[string][]int)
	for _, cur := range races {
		for i := 0; i < len(cur); i++ {
			name := cur[i]
			if i < 10 {
				points[name] += scores[i]
			}
			if len(wins[name]) == 0 {
				wins[name] = make([]int, 50)
			}
			wins[name][i]++
		}
	}

	find := func(players []string, x int) string {
		for i := x; len(players) > 1; i++ {
			var cnt int
			for _, player := range players {
				cnt = max(cnt, wins[player][i])
			}
			var next []string
			for _, player := range players {
				if cnt == wins[player][i] {
					next = append(next, player)
				}
			}
			players = next
		}
		return players[0]
	}

	var best int
	for _, v := range points {
		best = max(best, v)
	}
	var players []string
	for k, v := range points {
		if v == best {
			players = append(players, k)
		}
	}

	first := find(players, 0)

	var most int

	for _, vs := range wins {
		most = max(most, vs[0])
	}

	var players2 []string

	for k, vs := range wins {
		if most == vs[0] {
			players2 = append(players2, k)
		}
	}

	if len(players2) > 1 {
		var tmp []string
		best = 0
		for _, s := range players2 {
			best = max(best, points[s])
		}
		for _, s := range players2 {
			if points[s] == best {
				tmp = append(tmp, s)
			}
		}
		players2 = tmp
	}
	second := find(players2, 1)

	return []string{first, second}
}
