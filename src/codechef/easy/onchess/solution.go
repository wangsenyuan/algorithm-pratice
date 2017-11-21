package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := toInt(scanner.Bytes())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n := toInt(scanner.Bytes())
		players := make([]Player, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			players[j] = toPlayer(scanner.Bytes())
		}
		res := solve(n, players)
		for i := 0; i < n; i++ {
			fmt.Println(res[i])
		}
	}
}

type Player struct {
	rating int
	min    int
	max    int
	time   int
	rated  bool
	color  int
}

func toPlayer(bs []byte) Player {
	pos := 0
	var rating, min, max, time, color int
	var rated bool
	pos = readIntToken(bs, pos, &rating)
	pos = readIntToken(bs, pos+1, &min)
	pos = readIntToken(bs, pos+1, &max)
	pos = readIntToken(bs, pos+1, &time)
	pos++
	if bs[pos] == 'r' {
		rated = true
	}
	for bs[pos] != ' ' {
		pos++
	}
	pos++
	if bs[pos] == 'w' {
		color = -1
	} else if bs[pos] == 'b' {
		color = 1
	}

	return Player{rating, min, max, time, rated, color}
}

func readIntToken(bs []byte, from int, res *int) int {
	var tmp int
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
		i++
	}
	*res = tmp
	return i
}

func toInt(bs []byte) int {
	var res int

	for i := 0; i < len(bs); i++ {
		res = res*10 + int(bs[i]-'0')
	}

	return res
}

func solve(n int, players []Player) []interface{} {
	res := make([]interface{}, n)
	pair := make([]bool, n)

	canPair := func(i, j int) bool {
		a, b := players[i], players[j]

		if a.rating < b.min || a.rating > b.max {
			return false
		}

		if b.rating < a.min || b.rating > a.max {
			return false
		}

		if a.time != b.time {
			return false
		}
		if a.rated != b.rated {
			return false
		}
		if a.color+b.color != 0 {
			return false
		}
		return true
	}

	for i := 0; i < n; i++ {
		j := 0
		for j < i {
			if !pair[j] && canPair(i, j) {
				break
			}
			j++
		}
		if j < i {
			res[i] = j + 1
			pair[i], pair[j] = true, true
		} else {
			res[i] = "wait"
		}
	}

	return res
}
