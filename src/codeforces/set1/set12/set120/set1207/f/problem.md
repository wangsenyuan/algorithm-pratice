You are given an array 𝑎
 consisting of 500000
 integers (numbered from 1
 to 500000
). Initially all elements of 𝑎
 are zero.

You have to process two types of queries to this array:

1
 𝑥
 𝑦
 — increase 𝑎𝑥
 by 𝑦
;
2
 𝑥
 𝑦
 — compute ∑𝑖∈𝑅(𝑥,𝑦)𝑎𝑖
, where 𝑅(𝑥,𝑦)
 is the set of all integers from 1
 to 500000
 which have remainder 𝑦
 modulo 𝑥
.
Can you process all the queries?

### ideas
1. i % x = y
2. i = y, x + y, x + 2 * y .... 
3. 如果x比较大（超过sqrt(5..))， 那么就brute force i
4. 如果x比较小 , 那么y也比较小， 如何利用y呢？
5. 可以准备一个额外的数组，分别表示dp[x][y] 表示当除数是x，余数是y时的和
6. 那么在修改i的时候，就可以去更新一边就好了
7. got