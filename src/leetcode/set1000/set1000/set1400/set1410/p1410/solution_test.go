package p1410

import "testing"

func runSample(t *testing.T, text string, expect string) {
	res := entityParser(text)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", text, expect, res)
	}
}

func TestSample1(t *testing.T) {
	text := "&amp; is an HTML entity but &ambassador; is not."
	expect := "& is an HTML entity but &ambassador; is not."
	runSample(t, text, expect)
}

func TestSample2(t *testing.T) {
	text := "and I quote: &quot;...&quot;"
	expect := "and I quote: \"...\""
	runSample(t, text, expect)
}
