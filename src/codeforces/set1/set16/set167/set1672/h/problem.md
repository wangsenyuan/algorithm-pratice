You have a binary string 𝑎
 of length 𝑛
 consisting only of digits 0
 and 1
.

You are given 𝑞
 queries. In the 𝑖
-th query, you are given two indices 𝑙
 and 𝑟
 such that 1≤𝑙≤𝑟≤𝑛
.

Let 𝑠=𝑎[𝑙,𝑟]
. You are allowed to do the following operation on 𝑠
:

Choose two indices 𝑥
 and 𝑦
 such that 1≤𝑥≤𝑦≤|𝑠|
. Let 𝑡
 be the substring 𝑡=𝑠[𝑥,𝑦]
. Then for all 1≤𝑖≤|𝑡|−1
, the condition 𝑡𝑖≠𝑡𝑖+1
 has to hold. Note that 𝑥=𝑦
 is always a valid substring.
Delete the substring 𝑠[𝑥,𝑦]
 from 𝑠
.
For each of the 𝑞
 queries, find the minimum number of operations needed to make 𝑠
 an empty string.

Note that for a string 𝑠
, 𝑠[𝑙,𝑟]
 denotes the subsegment 𝑠𝑙,𝑠𝑙+1,…,𝑠𝑟
.

### ideas
1. 0011  => 这个确实可以在两次内删除掉， 先删除中间的01, 然后删除 01
2. 不考虑子串，在s中，如果有一段连续的长度为l的1， 那么结果不会小于l
3. 考虑 L0101R
4. 感觉就是这段内，连续的l sum(l - 1)?  不是
5. 比如  110110 >> 110 >> 1 => ~
6. 1001110110 => 10110110 => 10110 => 10 => ~
7. 1001110110 => 1001110 => 10110 => 10 => ~
8. 
9. 