package p1244

const MAX_N = 100*1000 + 10

type Leaderboard struct {
	size   int
	arr    []int
	scores map[int]int
	sum    []int
	total  int
}

func Constructor() Leaderboard {
	arr := make([]int, 4*MAX_N)
	scores := make(map[int]int)
	sum := make([]int, 4*MAX_N)

	return Leaderboard{MAX_N, arr, scores, sum, 0}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	n := this.size

	// v := this.scores[playerId] + score

	var loop func(i int, start int, end int, score int, v int)

	loop = func(i int, start int, end int, score int, v int) {
		if start == end {
			this.arr[i] += v
			this.sum[i] += v * score
			return
		}
		mid := (start + end) / 2
		if mid+1 <= score {
			loop(2*i+2, mid+1, end, score, v)
		} else {
			loop(2*i+1, start, mid, score, v)
		}
		this.arr[i] = this.arr[2*i+1] + this.arr[2*i+2]
		this.sum[i] = this.sum[2*i+1] + this.sum[2*i+2]
	}

	// this.total -= this.scores[playerId]

	if this.scores[playerId] > 0 {
		loop(0, 0, n-1, this.scores[playerId], -1)
	}
	this.scores[playerId] += score

	if this.scores[playerId] > 0 {
		loop(0, 0, n-1, this.scores[playerId], 1)
	}

	this.total += score
}

func (this *Leaderboard) Top(K int) int {
	m := len(this.scores)
	// total m players
	x := m - K

	var loop func(i int, start int, end int, cnt int) int

	loop = func(i int, start int, end int, cnt int) int {
		if cnt <= 0 {
			return 0
		}

		if start == end {
			return cnt * start
		}

		if this.arr[i] <= cnt {
			return this.sum[i]
		}

		mid := (start + end) / 2

		return loop(2*i+1, start, mid, min(cnt, this.arr[2*i+1])) + loop(2*i+2, mid+1, end, cnt-this.arr[2*i+1])
	}

	ls := loop(0, 0, this.size-1, x)

	return this.total - ls
}

func (this *Leaderboard) Reset(playerId int) {
	this.AddScore(playerId, -this.scores[playerId])
	delete(this.scores, playerId)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
