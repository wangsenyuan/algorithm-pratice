You are given a string s consisting only of characters 0 and 1. A substring [l, r] of s is a string slsl + 1sl + 2...
sr, and its length equals to r - l + 1. A substring is called balanced if the number of zeroes (0) equals to the number
of ones in this substring.

You have to determine the length of the longest balanced substring of s.

Input
The first line contains n (1 ≤ n ≤ 100000) — the number of characters in s.

The second line contains a string s consisting of exactly n characters. Only characters 0 and 1 can appear in s.

Output
If there is no non-empty balanced substring in s, print 0. Otherwise, print the length of the longest balanced
substring.

### ideas

1. 前缀和相等的两端之间的长度为可能答案