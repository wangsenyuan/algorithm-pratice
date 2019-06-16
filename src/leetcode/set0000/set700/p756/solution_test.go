package p756

import "testing"

func runSample(t *testing.T, bottom string, allowed []string, expect bool) {
	res := pyramidTransition(bottom, allowed)
	if res != expect {
		t.Errorf("sample: %s %v; should give answer %v, but got %v", bottom, allowed, expect, res)
	}
}

func TestSample1(t *testing.T) {
	bottom := "XYZ"
	allowed := []string{
		"XYD", "YZE", "DEA", "FFF",
	}
	expect := true
	runSample(t, bottom, allowed, expect)
}

func TestSample2(t *testing.T) {
	bottom := "XXYX"
	allowed := []string{
		"XXX", "XXY", "XYX", "XYY", "YXZ",
	}
	expect := false
	runSample(t, bottom, allowed, expect)
}

func TestSample3(t *testing.T) {
	bottom := "ABCD"
	allowed := []string{
		"BCE", "BCF", "ABA", "CDA", "AEG", "FAG", "GGG",
	}
	expect := false
	runSample(t, bottom, allowed, expect)
}

func TestSample4(t *testing.T) {
	bottom := "DACCDGDDCFCGCFAGFADF"
	allowed := []string{"BGF", "AGE", "AGC", "AGA", "CCE", "CCD", "EGA", "CCF", "CCA", "DCD", "DCA", "FGE", "FGA", "FGB", "BFB", "BFG", "BFD", "ECC", "DBA", "FAF", "DBF", "FDA", "FDC", "FDE", "BEA", "BEE", "AEA", "AEC", "AED", "EEE", "DEA", "DEC", "EEA", "CCG", "EEC", "DEG", "CEE", "CED", "CEC", "CEA", "GEC", "GEA", "GEF", "GEE", "BDE", "BDD", "GCE", "AFC", "DDC", "DDB", "EFB", "EFE", "DDE", "DDD", "CBC", "CBE", "ACB", "ACE", "BCD", "BCE", "BCA", "BCB", "BCC", "DGB", "ECF", "DGF", "ECB", "ECA", "CGD", "CGF", "FCE", "FEF", "BBF", "BBD", "ADG", "ADD", "ADA", "DFD", "DFC", "CDE", "CDF", "CDG", "EDF", "EDG", "EDD", "FBA", "FBG", "FBF", "GDF", "AAE", "AAD", "AAC", "BAG", "BAB", "BAC", "BAA", "CAF", "CAD", "DAD", "DAB", "DED", "EAD", "EAG", "EAF", "FAC", "GAD", "GAC", "GAB", "ABA", "ABF", "EBD", "EBF", "EBA", "EBB", "EBC", "CFF", "CFE", "CFC", "GFC", "GFA", "GFG", "GFD"}
	expect := true
	runSample(t, bottom, allowed, expect)
}
