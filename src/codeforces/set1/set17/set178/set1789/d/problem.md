Serval has two 𝑛
-bit binary integer numbers 𝑎
 and 𝑏
. He wants to share those numbers with Toxel.

Since Toxel likes the number 𝑏
 more, Serval decides to change 𝑎
 into 𝑏
 by some (possibly zero) operations. In an operation, Serval can choose any positive integer 𝑘
 between 1
 and 𝑛
, and change 𝑎
 into one of the following number:

𝑎⊕(𝑎≪𝑘)
𝑎⊕(𝑎≫𝑘)
In other words, the operation moves every bit of 𝑎
 left or right by 𝑘
 positions, where the overflowed bits are removed, and the missing bits are padded with 0
. The bitwise XOR of the shift result and the original 𝑎
 is assigned back to 𝑎
.

Serval does not have much time. He wants to perform no more than 𝑛
 operations to change 𝑎
 into 𝑏
. Please help him to find out an operation sequence, or determine that it is impossible to change 𝑎
 into 𝑏
 in at most 𝑛
 operations. You do not need to minimize the number of operations.

In this problem, 𝑥⊕𝑦
 denotes the bitwise XOR operation of 𝑥
 and 𝑦
. 𝑎≪𝑘
 and 𝑎≫𝑘
 denote the logical left shift and logical right shift.

### ideas
1. a ^ (a >> k) 那么a的最高k位是保留的
2. b = 11000 = x ^ (x >> ?) or x ^ (x << ?)
3. 如果是左移k, 那么必须满足 b[i] = 0, 所以 x[n] = x[n-k], x[n-1] = x[n-k-1] ... x[n-k+1] = x[1]
4. 123456   001234
5. x[i] != x[i-k]  (如果 i > k)
6. 如果 b[i] = 0 => x[i] = x[i-k]