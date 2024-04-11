Childan is making up a legendary story and trying to sell his forgery — a necklace with a strong sense of "Wu" to the
Kasouras. But Mr. Kasoura is challenging the truth of Childan's story. So he is going to ask a few questions about
Childan's so-called "personal treasure" necklace.

This "personal treasure" is a multiset 𝑆
of 𝑚
"01-strings".

A "01-string" is a string that contains 𝑛
characters "0" and "1". For example, if 𝑛=4
, strings "0110", "0000", and "1110" are "01-strings", but "00110" (there are 5
characters, not 4
) and "zero" (unallowed characters) are not.

Note that the multiset 𝑆
can contain equal elements.

Frequently, Mr. Kasoura will provide a "01-string" 𝑡
and ask Childan how many strings 𝑠
are in the multiset 𝑆
such that the "Wu" value of the pair (𝑠,𝑡)
is not greater than 𝑘
.

Mrs. Kasoura and Mr. Kasoura think that if 𝑠𝑖=𝑡𝑖
(1≤𝑖≤𝑛
) then the "Wu" value of the character pair equals to 𝑤𝑖
, otherwise 0
. The "Wu" value of the "01-string" pair is the sum of the "Wu" values of every character pair. Note that the length of
every "01-string" is equal to 𝑛
.

For example, if 𝑤=[4,5,3,6]
, "Wu" of ("1001", "1100") is 7
because these strings have equal characters only on the first and third positions, so 𝑤1+𝑤3=4+3=7
.

You need to help Childan to answer Mr. Kasoura's queries. That is to find the number of strings in the multiset 𝑆
such that the "Wu" value of the pair is not greater than 𝑘
.

Input
The first line contains three integers 𝑛
, 𝑚
, and 𝑞
(1≤𝑛≤12
, 1≤𝑞,𝑚≤5⋅105
) — the length of the "01-strings", the size of the multiset 𝑆
, and the number of queries.

The second line contains 𝑛
integers 𝑤1,𝑤2,…,𝑤𝑛
(0≤𝑤𝑖≤100
) — the value of the 𝑖
-th caracter.

Each of the next 𝑚
lines contains the "01-string" 𝑠
of length 𝑛
— the string in the multiset 𝑆
.

Each of the next 𝑞
lines contains the "01-string" 𝑡
of length 𝑛
and integer 𝑘
(0≤𝑘≤100
) — the query.

### ideas

1. 很绕的一个题目
2. 好像每位是可以单独计算的，如果t[i] = 0， 那么= cnt[i][0] * w[i], else cnt[i][1] * w[i]
3. 忘记了要找的是符合条件 f(s, t) <= k 的s的数量
4. 对于给定的t, 是否可以找到它(0, 4096)范围内的s，符合条件的进行处理？