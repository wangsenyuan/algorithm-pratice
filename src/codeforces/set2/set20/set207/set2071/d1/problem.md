This is the easy version of the problem. The difference between the versions is that in this version, 𝑙=𝑟
. You can hack only if you solved all versions of this problem.

You are given a positive integer 𝑛
 and the first 𝑛
 terms of an infinite binary sequence 𝑎
, which is defined as follows:

For 𝑚>𝑛
, 𝑎𝑚=𝑎1⊕𝑎2⊕…⊕𝑎⌊𝑚2⌋
∗
.
Your task is to compute the sum of elements in a given range [𝑙,𝑟]
: 𝑎𝑙+𝑎𝑙+1+…+𝑎𝑟
.

### ideas
1. a[i] (0, 1)
2. a[i] = 0的那些，没有贡献。只有a[i] = 1 的需要考虑
3. a[i] = a[1]... a[i/2]
4. 考虑 a[n+1] = a[1] ^ ... ^ a[(n+1) / 2]