And where the are the phone numbers?

You are given a string s consisting of lowercase English letters and an integer k. Find the lexicographically smallest
string t of length k, such that its set of letters is a subset of the set of letters of s and s is lexicographically
smaller than t.

It's guaranteed that the answer exists.

Note that the set of letters is a set, not a multiset. For example, the set of letters of abadaba is {a, b, d}.

String p is lexicographically smaller than string q, if p is a prefix of q, is not equal to q or there exists i, such
that pi<qi and for all j<i it is satisfied that pj = qj. For example, abc is lexicographically smaller than abcd , abd
is lexicographically smaller than abec, afa is not lexicographically smaller than ab and a is not lexicographically
smaller than a.

### ideas

1. 假设在第i位，使得s[i] < t[i]，前面的都一致，i后面使用最小值
2. 显然i越大越好
3. 那岂不是就从往前判断，能不能存在x > s[i]就可以了吗？