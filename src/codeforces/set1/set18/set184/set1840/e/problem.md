You are given two strings of equal length 𝑠1
and 𝑠2
, consisting of lowercase Latin letters, and an integer 𝑡
.

You need to answer 𝑞
queries, numbered from 1
to 𝑞
. The 𝑖
-th query comes in the 𝑖
-th second of time. Each query is one of three types:

block the characters at position 𝑝𝑜𝑠
(indexed from 1
) in both strings for 𝑡
seconds;
swap two unblocked characters;
determine if the two strings are equal at the time of the query, ignoring blocked characters.
Note that in queries of the second type, the characters being swapped can be from the same string or from 𝑠1
and 𝑠2
.

Input
The first line of the input contains a single integer 𝑇
(1≤𝑇≤104
) — the number of test cases.

Then follow the descriptions of the test cases.

The first line of each test case contains a string 𝑠1
consisting of lowercase Latin letters (length no more than 2⋅105
).

The second line of each test case contains a string 𝑠2
consisting of lowercase Latin letters (length no more than 2⋅105
).

The strings have equal length.

The third line of each test case contains two integers 𝑡
and 𝑞
(1≤𝑡,𝑞≤2⋅105
). The number 𝑡
indicates the number of seconds for which a character is blocked. The number 𝑞
corresponds to the number of queries.

Each of the next 𝑞
lines of each test case contains a single query. Each query is one of three types:

"1 𝑝𝑜𝑠
" — block the characters at position 𝑝𝑜𝑠
in both strings for 𝑡
seconds;
"2 1/2 𝑝𝑜𝑠1 1/2 𝑝𝑜𝑠2
" — swap two unblocked characters. The second number in the query indicates the number of the string from which the
first character for the swap is taken. The third number in the query indicates the position in that string of that
character. The fourth number in the query indicates the number of the string from which the second character for the
swap is taken. The fifth number in the query indicates the position in that string of that character;
"3
" — determine if the two strings are equal at the time of the query, ignoring blocked characters.
For queries of the first type, it is guaranteed that at the time of the query, the characters at position 𝑝𝑜𝑠
are not blocked.

For queries of the second type, it is guaranteed that the characters being swapped are not blocked.

All values of 𝑝𝑜𝑠,𝑝𝑜𝑠1,𝑝𝑜𝑠2
are in the range from 1
to the length of the strings.

The sum of the values of 𝑞
over all test cases, as well as the total length of the strings 𝑠1
, does not exceed 2⋅105
.

### thoughts

1. 假设没有操作1和操作2，如何快速的判断s1 == s2
2. divide and conquer?
3. block掉一个后，可以重新计算一遍, 高度是logn, 所以复杂度也是logn
4. swap也一样
5. 到时间unblock也一样