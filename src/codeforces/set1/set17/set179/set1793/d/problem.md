In winter, the inhabitants of the Moscow Zoo are very bored, in particular, it concerns gorillas. You decided to entertain them and brought a permutation 𝑝
 of length 𝑛
 to the zoo.

A permutation of length 𝑛
 is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in any order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 occurs twice in the array) and [1,3,4]
 is also not a permutation (𝑛=3
, but 4
 is present in the array).

The gorillas had their own permutation 𝑞
 of length 𝑛
. They suggested that you count the number of pairs of integers 𝑙,𝑟
 (1≤𝑙≤𝑟≤𝑛
) such that MEX([𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟])=MEX([𝑞𝑙,𝑞𝑙+1,…,𝑞𝑟])
.

The MEX
 of the sequence is the minimum integer positive number missing from this sequence. For example, MEX([1,3])=2
, MEX([5])=1
, MEX([3,1,2,6])=4
.

You do not want to risk your health, so you will not dare to refuse the gorillas.


### ideas
1. 略一思考，还有点难～
2. 假设p[i] = 0, q[j] = 0,
3. 那么任何不包括(i, j)的区间[l...r]它们的mex = 0
4. 那是不是就可以这样算呢？
5. 不包括(i, j)的区间 = 总区间数 n * (n + 1) / 2 - 包含它们的区间数 = (i+1) * (n - j) ?
6. 前面处理的是 mex = 0的情况， 
7. 接下来处理mex = 1的情况（它们必须把0包括进去）
8. 好像不大对，比如如果0在两头的时候
9. 那些mex = 0 的区间 = 肯定不包含任何一个0的区间
10. 正难则反。有办法计算出，区间[l...r]中 mex(q) != mex(p)的方式吗？
11. 假设 mex(p[l..r]) = x, mex(q[l...r]) = y
12. 它们不相同
13. 考虑一个序列，如何找到mex
14. 假设序列是 [2,3,1,5,4],
15. mex[3:3] = 2, 其他的等于 1
16. mex[1:3] = 4, 其他的等于 1
17. mex[1:5] = 6, 其他的等于1
18. [1 3 2], [2 1 3]
19. p_mex[1:2] = 2, q_mex[1:2] = 3 但是其他都等于1 
20. p_mex[1:3] = 4, q_mex[1:3] = 4
21. [4, 1, 3, 2], [4, 3, 1, 2]
22. [4, 1, 3] mex = 2
23. [4, 3, 1] mex = 2
24. 这个区间[1:3]它就是包含了1,但是不包含2，所以这样是这样的区间它们的mex = 2
25. 所以如果某个区间包含了1....x-1, 但是不包含x,那么这样的区间是相等的
26. mex = 0的区间假设已经get到了
27. 那么计算 mex = x的区间，就可以用到上面的方式