You are given an array 𝑏1,𝑏2,…,𝑏𝑛
.

An anonymous informant has told you that the array 𝑏
was obtained as follows: initially, there existed an array 𝑎1,𝑎2,…,𝑎𝑛
, after which the following two-component operation was performed 𝑘
times:

A fixed point†
𝑥
of the array 𝑎
was chosen.
Then, the array 𝑎
was cyclically shifted to the left‡
exactly 𝑥
times.
As a result of 𝑘
such operations, the array 𝑏1,𝑏2,…,𝑏𝑛
was obtained. You want to check if the words of the anonymous informant can be true or if they are guaranteed to be
false.

†
A number 𝑥
is called a fixed point of the array 𝑎1,𝑎2,…,𝑎𝑛
if 1≤𝑥≤𝑛
and 𝑎𝑥=𝑥
.

‡
A cyclic left shift of the array 𝑎1,𝑎2,…,𝑎𝑛
is the array 𝑎2,…,𝑎𝑛,𝑎1
.

### thoughts

1. 如果存在a 是b从某个地方开始的(环)n个元素
2. 考虑那些<= n的元素，它们可以作为fixed point
3. 这些数和k有什么关系？
4. 感觉和最大公倍数有关系
5. k % lcm(dist)