You are given an array of 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
.

In one operation you split the array into two parts: a non-empty prefix and a non-empty suffix. The value of each part is the bitwise XOR of all elements in it. Next, discard the part with the smaller value. If both parts have equal values, you can choose which one to discard. Replace the array with the remaining part.

The operations are being performed until the length of the array becomes 1
. For each 𝑖
 (1≤𝑖≤𝑛
), determine whether it is possible to achieve the state when only the 𝑖
-th element (with respect to the original numbering) remains.

Formally, you have two numbers 𝑙
 and 𝑟
, initially 𝑙=1
 and 𝑟=𝑛
. The current state of the array is [𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟]
.

As long as 𝑙<𝑟
, you apply the following operation:

Choose an arbitrary 𝑘
 from the set {𝑙,𝑙+1,…,𝑟−1}
. Denote 𝑥=𝑎𝑙⊕𝑎𝑙+1⊕…⊕𝑎𝑘
 and 𝑦=𝑎𝑘+1⊕𝑎𝑘+2⊕…⊕𝑎𝑟
, where ⊕
 denotes the bitwise XOR operation.
If 𝑥<𝑦
, set 𝑙=𝑘+1
.
If 𝑥>𝑦
, set 𝑟=𝑘
.
If 𝑥=𝑦
, either set 𝑙=𝑘+1
, or set 𝑟=𝑘
.
For each 𝑖
 (1≤𝑖≤𝑛
), determine whether it is possible to achieve 𝑙=𝑟=𝑖
.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤10000
). The description of the test cases follows.

The first line of each test case contains one integer 𝑛
 (1≤𝑛≤10000
).

The second line of each test case contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (0≤𝑎𝑖<260
).

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 10000
.

Output
For each test case, output a single string of length 𝑛
 where the 𝑖
-th element is equal to 1 if it is possible to achieve 𝑙=𝑟=𝑖
 and is equal to 0 otherwise.

### ideas
1. 考虑对于一个i来说，要怎么样才能保留它
2. i必须每次都留在大的那边
3. 有点无处下手呐

### solution
1. [reference](https://www.luogu.com.cn/article/2ehhrzmk)
   
