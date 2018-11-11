package p939

func minAreaRect(points [][]int) int {
	mem := make(map[PNT]bool)

	n := len(points)

	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		mem[PNT{x, y}] = true
	}

	var res int

	for i := 0; i < n; i++ {
		a := PNT{points[i][0], points[i][1]}
		for j := i + 1; j < n; j++ {
			b := PNT{points[j][0], points[j][1]}

			x := PNT{a.x, b.y}
			y := PNT{b.x, a.y}

			if mem[x] && mem[y] {
				tmp := area(a, b)
				if tmp > 0 && (res == 0 || tmp < res) {
					res = tmp
				}
			}
		}
	}
	return res
}

func area(a, b PNT) int {
	return abs(b.y-a.y) * abs(b.x-a.x)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type PNT struct {
	x, y int
}
