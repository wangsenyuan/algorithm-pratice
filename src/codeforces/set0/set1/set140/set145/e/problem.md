Petya loves lucky numbers very much. Everybody knows that lucky numbers are positive integers whose decimal record
contains only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

Petya brought home string s with the length of n. The string only consists of lucky digits. The digits are numbered from
the left to the right starting with 1. Now Petya should execute m queries of the following form:

switch l r — "switch" digits (i.e. replace them with their opposites) at all positions with indexes from l to r,
inclusive: each digit 4 is replaced with 7 and each digit 7 is replaced with 4 (1 ≤ l ≤ r ≤ n);
count — find and print on the screen the length of the longest non-decreasing subsequence of string s.
Subsequence of a string s is a string that can be obtained from s by removing zero or more of its elements. A string is
called non-decreasing if each successive digit is not less than the previous one.

Help Petya process the requests.

### ideas

1. 将4看作0，7看作1，就变成了0/1字符串
2. 不考虑switch的情况，如何count呢？
3. 正常情况下，应该是一个位置i的前面的所有的0+后面所有的1（它自己的值不重要)
4. 但如何计算一段[l..r]内的呢？
5. 假设每段上维护了以下值value (最大值), cnt[0]/cnt[1]
6. 那么value(l...r) = max(value of first half + 右边cnt[1], 左边cnt[0] + value of second half)
7. 但是翻转的时候， value怎么变化呢？