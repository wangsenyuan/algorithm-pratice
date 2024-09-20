You are given a binary string 𝑠
 of length 𝑛
. For each pair of integers (𝑙,𝑟)
 (1≤𝑙≤𝑟≤𝑛)
, count the number of pairs (𝑥,𝑦)
 (𝑙≤𝑥≤𝑦≤𝑟)
 such that the amount of 𝟶
 equals the amount of 𝟷
 in the substring 𝑠𝑥𝑠𝑥+1...𝑠𝑦
.

### ideas
1. s[x..y] 的0和1的数量一致
2. 可以看一下x,y在多少个l...r中间出现过？
3. (n - y) * (x + 1) => (n - x) * sum(x + 1)
4. ???0101???
5. 