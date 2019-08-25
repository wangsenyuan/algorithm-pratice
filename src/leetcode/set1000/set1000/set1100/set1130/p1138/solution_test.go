package p1138

import (
	"bytes"
	"testing"
)

var board = []string{
	"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z",
}

func runSample(t *testing.T, target string, expect string) {
	res := alphabetBoardPath(target)

	if len(res) != len(expect) {
		t.Errorf("Sample %s, expect %s, but got %s case one", target, expect, res)
		return
	}
	var x, y int
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		if res[i] == 'D' {
			x++
		} else if res[i] == 'U' {
			x--
		} else if res[i] == 'L' {
			y--
		} else if res[i] == 'R' {
			y++
		} else if x >= 0 && x <= 5 && y >= 0 && y < len(board[x]) {
			buf.WriteByte(board[x][y])
		} else {
			t.Errorf("Sample %s, expect %s, but got %s", target, expect, res)
			return
		}
	}

	if target != buf.String() {
		t.Errorf("Sample %s, expect %s, but got %s", target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "leet", "DDR!UURRR!!DDD!")
}

func TestSample2(t *testing.T) {
	runSample(t, "code", "RR!DDRR!UUL!R!")
}

func TestSample3(t *testing.T) {
	runSample(t, "zb", "DDDDD!UUUUUR!")
}
