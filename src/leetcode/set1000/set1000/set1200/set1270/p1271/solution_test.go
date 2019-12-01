package p1271

import "testing"

func runSample(t *testing.T, num string, expect string) {
	res := toHexspeak(num)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", num, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "257", "IOI")
}
