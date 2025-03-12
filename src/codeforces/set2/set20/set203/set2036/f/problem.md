The first line of input contains a single integer 𝑡
 (1≤𝑡≤104)
 — the number of XOR queries on the segment. The following 𝑡
 lines contain the queries, each consisting of the integers 𝑙
, 𝑟
, 𝑖
, 𝑘
 (1≤𝑙≤𝑟≤1018
, 0≤𝑖≤30
, 0≤𝑘<2𝑖)
.

Output
For each query, output a single integer — the XOR of all integers 𝑥
 in the range [𝑙,𝑟]
 such that 𝑥≢𝑘mod2𝑖
.


### ideas. 
1. 也就是说在区间l...r中间，寻找x 的低i位不是k的情况
2. 还需要把这些x xor在一起～～～
3. 