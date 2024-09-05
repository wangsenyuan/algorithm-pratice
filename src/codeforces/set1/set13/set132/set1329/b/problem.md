Dreamoon likes sequences very much. So he created a problem about the sequence that you can't find in OEIS:

You are given two integers 𝑑,𝑚
, find the number of arrays 𝑎
, satisfying the following constraints:

The length of 𝑎
 is 𝑛
, 𝑛≥1
1≤𝑎1<𝑎2<⋯<𝑎𝑛≤𝑑
Define an array 𝑏
 of length 𝑛
 as follows: 𝑏1=𝑎1
, ∀𝑖>1,𝑏𝑖=𝑏𝑖−1⊕𝑎𝑖
, where ⊕
 is the bitwise exclusive-or (xor). After constructing an array 𝑏
, the constraint 𝑏1<𝑏2<⋯<𝑏𝑛−1<𝑏𝑛
 should hold.
Since the number of possible arrays may be too large, you need to find the answer modulo 𝑚
.

Input
The first line contains an integer 𝑡
 (1≤𝑡≤100
) denoting the number of test cases in the input.

Each of the next 𝑡
 lines contains two integers 𝑑,𝑚
 (1≤𝑑,𝑚≤109
).

Note that 𝑚
 is not necessary the prime!

Output
For each test case, print the number of arrays 𝑎
, satisfying all given constrains, modulo 𝑚
.


### ideas
1. 考虑 1 <= a[1] <... <= a[n] <= d
2. 且a的前缀异或递增
3. b[i] = b[i-1] ^ a[i]
4. a[i] = b[i] ^ b[i-1] 且 b[i-1] < b[i]
5. a需要满足什么性质
6. b[1] = a[1], b[2] = b[1] ^ a[2]
7. b[1] < b[2] => a[1] < a[1] ^ a[2]
8. 如果a[1]和a[2]的最高位相同，显然是不行的。因为这样子a[1] > a[1] ^ a[2]了
9. 只要a[1]的最高位 < a[2]的最高位，上面那个式子，就是成立的，且也满足 a[1] <= a[2]的条件
10. a[2]和a[3]呢？ b[2] < b[3]
11. b[2] = a[1] ^ a[2]
12. b[3] = a[1] ^ a[2] ^ a[3]
13. 那么这里已经知道 b[2]的最高位 = a[2]的最高位
14. 如果a[3]的最高位 > a[2]的最高位，b[2] < b[3]肯定是成立的
15. 如果最高位相同呢, b[2] < b[3] 不成立（也没有成立的可能性)
16. 如果 a[3]的最高位 < a[2]的最高位呢？
17. 也不行，这样子违反了 a[2] <= a[3]的条件
18. 所以，a完全有最高位决定
19. n不超过d的最高位，假设假设 a[i] = 0 表示这一位没有出现
20. 那么(1 << i)