package patternmatching

import "testing"

func runSample(t *testing.T, pattern, value string, expect bool) {
	res := patternMatching(pattern, value)

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", pattern, value, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "bbba", "xxxxxxy", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "abba", "dogcatcatdog", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "abba", "dogcatcatfish", false)
}

func TestSample4(t *testing.T) {
	runSample(t, "bbbbbbbbbbbbbbabbbbb", "ppppppppppppppjsftcleifftfthiehjiheyqkhjfkyfckbtwbelfcgihlrfkrwireflijkjyppppg", true)
}

func TestSample5(t *testing.T) {
	runSample(t, "abb", "dryqxzysggjljxdxag", true)
}

func TestSample6(t *testing.T) {
	runSample(t, "ab", "", false)
}
