package p1276

import "testing"

func runSample(t *testing.T, tom int, chees int, can bool) {
	res := numOfBurgers(tom, chees)
	if len(res) == 0 && can {
		t.Errorf("Sample %d %d, got wrong answer %v", tom, chees, res)
	} else if can {
		x := res[0]
		y := res[1]
		if 4*x+2*y != tom || x+y != chees {
			t.Errorf("Sample %d %d, got wrong answer %v", tom, chees, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 16, 7, true)
}
