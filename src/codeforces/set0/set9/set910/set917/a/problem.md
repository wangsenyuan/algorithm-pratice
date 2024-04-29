Thus, he came up with a puzzle to tell his mom his coordinates. His coordinates are the answer to the following problem.

A string consisting only of parentheses ('(' and ')') is called a bracket sequence. Some bracket sequence are called
correct bracket sequences. More formally:

Empty string is a correct bracket sequence.
if s is a correct bracket sequence, then (s) is also a correct bracket sequence.
if s and t are correct bracket sequences, then st (concatenation of s and t) is also a correct bracket sequence.
A string consisting of parentheses and question marks ('?') is called pretty if and only if there's a way to replace
each question mark with either '(' or ')' such that the resulting string is a non-empty correct bracket sequence.

Will gave his mom a string s consisting of parentheses and question marks (using Morse code through the lights) and his
coordinates are the number of pairs of integers (l, r) such that 1 ≤ l ≤ r ≤ |s| and the string slsl + 1... sr is
pretty, where si is i-th character of s.

Joyce doesn't know anything about bracket sequences, so she asked for your help.

### ideas

1. dp[i..j] = dp[i+1..j] + dp[i...j-1] - dp[i+1...j-1] + v if i...j is good