package p1516

const N = 100010

var win []bool

func init() {
	win = make([]bool, N)
	for i := 1; i < N; i++ {
		for x := 1; x*x <= i; x++ {
			y := x * x
			if !win[i-y] {
				win[i] = true
				break
			}
		}
	}
}

func winnerSquareGame(n int) bool {
	return win[n]
}
