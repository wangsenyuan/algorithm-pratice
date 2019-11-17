package p1256

import "testing"

func runSample(t *testing.T, num int, expect string) {
	res := encode(num)

	if res != expect {
		t.Errorf("Sample %d, expect %s, but got %s", num, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, "")
	runSample(t, 1, "0")
	runSample(t, 2, "1")
	runSample(t, 3, "00")
	runSample(t, 4, "01")
	runSample(t, 5, "10")
	runSample(t, 6, "11")
	runSample(t, 7, "000")
}

func TestSample2(t *testing.T) {
	runSample(t, 23, "1000")
}

func TestSample3(t *testing.T) {
	runSample(t, 107, "101100")
}
