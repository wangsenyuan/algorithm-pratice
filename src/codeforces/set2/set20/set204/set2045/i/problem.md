You are given an array of 𝑁
 integers: [𝐴1,𝐴2,…,𝐴𝑁]
.

A subsequence can be derived from an array by removing zero or more elements without changing the order of the remaining elements. For example, [2,1,2]
, [3,3]
, [1]
, and [3,2,1,3,2]
 are subsequences of array [3,2,1,3,2]
, while [1,2,3]
 is not a subsequence of array [3,2,1,3,2]
.

A subsequence is microwavable if the subsequence consists of at most two distinct values and each element differs from its adjacent elements. For example, [2,1,2]
, [3,2,3,2]
, and [1]
 are microwavable, while [3,3]
 and [3,2,1,3,2]
 are not microwavable.

Denote a function 𝑓(𝑥,𝑦)
 as the length of the longest microwavable subsequence of array 𝐴
 such that each element within the subsequence is either 𝑥
 or 𝑦
. Find the sum of 𝑓(𝑥,𝑦)
 for all 1≤𝑥<𝑦≤𝑀
.

### ideas
1. f(x, y) 1 <= x < y <= m 的microwable序列的最大长度
2. 如果从x,y出发感觉好难搞。能不能从长度出发，计算有多少种(x,y)组合，能够得到长度i
3. g(i) * i 
4. 怎么计算g(i)， 也没有头绪
5. 从1.。。m迭代y，然后看和y能形成最长的序列的长度
6. 这个序列的长度，不会超过 2 * freq[y] + 1
7. 但是肯定也不能一个个的试验吧， 比如序列 1, 1, 3, 2, 3, 1
8. f(x, y) 肯定以x/y开始，并以x/y结束
9. 所以，x/y第一次出现和最后一次出现比较重要
10. x第一次出现，并且x最后一次出现
11. 这时候的y，只能是x中间的数字（如果出现在外围，就不符合x在头/尾的情况）
12. 想不出来，算了