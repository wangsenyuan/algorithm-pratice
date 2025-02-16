package p3458

import "testing"

func runSample(t *testing.T, s string, k int, expect bool) {
	res := maxSubstringLength(s, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but go %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcdbaefab"
	k := 2
	expect := true
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "cdefdc"
	k := 3
	expect := false
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "abeabe"
	k := 0
	expect := true
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "ddjlopbgngpoenkqktvuuvpygctyquoeqglszijjiifljfiswiladdfwzislzdd"
	k := 6
	expect := false
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "yefby"
	k := 3
	expect := true
	runSample(t, s, k, expect)
}
func TestSample6(t *testing.T) {
	s := "gaixgqpgdrhhxuurgrriwovkbjjbffnlnjcnyzbmblymcmvnftlpaiqepgxgiperurhgduaqpqxqgaxdexur"
	k := 6
	expect := false
	runSample(t, s, k, expect)
}

func TestSample7(t *testing.T) {
	s := "nbuirvanjiccnsyyyoirleqsrwrvxepaglcidqplyryujytzqoncxjgwdmatytgwhzyhlsodrbzrpbbitovtdasazjtoyyfhowqqrzuvjveydceouscrfazzoblqhalhfybwheybkpcroijxvarrtqrqnmwslkpdducfeblvfecyjyulxgahxlzlyztssfzwvfujrriryslkvdwhmkcyebfhkadrahunvxivkwitilyzknwyujtylahgmlddymlbrbrniomepbmdieasuvdcqnzfwspxewbbpruxrznjxwnjjxvblxyrgv"
	k := 1
	expect := false
	runSample(t, s, k, expect)
}

func TestSample8(t *testing.T) {
	s := "gjfg"
	k := 1
	expect := true
	runSample(t, s, k, expect)
}
