You are given two arrays 𝑎1,𝑎2,…,𝑎𝑛
 and 𝑏1,𝑏2,…,𝑏𝑛
.

You must perform the following operation exactly once:

choose any indices 𝑙
 and 𝑟
 such that 1≤𝑙≤𝑟≤𝑛
;
swap 𝑎𝑖
 and 𝑏𝑖
 for all 𝑖
 such that 𝑙≤𝑖≤𝑟
.
Find the maximum possible value of gcd(𝑎1,𝑎2,…,𝑎𝑛)+gcd(𝑏1,𝑏2,…,𝑏𝑛)
 after performing the operation exactly once. Also find the number of distinct pairs (𝑙,𝑟)
 which achieve the maximum value.



### ideas
1. gcd(a) 不会大于a[0]（假设a[0]没有被交换）的因子
2. 对于a[0]的因子x,假设要得到x的结果，必须交换区间[lx, rx]
3. 假设两个因子x, y， 且x > y, 那么[lx, rx]必然是包含区间[ly, ry]的
4. 如果在数组b中， [lx, rx] 可以得到gcd(x)， 那么交换后，至少可以得到x
5. 且扩大(lx, rx) 对于a来说，肯定可以得到x，
6. 比如，交换[lx, rx] 肯定也能得到y