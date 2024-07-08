This is an easier version of the problem with smaller constraints.

Korney Korneevich dag up an array 𝑎
 of length 𝑛
. Korney Korneevich has recently read about the operation bitwise XOR, so he wished to experiment with it. For this purpose, he decided to find all integers 𝑥≥0
 such that there exists an increasing subsequence of the array 𝑎
, in which the bitwise XOR of numbers is equal to 𝑥
.

It didn't take a long time for Korney Korneevich to find all such 𝑥
, and he wants to check his result. That's why he asked you to solve this problem!

A sequence 𝑠
 is a subsequence of a sequence 𝑏
 if 𝑠
 can be obtained from 𝑏
 by deletion of several (possibly, zero or all) elements.

A sequence 𝑠1,𝑠2,…,𝑠𝑚
 is called increasing if 𝑠1<𝑠2<…<𝑠𝑚
.

### ideas
1. a[i] < 500, 那么就可以使用dp了
2. dp[x] 表示存在能产生xor = x的递增序列，且最这个序列的最后一个元素是dp[x]
3. 