Suppose you have two points 𝑝=(𝑥𝑝,𝑦𝑝)
 and 𝑞=(𝑥𝑞,𝑦𝑞)
. Let's denote the Manhattan distance between them as 𝑑(𝑝,𝑞)=|𝑥𝑝−𝑥𝑞|+|𝑦𝑝−𝑦𝑞|
.

Let's say that three points 𝑝
, 𝑞
, 𝑟
 form a bad triple if 𝑑(𝑝,𝑟)=𝑑(𝑝,𝑞)+𝑑(𝑞,𝑟)
.

Let's say that an array 𝑏1,𝑏2,…,𝑏𝑚
 is good if it is impossible to choose three distinct indices 𝑖
, 𝑗
, 𝑘
 such that the points (𝑏𝑖,𝑖)
, (𝑏𝑗,𝑗)
 and (𝑏𝑘,𝑘)
 form a bad triple.

You are given an array 𝑎1,𝑎2,…,𝑎𝑛
. Calculate the number of good subarrays of 𝑎
. A subarray of the array 𝑎
 is the array 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
 for some 1≤𝑙≤𝑟≤𝑛
.

Note that, according to the definition, subarrays of length 1
 and 2
 are good.

### ideas
1. bad ，就是哈对于三个点来说，一个点位于另外两个点组成的长方形中
2. b1, b2, ... bm
3. is good if no bad 三元组 (i < j < k)
4. 那么就是 b[j] < min(b[i], b[k]) 或者 b[j] > max(b[i], b[k])
5. 那么考虑 1, m, 所有的其他的b[i]都在1...m的上面或者下面
6. 这里不妨设 b[1] < b[m]
7. b[2] > b[m], b[3] > b[2] (但是这里就有问题了)
8. 那么1...3 所形成的区域，就包括了2了。 所以 b[3] < b[1]， 但是这时，貌似不能4 != m 了吧。
9. 如果存在， b[4] < b[3], 会造成3在1...4的区域内，要么b[4] > b[2], 造成2在1，4的区域内
10. 也就是说m, 最多是4
11. 如果m = 2, 很好结果， 如果 m = 3, 也比较好解决，就是找到前面比a[r]大的，在r前面有多少 < a[r]的即可
12. m = 4, 这个要怎么计算呢？
13. a, b, c, d 那么d 必须处在(b, c)的中间， 且a也必须处在(b, c)的中间
14. 但是显然是不能迭代(b, c)的，太慢了
15. subarray是连续的
16. 那就简单了
17. 
    