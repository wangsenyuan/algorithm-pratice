The string t1t2…tk
 is good if each letter of this string belongs to at least one palindrome of length greater than 1.

A palindrome is a string that reads the same backward as forward. For example, the strings A, BAB, ABBA, BAABBBAAB are palindromes, but the strings AB, ABBBAA, BBBA are not.

Here are some examples of good strings:

t
 = AABBB (letters t1
, t2
 belong to palindrome t1…t2
 and letters t3
, t4
, t5
 belong to palindrome t3…t5
);
t
 = ABAA (letters t1
, t2
, t3
 belong to palindrome t1…t3
 and letter t4
 belongs to palindrome t3…t4
);
t
 = AAAAA (all letters belong to palindrome t1…t5
);
You are given a string s
 of length n
, consisting of only letters A and B.

You have to calculate the number of good substrings of string s
.


### ideas
1. 似乎除了BA，所有的substring都是好的
2. 还有1