Catherine received an array of integers as a gift for March 8. Eventually she grew bored with it, and she started calculated various useless characteristics for it. She succeeded to do it for each one she came up with. But when she came up with another one — xor of all pairwise sums of elements in the array, she realized that she couldn't compute it for a very large array, thus she asked for your help. Can you do it? Formally, you need to compute

(𝑎1+𝑎2)⊕(𝑎1+𝑎3)⊕…⊕(𝑎1+𝑎𝑛)⊕(𝑎2+𝑎3)⊕…⊕(𝑎2+𝑎𝑛)…⊕(𝑎𝑛−1+𝑎𝑛)
Here 𝑥⊕𝑦
 is a bitwise XOR operation (i.e. 𝑥
 ^ 𝑦
 in many modern programming languages). You can read about it in Wikipedia:

 ### ideas
 1. 先考虑a1, (a1 + a2) ^ (a1 + a3) ^ ... ^ (a1 + an)
 2. 首先这个式子有n-1项
 3. 考虑它可能产生的最高位h，如果有奇数个和出现了h位，那么就可以贡献h
 4. 这样也不对。因为还需要考虑a2, a3, ... an-1等
 5. 考虑最高位的数字a[i], 它的最高位h，会有多少个数和它相加能够出现呢？