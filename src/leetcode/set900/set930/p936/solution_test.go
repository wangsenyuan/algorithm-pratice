package p936

import "testing"

func runSample(t *testing.T, stamp, target string, hasAnswer bool) {
	res := movesToStamp(stamp, target)

	if (len(res) != 0) != hasAnswer {
		t.Errorf("Sample %s %s, expect %t, but got %v", stamp, target, hasAnswer, res)
		return
	}
	if hasAnswer {
		buf := make([]byte, len(target))
		for _, i := range res {
			for j := 0; j < len(stamp); j++ {
				buf[i+j] = stamp[j]
			}
		}
		str := string(buf)
		if str != target {
			t.Errorf("Sample %s %s, got %v, give wrong result %s", stamp, target, res, str)
		}
	}
}

func TestSample1(t *testing.T) {
	stamp := "abc"
	target := "ababc"
	runSample(t, stamp, target, true)
}

func TestSample2(t *testing.T) {
	stamp := "abca"
	target := "ababc"
	runSample(t, stamp, target, true)
}
