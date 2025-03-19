package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if len(res) != len(expect) {
		t.Errorf("结果长度不匹配,期望 %d, 得到 %d", len(expect), len(res))
		return
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expect[i] {
			t.Errorf("第%d个结果不匹配,期望 %d, 得到 %d", i, expect[i], res[i])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 10
codeforces
codeforces
codehorses
`

	expect := []int{0, 12}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 9
vkontakte
vcontacte
vkontrakte
vkollapse
vkrokodile
vtopke
vkapuste
vpechke
vk
vcodeforcese
`

	expect := []int{18, 14, 36, 47, 14, 29, 30, 0, 84}

	runSample(t, s, expect)
}
