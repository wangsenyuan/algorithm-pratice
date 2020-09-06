package p1576

import "testing"

func runSample(t *testing.T, s string) {
	res := modifyString(s)

	for i := 1; i < len(res); i++ {
		if res[i] == res[i-1] || res[i] == '?' {
			t.Errorf("Sample result %s, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "?zs")
}

func TestSample2(t *testing.T) {
	runSample(t, "ubv?w")
}

func TestSample3(t *testing.T) {
	runSample(t, "j?qg??b")
}

func TestSample4(t *testing.T) {
	runSample(t, "??yw?ipkj?")
}

func TestSample5(t *testing.T) {
	runSample(t, "h???")
}

func TestSample6(t *testing.T) {
	runSample(t, "wz???a??n")
}
