package p1366

import "testing"

func runSample(t *testing.T, votes []string, expect string) {
	res := rankTeams(votes)

	if res != expect {
		t.Errorf("Sample %v, expect %s, but got %s", votes, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []string{"ABC", "ACB", "ABC", "ACB", "ACB"}, "ACB")
}

func TestSample2(t *testing.T) {
	runSample(t, []string{"WXYZ", "XYZW"}, "XWYZ")
}

func TestSample3(t *testing.T) {
	runSample(t, []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}, "ZMNAGUEDSJYLBOPHRQICWFXTVK")
}

func TestSample4(t *testing.T) {
	runSample(t, []string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}, "ABC")
}
