package p913

const DRAW = 0
const MOUSE = 1
const CAT = 2

func catMouseGame(graph [][]int) int {
	n := len(graph)
	color := initTripDimSlice(50, 50, 3)
	degree := initTripDimSlice(50, 50, 3)

	for m := 0; m < n; m++ {
		for c := 0; c < n; c++ {
			// mouse can move out with len(graph[m]) ways
			degree[m][c][1] = len(graph[m])
			// cat move ways
			degree[m][c][2] = len(graph[c])
			for i := 0; i < len(graph[c]); i++ {
				if graph[c][i] == 0 {
					//cat can't move to hole
					degree[m][c][2]--
					break
				}
			}
		}
	}

	que := make([]State, 0, 50*50*3)
	for i := 0; i < n; i++ {
		// 1 for cat, 2 for mouse
		for t := 1; t <= 2; t++ {
			// when mouse is at hole, mouse wins
			color[0][i][t] = MOUSE
			que = append(que, State{0, i, t, MOUSE})
			if i > 0 {
				// when cat reach mouse, cat wins
				color[i][i][t] = CAT
				que = append(que, State{i, i, t, CAT})
			}
		}
	}

	findParent := func(m, c, t int) [][]int {
		res := make([][]int, 0, 10)
		if t == 2 {
			for i := 0; i < len(graph[m]); i++ {
				m2 := graph[m][i]
				// m2, c, and cat's turn
				res = append(res, []int{m2, c, 1})
			}
		} else {
			for i := 0; i < len(graph[c]); i++ {
				c2 := graph[c][i]
				if c2 > 0 {
					res = append(res, []int{m, c2, 2})
				}
			}
		}
		return res
	}

	var front int
	for front < len(que) {
		cur := que[front]
		front++
		i := cur.mouse
		j := cur.cat
		t := cur.turn
		c := cur.color
		parents := findParent(i, j, t)
		for _, parent := range parents {
			i2 := parent[0]
			j2 := parent[1]
			t2 := parent[2]
			if color[i2][j2][t2] == DRAW {
				if t2 == c {
					color[i2][j2][t2] = c
					que = append(que, State{i2, j2, t2, c})
				} else {
					degree[i2][j2][t2]--
					if degree[i2][j2][t2] == 0 {
						color[i2][j2][t2] = 3 - t2
						que = append(que, State{i2, j2, t2, 3 - t2})
					}
				}
			}
		}
	}

	return color[1][2][1]
}

type State struct {
	mouse int
	cat   int
	turn  int
	color int
}

func initTripDimSlice(a, b, c int) [][][]int {
	res := make([][][]int, a)
	for i := 0; i < a; i++ {
		res[i] = make([][]int, b)
		for j := 0; j < b; j++ {
			res[i][j] = make([]int, c)
		}
	}
	return res
}
