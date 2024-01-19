Monocarp finally got the courage to register on ForceCoders. He came up with a handle but is still thinking about the
password.

He wants his password to be as strong as possible, so he came up with the following criteria:

the length of the password should be exactly 𝑚
;
the password should only consist of digits from 0
to 9
;
the password should not appear in the password database (given as a string 𝑠
) as a subsequence (not necessarily contiguous).
Monocarp also came up with two strings of length 𝑚
: 𝑙
and 𝑟
, both consisting only of digits from 0
to 9
. He wants the 𝑖
-th digit of his password to be between 𝑙𝑖
and 𝑟𝑖
, inclusive.

Does there exist a password that fits all criteria?

### thoughts

1. m <= 10, 每个字符10， 10 ** m种可能性
2. 所以问题的关键是给定字符x,判断它是否是s的子序列
3. 但是如果每个x都去判断，复杂性 100 * O(len(s))
4. 所以需要更快的方式
5. 如果能计算出所有s的长度为m的子序列，