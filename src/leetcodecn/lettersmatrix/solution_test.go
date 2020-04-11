package lettersmatrix

import "testing"

func runSample(t *testing.T, words []string, area int) {
	res := maxRectangle(words)

	if (area == 0) != (len(res) == 0) {
		t.Fatalf("Sample expect %d, but got %v", area, res)
	}

	if area == 0 {
		return
	}

	h, w := len(res), len(res[0])

	if w*h != area {
		t.Fatalf("Sample expect %d, but got %v", area, res)
	}

	mem := make(map[string]bool)

	for i := 0; i < len(words); i++ {
		mem[words[i]] = true
	}

	for j := 0; j < w; j++ {
		buf := make([]byte, h)
		for i := 0; i < h; i++ {
			buf[i] = res[i][j]
		}
		w := string(buf)
		if !mem[w] {
			t.Fatalf("Sample expect %d, but got %v", area, res)
		}
	}
}

func TestSample1(t *testing.T) {
	words := []string{"this", "real", "hard", "trh", "hea", "iar", "sld"}
	expect := 12
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"aa"}
	expect := 4
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"hv", "pi", "iu", "w", "yk", "lu", "dl", "e", "r", "pl"}
	expect := 4
	runSample(t, words, expect)
}
