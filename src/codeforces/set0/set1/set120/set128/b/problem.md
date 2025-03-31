One day in the IT lesson Anna and Maria learned about the lexicographic order.

String x is lexicographically less than string y, if either x is a prefix of y (and x ≠ y), or there exists such i (1 ≤ i ≤ min(|x|, |y|)), that xi < yi, and for any j (1 ≤ j < i) xj = yj. Here |a| denotes the length of the string a. The lexicographic comparison of strings is implemented by operator < in modern programming languages​​.

The teacher gave Anna and Maria homework. She gave them a string of length n. They should write out all substrings of the given string, including the whole initial string, and the equal substrings (for example, one should write out the following substrings from the string "aab": "a", "a", "aa", "ab", "aab", "b"). The resulting strings should be sorted in the lexicographical order. The cunning teacher doesn't want to check all these strings. That's why she said to find only the k-th string from the list. Help Anna and Maria do the homework.

### ideas
1. trie？