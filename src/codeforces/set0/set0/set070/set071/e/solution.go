package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		m := len(cur)
		for i := 0; i < m-1; i++ {
			buf.WriteString(cur[i])
			if i < m-2 {
				buf.WriteString("+")
			}
		}
		buf.WriteString("->")
		buf.WriteString(cur[m-1])
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

var atoms map[string]int

func init() {
	atoms = map[string]int{
		"H":  1,
		"He": 2,
		"Li": 3,
		"Be": 4,
		"B":  5,
		"C":  6,
		"N":  7,
		"O":  8,
		"F":  9,
		"Ne": 10,
		"Na": 11,
		"Mg": 12,
		"Al": 13,
		"Si": 14,
		"P":  15,
		"S":  16,
		"Cl": 17,
		"Ar": 18,
		"K":  19,
		"Ca": 20,
		"Sc": 21,
		"Ti": 22,
		"V":  23,
		"Cr": 24,
		"Mn": 25,
		"Fe": 26,
		"Co": 27,
		"Ni": 28,
		"Cu": 29,
		"Zn": 30,
		"Ga": 31,
		"Ge": 32,
		"As": 33,
		"Se": 34,
		"Br": 35,
		"Kr": 36,
		"Rb": 37,
		"Sr": 38,
		"Y":  39,
		"Zr": 40,
		"Nb": 41,
		"Mo": 42,
		"Tc": 43,
		"Ru": 44,
		"Rh": 45,
		"Pd": 46,
		"Ag": 47,
		"Cd": 48,
		"In": 49,
		"Sn": 50,
		"Sb": 51,
		"Te": 52,
		"I":  53,
		"Xe": 54,
		"Cs": 55,
		"Ba": 56,
		"La": 57,
		"Ce": 58,
		"Pr": 59,
		"Nd": 60,
		"Pm": 61,
		"Sm": 62,
		"Eu": 63,
		"Gd": 64,
		"Tb": 65,
		"Dy": 66,
		"Ho": 67,
		"Er": 68,
		"Tm": 69,
		"Yb": 70,
		"Lu": 71,
		"Hf": 72,
		"Ta": 73,
		"W":  74,
		"Re": 75,
		"Os": 76,
		"Ir": 77,
		"Pt": 78,
		"Au": 79,
		"Hg": 80,
		"Tl": 81,
		"Pb": 82,
		"Bi": 83,
		"Po": 84,
		"At": 85,
		"Rn": 86,
		"Fr": 87,
		"Ra": 88,
		"Ac": 89,
		"Th": 90,
		"Pa": 91,
		"U":  92,
		"Np": 93,
		"Pu": 94,
		"Am": 95,
		"Cm": 96,
		"Bk": 97,
		"Cf": 98,
		"Es": 99,
		"Fm": 100,
	}
}

func solve(source string, target string) [][]string {
	ss := strings.Split(source, " ")
	tt := strings.Split(target, " ")
	n := len(ss)
	m := len(tt)

	vals := make([]int, m)
	for i := 0; i < m; i++ {
		vals[i] = atoms[tt[i]]
	}

	dp := make([]int, 1<<n)

	// 然后把那些可以从source 组成 target 的， 连一条线
	for state := 1; state < 1<<n; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				dp[state] += atoms[ss[i]]
			}
		}
	}
	fp := make([]int, 1<<n)
	for i := 0; i < len(fp); i++ {
		fp[i] = -1
	}

	prev := make([]int, 1<<n)

	T := 1 << n

	fp[0] = 0

	for state := 0; state < T; state++ {
		tmp := (T - 1) ^ state
		j := fp[state]
		if j < 0 {
			continue
		}
		for next := tmp; next > 0; next = (next - 1) & tmp {
			if dp[next] == vals[j] {
				fp[state|next] = j + 1
				prev[state|next] = state
			}
		}
	}
	if fp[T-1] != m {
		return nil
	}

	ans := make([][]string, m)

	cur := T - 1
	for i := m - 1; i >= 0; i-- {
		state := cur ^ prev[cur]

		for j := 0; j < n; j++ {
			if (state>>j)&1 == 1 {
				ans[i] = append(ans[i], ss[j])
			}
		}

		ans[i] = append(ans[i], tt[i])

		cur = prev[cur]
	}

	return ans
}
