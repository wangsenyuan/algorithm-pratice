package p1694

import "testing"

func runSample(t *testing.T, str string, expect string) {
	res := reformatNumber(str)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", str, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1-23-45 6", "123-456")
}
