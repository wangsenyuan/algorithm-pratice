package p752

// import "fmt"

func openLock(deadends []string, target string) int {
	N := 10000

	dead := make([]bool, N)
	for i := 0; i < len(deadends); i++ {
		x := toNum(deadends[i])
		dead[x] = true
	}
	if dead[0] {
		return -1
	}
	t := toNum(target)

	que := make([]int, N)
	head, tail := 0, 0
	que[tail] = 0
	tail++

	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = -1
	}
	dist[0] = 0
	for head < tail {
		tmp := tail
		for head < tmp {
			v := que[head]
			head++
			if v == t {
				return dist[v]
			}

			// fmt.Printf("[debug] visit %d\n", v)
			for i := 0; i < 4; i++ {
				x := rotate(v, i, 1)
				// fmt.Printf("[debug] rotate clock got %d from %d\n", x, v)
				if !dead[x] && dist[x] == -1 {
					que[tail] = x
					dist[x] = dist[v] + 1
					tail++
				}
				y := rotate(v, i, -1)
				// fmt.Printf("[debug] rotate anti-clock got %d from %d\n", y, v)
				if !dead[y] && dist[y] == -1 {
					que[tail] = y
					dist[y] = dist[v] + 1
					tail++
				}
			}
		}
	}

	return -1
}

func rotate(x int, pos int, dir int) int {
	a, b, c, d := x/1000, x/100%10, x/10%10, x%10
	if pos == 0 {
		a = (a + dir + 10) % 10
	} else if pos == 1 {
		b = (b + dir + 10) % 10
	} else if pos == 2 {
		c = (c + dir + 10) % 10
	} else {
		d = (d + dir + 10) % 10
	}
	return a*1000 + b*100 + c*10 + d
}

func toNum(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
