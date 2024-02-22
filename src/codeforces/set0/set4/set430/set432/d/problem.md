You have a string s = s1s2...s|s|, where |s| is the length of string s, and si its i-th character.

Let's introduce several definitions:

A substring s[i..j] (1 ≤ i ≤ j ≤ |s|) of string s is string sisi + 1...sj.
The prefix of string s of length l (1 ≤ l ≤ |s|) is string s[1..l].
The suffix of string s of length l (1 ≤ l ≤ |s|) is string s[|s| - l + 1..|s|].
Your task is, for any prefix of string s which matches a suffix of string s, print the number of times it occurs in
string s as a substring.

### thoughts

1. lps的计数？