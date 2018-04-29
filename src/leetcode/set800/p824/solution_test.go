package p824

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := toGoatLatin(S)
	if res != expect {
		t.Errorf("sample %s, expect %s, but got %s", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa")
}

func TestSample2(t *testing.T) {
	runSample(t, "The quick brown fox jumped over the lazy dog",
		"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa")
}
