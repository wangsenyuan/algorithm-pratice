package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	cards := strings.Split(s, " ")
	res := solve(cards)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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

func solve(card []string) bool {
	n := len(card)

	if n == 1 {
		return true
	}

	if n == 2 {
		return card[0][0] == card[1][0] || card[0][1] == card[1][1]
	}

	get := func(card string) int {
		var num, suit int
		if card[1] == 'S' {
			suit = 0
		} else if card[1] == 'D' {
			suit = 1
		} else if card[1] == 'H' {
			suit = 2
		} else {
			suit = 3
		}

		x := card[0]
		if x >= '2' && x <= '9' {
			num = int(x - '0')
		} else if x == 'A' {
			num = 1
		} else if x == 'T' {
			num = 10
		} else if x == 'J' {
			num = 11
		} else if x == 'Q' {
			num = 12
		} else {
			num = 13
		}
		num--
		return num*4 + suit
	}
	N := 4 * 13

	getState := func(i int, a, b, c int) int {
		// n-2, n - 1, n
		state := (c*N+b)*N + a
		return state*n + i
	}

	calcState := func(state int) (i int, a int, b int, c int) {
		i = state % n
		state /= n
		a = state % N
		state /= N
		b = state % N
		state /= N
		c = state
		return
	}

	vis := make([]bool, n*N*N*N)

	que := make([]int, n*N*N*N)

	begin := getState(n-1, get(card[n-3]), get(card[n-2]), get(card[n-1]))

	var head, tail int
	que[head] = begin
	head++

	vis[begin] = true

	for tail < head {
		cur := que[tail]
		i, a, b, c := calcState(cur)
		tail++
		if i == 2 {
			if bruteForce(a, b, c) {
				return true
			}
			continue
		}
		// i > 2
		d := get(card[i-3])
		if c/4 == b/4 || c%4 == b%4 {
			next := getState(i-1, d, a, c)
			if !vis[next] {
				vis[next] = true
				que[head] = next
				head++
			}
		}
		if c/4 == d/4 || c%4 == d%4 {
			next := getState(i-1, c, a, b)
			if !vis[next] {
				vis[next] = true
				que[head] = next
				head++
			}
		}
	}

	return false
}

func bruteForce(a int, b int, c int) bool {
	if c/4 != b/4 && c%4 != b%4 {
		// 最后一个不能移动到倒数第二个
		return false
	}
	if c/4 != a/4 && c%4 != a%4 {
		return false
	}
	return true
}
