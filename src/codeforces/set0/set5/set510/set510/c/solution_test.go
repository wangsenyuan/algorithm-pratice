package main

import "testing"

func runSample(t *testing.T, words []string, expect bool) {
	res := solve(words)

	if (res != "Impossible") != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}

	order := make([]int, 26)

	for i, x := range []byte(res) {
		order[int(x-'a')] = i
	}

	for i := 0; i+1 < len(words); i++ {
		a := words[i]
		b := words[i+1]
		var k int
		for k < len(a) && k < len(b) && a[k] == b[k] {
			k++
		}
		if k == len(a) {
			continue
		}
		x := int(a[k] - 'a')
		y := int(b[k] - 'a')
		if order[x] > order[y] {
			t.Fatalf("Sample result %s, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"rivest",
		"shamir",
		"adleman",
	}
	expect := true
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"tourist",
		"petr",
		"wjmzbmr",
		"yeputons",
		"vepifanov",
		"scottwu",
		"oooooooooooooooo",
		"subscriber",
		"rowdark",
		"tankengineer",
	}
	expect := false
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"car",
		"care",
		"careful",
		"carefully",
		"becarefuldontforgetsomething",
		"otherwiseyouwillbehacked",
		"goodluck",
	}
	expect := true
	runSample(t, words, expect)
}

func TestSample4(t *testing.T) {
	words := []string{
		"adjcquqdqcsvixgwglwrrmkhdsbebbjvcgz",
		"aoalcgxovldqfzaorahdigyojknviaztpcmxlvovafhjphvshyfiqqtqbxjjmq",
		"bijtrgtxgzfcpuchdsmkvqrhgdiidnnunwsxscqqn",
		"bmxqnwdcyberszmofrwoc",
		"bomjxrqsfskevfxphcqoqkbbomcyurwlrnhrhctntzlylvwulbdkcdppgykichjtpukfnlxfcev",
		"bynjwjcrqayjpcijndyzmctzbaoqeugwrmhpkzvyygkovpckbzetshne",
		"cfhp",
		"cgdlozuqzwhfssrcxarcjflftmwngzonwfsvtukxvbsoubbbpbfwjfngelekpmoadizdsortkabmsw",
		"cgffiwrszzdbzjyhvgpnfsapufjqfespxbfljgtdgsmfeecqfwfvkweiacditmsnfaldcqlrycllmccmodlnlbkwv",
	}
	expect := true
	runSample(t, words, expect)
}
