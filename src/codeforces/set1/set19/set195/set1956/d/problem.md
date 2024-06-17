Nene gave you an array of integers 𝑎1,𝑎2,…,𝑎𝑛
 of length 𝑛
.

You can perform the following operation no more than 5⋅105
 times (possibly zero):

Choose two integers 𝑙
 and 𝑟
 such that 1≤𝑙≤𝑟≤𝑛
, compute 𝑥
 as MEX({𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟})
, and simultaneously set 𝑎𝑙:=𝑥,𝑎𝑙+1:=𝑥,…,𝑎𝑟:=𝑥
.
Here, MEX
 of a set of integers {𝑐1,𝑐2,…,𝑐𝑘}
 is defined as the smallest non-negative integer 𝑚
 which does not occur in the set 𝑐
.

Your goal is to maximize the sum of the elements of the array 𝑎
. Find the maximum sum and construct a sequence of operations that achieves this sum. Note that you don't need to minimize the number of operations in this sequence, you only should use no more than 5⋅105
 operations in your solution.


#### ideas
1. 没啥头绪
2. 假设选取了区间(l...r) 并将其中所有的数设置成了 mex(l...r) = x
3. 下一次可以再选取到区间(l...r)吗？好像是可以的 
4. 但是这里有多种情况，
   1. 选取后，新的mex(l1...r1) > x, 
   2. 选取后， < x
   3. 选取后 = x
5. 后面两种情况，似乎不是更优的方案， 先不考虑
6. 假设选取后更大了，那么这时候意味着 (l1....r1)区间内，至少包括（0...x-1)
7. 且 （l...r) = x, 所以才能更大 （l1....r1) 不包括（l..r)肯定不包含x,否则的话，可以直接到x+1, 而不需要额外操作（l..r)
8. 考虑最后的情况，应该是类似 [x,...x, y,...y,z...z, ...]
9. 多段连续的相同的数值
10. x = [1...i] 的 mex， y = [i+1, ...j] 的mex （或者原来的值）
11. dp[i]表示前面能够得到的最大值
12. 如果[i...j]能够达到mex x, 那么 dp[j] = dp[i-1] + x * (j - i + 1)  
13. 但是如何知道[i...j]能够达到x呢？
14. 是一次操作就能达到？还是多次操作能够达到？
15. 比如 [i...j]能够达到x, 然后[j...r]有[0...x-1]， 那么[i...r]就能到x+1
16. 怎么感觉 [0....x-1] 如果每个数至少有两个时，那么就可以到达 x+1
17. 好像也不一定， 它们的顺序也有关系比如 【0，1, 2，0， 1, 2】 => 4
18. 但是 [2, 2, 1, 0, 1, 0] 这个是到不了4的， 好像是可以到达4的
19. 那是不是可以推导到n的时候也成立？
20. [1, 1, 2, 2, 0, 0]不成立了
21. 即使成立，似乎要找出那些i也是很麻烦的
22. 等等 n <= 18
23. dp[l...r] 是能够到到的最大的mex
24. dp[l...r] = dp[l+1, r], dp[l, r-1] 或者 = dp[l...i]和i+1....r组成一个新的数组
25. 或者[l...i-1] + dp[i..r]的组合
26. 有其他的组合吗？比如中间和两头的？应该是没有的
27. 不大对～这个无法处理，左边达到x，右边[0...x-1]的情况
28. 