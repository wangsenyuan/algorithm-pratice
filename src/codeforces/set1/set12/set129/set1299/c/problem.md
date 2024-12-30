There are 𝑛
 water tanks in a row, 𝑖
-th of them contains 𝑎𝑖
 liters of water. The tanks are numbered from 1
 to 𝑛
 from left to right.

You can perform the following operation: choose some subsegment [𝑙,𝑟]
 (1≤𝑙≤𝑟≤𝑛
), and redistribute water in tanks 𝑙,𝑙+1,…,𝑟
 evenly. In other words, replace each of 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
 by 𝑎𝑙+𝑎𝑙+1+⋯+𝑎𝑟𝑟−𝑙+1
. For example, if for volumes [1,3,6,7]
 you choose 𝑙=2,𝑟=3
, new volumes of water will be [1,4.5,4.5,7]
. You can perform this operation any number of times.

What is the lexicographically smallest sequence of volumes of water that you can achieve?

As a reminder:

A sequence 𝑎
 is lexicographically smaller than a sequence 𝑏
 of the same length if and only if the following holds: in the first (leftmost) position where 𝑎
 and 𝑏
 differ, the sequence 𝑎
 has a smaller element than the corresponding element in 𝑏
.

### ideas
1. 为了让当前位置足够小，应该找到包含当前位置（开始）的sum / 长度最小的那段
2. 这是一个难点。
3. 还有一个难点，res[i]计算出来后，后续的也需要变成avg
4. 从例子的结果来看，这个应该是一个递增序列
5. 不过这个也很证明，如果不是一个递增序列，假设 res[i] > res[i+1]
6. 那么在处理i的时候，将i+1包含进去，就肯定能够得到一个更优的结果
7. 所以 res[1] <= res[2] ... <= res[n]
8. res[i]能达到的最小值 = (pref(j) - pref(i)) / (j - i)
9.  v = (pj - pi) / (j - i) 
10. 怎么更新这个东西，并重新获得最小值
11. (pj - pi) / (i - j)能得到最小值
12. 也没法二分～
13. 如果i的后面有一个j，a[j] < a[i], 那么这个j有可能是一个答案（但却不一定）
14. 如果没有这样的j，那么a[i]本身就是最小的
15. 或者是否能够检查, avg <= x一个具体的值是否成立？
16. 假设在处理到i时的最小的avg=x/y
17. 现在处理i-1, w
18. 如果 w >= x/y, 那么似乎(w + x) / (y + 1)仍然是最小的？
19. 如果 w <= x/y, 那么w就是最优的
20. 但是 如果 w >= x/y, 那么就不是的
21. (w + x) / (y + 1) = u
22. w + x = (y + 1) * u
23. 假设知道x/y，如何根据新的w，找到最小的y，使的(w + x) / (y+1) 最小
24. 所有的数+w这个比较好弄，就是个range update
25. 但是如何知道这个区间内的最小的 (w + x) / (y+1)呢？

### solution

Let's try to make the operation simpler. When we apply the operation, only the sum of the segment matters. And so let's instead define the operation on prefix sum array:

Replace each of 𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟
 by 𝑝𝑖=𝑝𝑙−1+𝑝𝑟−𝑝𝑙−1𝑟−𝑙+1⋅(𝑖−𝑙+1)
. You may see how similar it is to a line function. Hence we get the idea to plot points (𝑖,𝑝𝑖)
 ((0,𝑝0=0)
 included), and our operation is just drawing a line between 2
 points on integer 𝑥
 coordinates.

Nicely if sequence 𝑎
 is lexicographically smaller than sequence 𝑏
, then prefix sum array of 𝑎
 is smaller than prefix sum array of 𝑏
.

So we need to find the lexicographically smallest array 𝑝
. And then it is easy to see the lexicographically smallest sequence 𝑝
 will be the lower part of the convex hull.

