This is the easy version of the problem. The difference between the versions is that in this version, string 𝑡
 consists of only '0' and '1'. You can hack only if you solved all versions of this problem.

Kevin has a binary string 𝑠
 of length 𝑛
. Kevin can perform the following operation:

Choose two adjacent blocks of 𝑠
 and swap them.
A block is a maximal substring∗
 of identical characters. Formally, denote 𝑠[𝑙,𝑟]
 as the substring 𝑠𝑙𝑠𝑙+1…𝑠𝑟
. A block is 𝑠[𝑙,𝑟]
 satisfying:

𝑙=1
 or 𝑠𝑙≠𝑠𝑙−1
.
𝑠𝑙=𝑠𝑙+1=…=𝑠𝑟
.
𝑟=𝑛
 or 𝑠𝑟≠𝑠𝑟+1
.
Adjacent blocks are two blocks 𝑠[𝑙1,𝑟1]
 and 𝑠[𝑙2,𝑟2]
 satisfying 𝑟1+1=𝑙2
.

For example, if 𝑠=𝟶𝟶𝟶1100𝟷𝟷𝟷
, Kevin can choose the two blocks 𝑠[4,5]
 and 𝑠[6,7]
 and swap them, transforming 𝑠
 into 𝟶𝟶𝟶0011𝟷𝟷𝟷
.

Given a string 𝑡
 of length 𝑛
 consisting of '0', '1' and '?', Kevin wants to determine the minimum number of operations required to perform such that for any index 𝑖
 (1≤𝑖≤𝑛
), if 𝑡𝑖≠
 '?' then 𝑠𝑖=𝑡𝑖
. If it is impossible, output −1
.

### ideas
1. 一个连续的区域，通过swap后，只会变大（如果在边界处，可能不变化）
2. 假设s将连续的block，编码成数字后，是 1,2,3,4
3. 那么交换2，3 =》4，6
4. 如果交换3,4 => 1,6,3
5. 交换1,2 => 2,4,4 => 6,4
6. 假设对t进行编码后，t[1], t[2] ... t[i]
7. 从上面的例子可以看出，一次交换后，arr的size - 2, 所以，结果成立的条件之一，是 size(s) - size(t) 是偶数（且如果存在答案， 差 / 2就是答案）
8. 那么剩下的就是如何确定答案是存在的
9. 假设t[i] != s[i], 那么t[i] = s[i+1] + s[i+3] + s[i+5]... s[j] 这个成立?
10. 假设在 S[i] != T[i] 这里是编码后的结果
11. 如果S[i] > T[i] 表明这里有更多的0（或者1）=》 false(因为区间无法变小)
12. 那么只能 S[i] < T[i] => 也就是说要交换过来T[i] - S[i] 个1过来 11001  / 11100
13. 那么此时就是把最靠近的1的这些部分移动到这里来