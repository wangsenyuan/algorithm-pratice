You are given two lists of segments [𝑎𝑙1,𝑎𝑟1],[𝑎𝑙2,𝑎𝑟2],…,[𝑎𝑙𝑛,𝑎𝑟𝑛]
 and [𝑏𝑙1,𝑏𝑟1],[𝑏𝑙2,𝑏𝑟2],…,[𝑏𝑙𝑛,𝑏𝑟𝑛]
.

Initially, all segments [𝑎𝑙𝑖,𝑎𝑟𝑖]
 are equal to [𝑙1,𝑟1]
 and all segments [𝑏𝑙𝑖,𝑏𝑟𝑖]
 are equal to [𝑙2,𝑟2]
.

In one step, you can choose one segment (either from the first or from the second list) and extend it by 1
. In other words, suppose you've chosen segment [𝑥,𝑦]
 then you can transform it either into [𝑥−1,𝑦]
 or into [𝑥,𝑦+1]
.

Let's define a total intersection 𝐼
 as the sum of lengths of intersections of the corresponding pairs of segments, i.e. ∑𝑖=1𝑛intersection_length([𝑎𝑙𝑖,𝑎𝑟𝑖],[𝑏𝑙𝑖,𝑏𝑟𝑖])
. Empty intersection has length 0
 and length of a segment [𝑥,𝑦]
 is equal to 𝑦−𝑥
.

What is the minimum number of steps you need to make 𝐼
 greater or equal to 𝑘
?

Input
The first line contains the single integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases.

The first line of each test case contains two integers 𝑛
 and 𝑘
 (1≤𝑛≤2⋅105
; 1≤𝑘≤109
) — the length of lists and the minimum required total intersection.

The second line of each test case contains two integers 𝑙1
 and 𝑟1
 (1≤𝑙1≤𝑟1≤109
) — the segment all [𝑎𝑙𝑖,𝑎𝑟𝑖]
 are equal to initially.

The third line of each test case contains two integers 𝑙2
 and 𝑟2
 (1≤𝑙2≤𝑟2≤109
) — the segment all [𝑏𝑙𝑖,𝑏𝑟𝑖]
 are equal to initially.

It's guaranteed that the sum of 𝑛
 doesn't exceed 2⋅105
.

### ideas
1. 分多钟情况， 一开始两个是完全分开的，先要将两者变成接触，并使得两个差距最少
2. 假设 r1 < l2, 那么可以操作 w = l2 - r1
3. 然后，这个时候，应该固定短的区域不变，让长的区域延展到短的边界
4. 然后再往反方向延展，这样可以最多 r2 - r1的区域时间内，得到重叠
