package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	f, err := os.Open("./A-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var n int
		fmt.Fscanf(f, "%d", &n)
		actions := make([]string, n)
		count := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscanf(f, "%s", &actions[i])
			fmt.Fscanf(f, "%d", &count[i])
		}
		res := solve(n, actions, count)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

func solve(n int, actions []string, count []int) int {
	var x, y, dir int
	var A int
	t := make([]int, 6010)
	b := make([]int, 6010)

	for i := 0; i < 6010; i++ {
		t[i] = math.MinInt32
		b[i] = math.MaxInt32
	}

	x0 := 3000
	var Ax, Bx, Cx, Dx int
	Bx = math.MinInt32
	Dx = math.MaxInt32
	// dir = 0, 1, 2, 3 (N, W, S, E)
	for i := 0; i < n; i++ {
		action := actions[i]
		for j := 0; j < count[i]; j++ {
			for k := 0; k < len(action); k++ {
				if action[k] == 'F' {
					forward(&x, &y, dir)
					x1 := x0 + x

					if dir == 1 {
						// from right to left
						A -= y
						if y > t[x1] {
							t[x1] = y
						}
						if y < b[x1] {
							b[x1] = y
						}
					} else if dir == 3 {
						// from left to right
						A += y
						if y > t[x1-1] {
							t[x1-1] = y
						}
						if y < b[x1-1] {
							b[x1-1] = y
						}
					}

					if x1 < Dx {
						Dx = x1
					}
					if x1 > Bx {
						Bx = x1
					}
					if y > t[Ax] {
						Ax = x1
					}
					if y < b[Cx] {
						Cx = x1
					}
				} else if action[k] == 'L' {
					turnLeft(&dir)
				} else {
					turnRight(&dir)
				}
			}
		}
	}
	if A < 0 {
		A = -A
	}

	H := make([]int, 6010)
	L := make([]int, 6010)

	H[Bx-1] = t[Bx-1]

	// East
	for x := Bx - 2; x >= Ax; x-- {
		H[x] = max(H[x+1], t[x])

	}
	L[Bx-1] = b[Bx-1]

	for x := Bx - 2; x >= Cx; x-- {
		L[x] = min(L[x+1], b[x])
	}
	// West
	H[Dx] = t[Dx]

	for x := Dx + 1; x < Ax; x++ {
		H[x] = max(H[x-1], t[x])
	}

	L[Dx] = b[Dx]

	for x := Dx + 1; x < Cx; x++ {
		L[x] = min(L[x-1], b[x])
	}

	var B int
	for x := Dx; x < Bx; x++ {
		B += H[x] - L[x]
	}

	return B - A
}

func forward(x, y *int, dir int) {
	if dir == 0 {
		(*y)++
	} else if dir == 1 {
		(*x)--
	} else if dir == 2 {
		(*y)--
	} else {
		(*x)++
	}
}

func turnLeft(dir *int) {
	*dir = (*dir + 1) % 4
}

func turnRight(dir *int) {
	*dir = (*dir - 1 + 4) % 4
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
