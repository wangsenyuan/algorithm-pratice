package p773

import "strconv"

var dist map[string]int

func init() {
	que := make([]string, 6*5*4*3*2*1)
	head, tail := 0, 0
	que[tail] = "123450"
	tail++
	dist = make(map[string]int)
	dist["123450"] = 0

	for head < tail {
		tmp := tail
		for head < tmp {
			cur := que[head]
			head++
			ns := swap(cur)
			for _, next := range ns {
				if _, found := dist[next]; !found {
					dist[next] = dist[cur] + 1
					que[tail] = next
					tail++
				}
			}
		}
	}
}

func swap(cur string) []string {
	// 1 2 3
	// 4 5 0
	bs := []byte(cur)
	if bs[0] == '0' {
		bs[0], bs[1] = bs[1], bs[0]
		s1 := string(bs)
		//change back
		bs[0], bs[1] = bs[1], bs[0]
		bs[0], bs[3] = bs[3], bs[0]
		s2 := string(bs)
		return []string{s1, s2}
	}
	if bs[1] == '0' {
		bs[1], bs[0] = bs[0], bs[1]
		s1 := string(bs)
		bs[1], bs[0] = bs[0], bs[1]
		bs[1], bs[4] = bs[4], bs[1]
		s2 := string(bs)
		bs[1], bs[4] = bs[4], bs[1]
		bs[1], bs[2] = bs[2], bs[1]
		s3 := string(bs)
		return []string{s1, s2, s3}
	}
	if bs[2] == '0' {
		bs[2], bs[1] = bs[1], bs[2]
		s1 := string(bs)
		bs[2], bs[1] = bs[1], bs[2]
		bs[2], bs[5] = bs[5], bs[2]
		s2 := string(bs)
		return []string{s1, s2}
	}

	if bs[3] == '0' {
		bs[3], bs[4] = bs[4], bs[3]
		s1 := string(bs)
		//change back
		bs[3], bs[4] = bs[4], bs[3]
		bs[3], bs[0] = bs[0], bs[3]
		s2 := string(bs)
		return []string{s1, s2}
	}

	if bs[4] == '0' {
		bs[4], bs[3] = bs[3], bs[4]
		s1 := string(bs)
		bs[4], bs[3] = bs[3], bs[4]
		bs[4], bs[1] = bs[1], bs[4]
		s2 := string(bs)
		bs[4], bs[1] = bs[1], bs[4]
		bs[4], bs[5] = bs[5], bs[4]
		s3 := string(bs)
		return []string{s1, s2, s3}
	}
	bs[5], bs[4] = bs[4], bs[5]
	s1 := string(bs)
	bs[5], bs[4] = bs[4], bs[5]
	bs[5], bs[2] = bs[2], bs[5]
	s2 := string(bs)
	return []string{s1, s2}
}

func slidingPuzzle(board [][]int) int {
	var str string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			str += strconv.Itoa(board[i][j])
		}
	}

	if d, found := dist[str]; found {
		return d
	}
	return -1
}
