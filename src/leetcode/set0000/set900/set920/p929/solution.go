package p929

import "strings"

func numUniqueEmails(emails []string) int {
	xs := make(map[string]bool)

	for _, email := range emails {
		x := formalize(email)
		xs[x] = true
	}

	return len(xs)
}

func formalize(email string) string {
	ss := strings.Split(email, "@")
	name := ss[0]
	domain := ss[1]

	for i := 0; i < len(name); i++ {
		if name[i] == '+' {
			// first +
			name = name[:i]
			break
		}
	}
	bs := []byte(name)
	var j int
	for i := 0; i < len(bs); i++ {
		if bs[i] == '.' {
			continue
		}
		bs[j] = bs[i]
		j++
	}
	name = string(bs[:j])
	return name + "@" + domain
}
