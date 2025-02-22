A string t is called an anagram of the string s, if it is possible to rearrange letters in t so that it is identical to the string s. For example, the string "aab" is an anagram of the string "aba" and the string "aaa" is not.

The string t is called a substring of the string s if it can be read starting from some position in the string s. For example, the string "aba" has six substrings: "a", "b", "a", "ab", "ba", "aba".

You are given a string s, consisting of lowercase Latin letters and characters "?". You are also given a string p, consisting of lowercase Latin letters only. Let's assume that a string is good if you can obtain an anagram of the string p from it, replacing the "?" characters by Latin letters. Each "?" can be replaced by exactly one character of the Latin alphabet. For example, if the string p = «aba», then the string "a??" is good, and the string «?bc» is not.

Your task is to find the number of good substrings of the string s (identical substrings must be counted in the answer several times).

Input
The first line is non-empty string s, consisting of no more than 105 lowercase Latin letters and characters "?". The second line is non-empty string p, consisting of no more than 105 lowercase Latin letters. Please note that the length of the string p can exceed the length of the string s.

Output
Print the single number representing the number of good substrings of string s.

Two substrings are considered different in their positions of occurrence are different. Thus, if some string occurs several times, then it should be counted the same number of times.

