package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %s", expect, res)
	}
	for i, j := 0, 0; i < len(res); i++ {
		for j < len(s) && s[j] != res[i] {
			j++
		}
		if j == len(s) {
			t.Fatalf("Sample result %s, not subseq of %s", res, s)
		}
		j++
	}
}

func TestSample1(t *testing.T) {
	s := "bbbabcbbb"
	expect := len("bbbcbbb")
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "rquwmzexectvnbanemsmdufrg"
	expect := len("rumenanemur")
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abaafafadfafadfafa"
	expect := bruteForce(s)
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := "abc"
	expect := bruteForce(s)
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "vjocwczmhxdzgfpcpohbtugxpxpbloifdfahsqwiuccqzicsghxyisvjsgzlyyttmpyhoccsfmtchgcvxejwvkkrmmuqxqczzmsvibjhhtxzokpwlggccyppfwwcjrijhvozgjljtfzrisuhfxyekacflqqbxlqllbgpvbsojkuvsujybbtxkifmxgnoqvljpsgetqpkoqkluktpfjxsuhjlzwsiljzchwmxgtpituseftblvbqfwvfgfggqhvgvoyzbgouubhkquwbyopsgoqjskbgbkxtqpjcfhywccsdzpziwlfjfysutcxxkumcjgtzrwubuobnfrmumqgzgvoblhztgeqtojhbmlqmuduttmbotwcpjhqoyxogzyxnwngydbpswbrchhuvtexkumfphsmgqfmuyuijfskgollzwukgwrwbbecpyhcyybmzsulsjmlskwfyukmqustmohgzomgoyobqbgbwjpcojlxcpwgpeqfgoxozbppusklf"
	expect := bruteForce(s)
	runSample(t, s, expect)
}
