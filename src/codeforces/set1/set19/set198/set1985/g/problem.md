Let 𝐷(𝑛)
 represent the sum of digits of 𝑛
. For how many integers 𝑛
 where 10𝑙≤𝑛<10𝑟
 satisfy 𝐷(𝑘⋅𝑛)=𝑘⋅𝐷(𝑛)
? Output the answer modulo 109+7
.

#### ideas
1. 如果k很大, 比如k=10, D(k . n) = D(n)
2. 因为0不贡献，如果 k = 11, D(k.n) != D(n), 除非 n = 0
3. 所以k>= 10 的时候，答案 = 0
4. 当k=2,3... 如果n中存在某个位和k相乘产生进位，也是不行的
5. 所以当n中的digit可以被枚举出来。只要保证产生的数字在范围内即可