You are given a non-empty string s consisting of lowercase letters. Find the number of pairs of non-overlapping palindromic substrings of this string.

In a more formal way, you have to find the quantity of tuples (a, b, x, y) such that 1 ≤ a ≤ b < x ≤ y ≤ |s| and substrings s[a... b], s[x... y] are palindromes.

A palindrome is a string that can be read the same way from left to right and from right to left. For example, "abacaba", "z", "abba" are palindromes.

A substring s[i... j] (1 ≤ i ≤ j ≤ |s|) of string s = s1s2... s|s| is a string sisi + 1... sj. For example, substring s[2...4] of string s = "abacaba" equals "bac".