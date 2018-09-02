package p899

import "testing"

func runSample(t *testing.T, s string, K int, expect string) {
	res := orderlyQueue(s, K)

	if res != expect {
		t.Errorf("Sample %s %d, expect %s, but got %s", s, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "cba", 1, "acb")
}

func TestSample2(t *testing.T) {
	runSample(t, "baaca", 3, "aaabc")
}

func TestSample3(t *testing.T) {
	runSample(t, "nhtq", 1, "htqn")
}

func TestSample4(t *testing.T) {
	runSample(t, "xmvzi", 2, "imvxz")
}
