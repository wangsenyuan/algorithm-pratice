You are given an array 𝑎
 of length 𝑛
 and an integer 𝑘
. You are also given a binary array 𝑏
 of length 𝑛
.

You can perform the following operation at most 𝑘
 times:

Select an index 𝑖
 (1≤𝑖≤𝑛
) such that 𝑏𝑖=1
. Set 𝑎𝑖=𝑎𝑖+1
 (i.e., increase 𝑎𝑖
 by 1
).
Your score is defined to be max𝑖=1𝑛(𝑎𝑖+median(𝑐𝑖))
, where 𝑐𝑖
 denotes the array of length 𝑛−1
 that you get by deleting 𝑎𝑖
 from 𝑎
. In other words, your score is the maximum value of 𝑎𝑖+median(𝑐𝑖)
 over all 𝑖
 from 1
 to 𝑛
.

Find the maximum score that you can achieve if you perform the operations optimally.

For an arbitrary array 𝑝
, median(𝑝)
 is defined as the ⌊|𝑝|+12⌋
-th smallest element of 𝑝
. For example, median([3,2,1,3])=2
 and median([6,2,4,5,1])=4
.

### ideas
1. 还是有点搞不清楚
2. 排除掉一个后，c[i]要怎么计算？这个确实可以二分
3. 但是这样子进行处理，也是 n * n * lg(n)
4. 而且隐隐觉得，不需要每个都处理
5. 如果b[i] = 1, 似乎还是应该把所有的k都加到最大的a[i]上，而不是增加到其他b[?] = 1 
6. 如果b[i] = 0, 那么此时，就是所有b[i] = 1 的，加其他b[i] = 0 的部分
7. 是不是也应该选择最大的那个a[i] (b[i] = 0)作为，然后去加其他的部分
8. 假设不是这样子，因为k次加在那些b[i] = 1的上面时，假设它们一点没有浪费，都添加到了右边，
9. 此时median = x, 选择的是a[i], 且a[i]不是最大值
10. 我们考虑把a[i]最大值换出来。这时有3种情况
11. a. 最大值也在左边（那把它换出来肯定是更好的选择）
12. b. 最大值在右边，但是a[i]也在右边（换出来更好）
13. c. 最大值在右边，但是a[i]在左边（也就是说把最大值换出来后， median要向左移动一位）
14. 那不妨就假设换进去的a[i]就是新的median，那么这么做的收益 = x + a[i] vs ax + a[i]
15. 只要 x <= ax成立，就可以做这个事情
16. 