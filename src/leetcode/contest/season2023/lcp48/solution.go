package lcp48

import "fmt"

func gobang(pieces [][]int) string {
	// 黑走两步，白走一步
	// 如果黑走一步胜，则黑胜
	// 如果白走一步胜，且黑无法阻止，则白胜
	// 如果黑走两步胜，且白无法阻止，则黑胜
	board := make(map[int]map[int]int)

	set := func(a, b, c int) {
		if board[a] == nil {
			board[a] = make(map[int]int)
		}
		board[a][b] = c
	}

	for _, piece := range pieces {
		a, b, c := piece[0], piece[1], piece[2]
		set(a, b, c)
	}

	get := func(a, b int) int {
		if v, ok := board[a]; !ok {
			return -1
		} else if w, ok2 := v[b]; !ok2 {
			return -1
		} else {
			return w
		}
	}

	checkWin := func(a, b, c, dx, dy int) bool {
		// go 5 steps
		var hole int
		for i := 0; i < 5; i++ {
			x := get(a, b)
			if x < 0 {
				// a hole
				if hole == 1 {
					return false
				}
				hole++
			} else if x != c {
				// already blocked
				return false
			}
			a += dx
			b += dy
		}

		return hole == 1
	}

	findWinPosition := func(color int) [][]int {
		// 有没有一步内即可以获胜的
		var res [][]int
		for _, piece := range pieces {
			a, b, c := piece[0], piece[1], piece[2]
			if c == color {
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue
						}
						if checkWin(a, b, c, dx, dy) {
							res = append(res, []int{a, b, dx, dy})
						}
					}
				}
			}
		}

		return res
	}

	findBlockingPosition := func(ps [][]int) [][]int {
		mem := make(map[string]bool)
		var res [][]int
		for _, p := range ps {
			a, b, dx, dy := p[0], p[1], p[2], p[3]
			for i := 0; i < 5; i++ {
				x := get(a, b)
				if x < 0 {
					// a hole
					key := fmt.Sprintf("%d-%d", a, b)
					if !mem[key] {
						res = append(res, []int{a, b})
						mem[key] = true
					}
					break
				}
				a += dx
				b += dy
			}
		}
		return res
	}

	black_win_positions := findWinPosition(0)

	if len(black_win_positions) > 0 {
		// win at the very first step
		return "Black"
	}
	// else no win positions for first player at one step

	white_win_positions := findWinPosition(1)

	if len(white_win_positions) > 0 {
		// must stop
		holes := findBlockingPosition(white_win_positions)
		if len(holes) > 1 {
			// second player can win in one step
			return "White"
		}
		// first step by first player to stop second player win
		a, b := holes[0][0], holes[0][1]
		set(a, b, 0)
		pieces = append(pieces, []int{a, b, 0})
		// still might won by first player
		black_win_positions = findWinPosition(0)

		holes = findBlockingPosition(black_win_positions)
		// second player have to stop first player by put white button at the blocking position
		if len(holes) > 1 {
			// too many positions
			return "Black"
		}
		return "None"
	}
	// no one-step winning, first player have to find two steps winning
	// the hard part
	// 是不是找出所有在两步内能获胜的line
	// 还是找出符合某种模式的形状
	// 如果是 .**.*. 中间有一个空格，且两端都有空格
	// .***.., 且两端都是空格, 且有一端至少有两个空格
	// 还有就是一个空格，放入后，可以造成两个连4，且它们都有获胜的位置的

	vis := make(map[string]bool)

	check3 := func(a, b int) bool {
		if vis[fmt.Sprintf("%d-%d", a, b)] {
			return false
		}
		vis[fmt.Sprintf("%d-%d", a, b)] = true
		// place black button at this position, and check if it can win in one more step
		holes := make(map[string]int)
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				// a, b 只是中间的某个节点
				u, v := a, b

				for i := 0; i < 5; i++ {
					u -= dx
					v -= dy
					color := get(u, v)
					if color == 1 {
						break
					}
					var hole [][]int
					x, y := u, v
					j := 0
					for ; j < 5; x, y, j = x+dx, y+dy, j+1 {
						if x == a && y == b {
							continue
						}
						color = get(x, y)
						if color == 1 {
							break
						}
						if color < 0 {
							hole = append(hole, []int{x, y})
						}
					}
					if len(hole) == 1 && j == 5 {
						x, y = hole[0][0], hole[0][1]
						holes[fmt.Sprintf("%d-%d", x, y)]++
					}
				}
			}
		}
		// 如果只有一个空位获胜，会被second player阻止
		return len(holes) > 1
	}

	check2 := func(a, b, dx, dy int) bool {
		//a, b is black
		x, y := a, b
		for i := 0; i < 3; i++ {
			c := get(x, y)
			if c == 1 {
				break
			}
			if c < 0 {
				tmp := check3(x, y)
				if tmp {
					return true
				}
			}
			x += dx
			y += dy
		}
		return false
	}

	checkWin2 := func(a, b int) bool {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				if check2(a, b, dx, dy) {
					return true
				}
			}
		}
		return false
	}

	for _, piece := range pieces {
		a, b, c := piece[0], piece[1], piece[2]
		if c == 0 && checkWin2(a, b) {
			return "Black"
		}
	}
	return "None"
}
