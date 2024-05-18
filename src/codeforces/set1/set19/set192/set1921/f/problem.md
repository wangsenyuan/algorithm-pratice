You are given an array 𝑎
 of 𝑛
 numbers. There are also 𝑞
 queries of the form 𝑠,𝑑,𝑘
.

For each query 𝑞
, find the sum of elements 𝑎𝑠+𝑎𝑠+𝑑⋅2+⋯+𝑎𝑠+𝑑⋅(𝑘−1)⋅𝑘
. In other words, for each query, it is necessary to find the sum of 𝑘
 elements of the array with indices starting from the 𝑠
-th, taking steps of size 𝑑
, multiplying it by the serial number of the element in the resulting sequence.

Input
Each test consists of several testcases. The first line contains one integer 𝑡
 (1≤𝑡≤104
) — the number of testcases. Next lines contain descriptions of testcases.

The first line of each testcase contains two numbers 𝑛,𝑞
 (1≤𝑛≤105,1≤𝑞≤2⋅105
) — the number of elements in the array 𝑎
 and the number of queries.

The second line contains 𝑛
 integers 𝑎1,...𝑎𝑛
 (−108≤𝑎1,...,𝑎𝑛≤108
) — elements of the array 𝑎
.

The next 𝑞
 lines each contain three integers 𝑠
, 𝑑
, and 𝑘
 (1≤𝑠,𝑑,𝑘≤𝑛
, 𝑠+𝑑⋅(𝑘−1)≤𝑛
 ).

It is guaranteed that the sum of 𝑛
 over all testcases does not exceed 105
, and that the sum of 𝑞
 over all testcases does not exceed 2⋅105
.

### ideas
1. for s, d, k
2. let f(s, d, k) = a[s] + a[s + d] * 2 + ... + a[s + (k - 1) * d] * k
3. f(s, d, k) = sum(a[s], a[s+d], ... a[s + (k - 1) * d]) + f(s + d, d, k - 1)
4. F(s, d) = a[s] + a[s + d] * 2 + a[s + 2 * d] * 3... + a[s + (k - 1) * d] * k + a[s + k * d] * (k + 1) ...
5. 这个变形没啥用
6. 还是得把k给提出来
7. F(s) = 等于从s开始的，s的倍数的sum
8. s, s + d, s + d + d, s + d + d + d
9. 这里有一个难点是，如果从后往前的算，那么k的范围没法控制
10. 那就从前往后处理？
11. F(i) 表示到i计算好的某种结果
12. 如果有一个query (s, d, k) 满足 s + (k - 1) * d = i
13. 如果能够从F[i]中把结果推导出来
14. 那F[i]需要保存什么信息呢？
15. F[i] 和 F[i-d]有什么关系？
16. 假设F[i-d]中保存了某个数的j次, 那么 F[i][j+1] = F[i-d][j] + (j+1) * a[i]
17. 不行～
18. 还是不够creative
19. 1， 2， 3， 4，...., n
20. 有点戏的貌似还是第3点
21. 就brute force处理，按照 (s, d) 把k排序？
22. 这样不行， d越小的时候， k就也可能很大，假设 都是 (s, 1), 那么这个潜在的影响就是 q * n
23. 主要是这个k的影响，不好处理
24. 如果没有k
25. f(s, d) = a[s] + a[s + d] + a[s + 2 * d] + a[s + 3 * d] .. a[s + i * d] ...
26. 没办法～～～～
27. 有个点，就是对于i来说 (i - s) 是可以被d整除
28. 那么这样子，d可以是 (i - s) (<= n) 的一个因子
29. 在d给定的情况下 k一定程度上时可以推导出来的
30. 但问题在于， 迭代i, s的复杂性，还是 n * n
31. 对于query (s, d, k) 来说， 可以把它分布在这样的位置上
32. s, s + d, s + 2 * d, s + 3 * d ... 
33. 这样子的最多时 sqrt(n)个
34. 反过来每个位置也最多这么多个查询，然后把查询给加上去就可以了
35. 当d足够大的时候，可以这样做，或者k足够小的时候
36. 当d小且k比较大的时候的时候，需要另外一种做法
37. 这样的d时比较少的，不超过sqrt(n)个

### solution

The key idea is that we know how to calculate the sum (𝑖−𝑙+1)⋅𝑎𝑖
 for 𝑙≤𝑖≤𝑟
 fast – we need to calculate all prefix sums 𝑖⋅𝑎𝑖
 and 𝑎𝑖
 for 1≤𝑖≤𝑘
, then take the difference between the 𝑟
-th and (𝑙−1)
-th of 𝑖⋅𝑎𝑖
 and subtract the difference between the 𝑟
-th and (𝑙−1)
-th multiplied by 𝑙−1
. This way queries with step 1
 will be processed in 𝑂(𝑛+𝑞)
 time, where 𝑞
 is the total amount of queries with step 1.

But this idea can be generalized to the following: we can precalculate all the prefix sums and all the prefix sums with multiplication by index for every 𝑑0≤𝑑
 in 𝑂(𝑛⋅𝑑)
 time, and then process all queries with step 𝑑0≤𝑑
 in 𝑂(1)
 time.

However, for all other queries we can process a single query in 𝑂(𝑛/𝑑)
 time, because the difference between consecutive elements in the resulting sequence is greater than 𝑑
.

Combining these two ideas, we get a solution with a time complexity 𝑂(𝑛⋅𝑑+𝑞⋅𝑛/𝑑)
. Setting 𝑑=𝑞√
, we get a solution with a time complexity 𝑂(𝑛𝑞√)
. The model solution fixes the value of 𝑑=322
, which is equal to 𝑀𝐴𝑋‾‾‾‾‾‾√
.

Interestingly, this solution can be generalized to calculate the sums (𝑖+1)2⋅𝑎𝑠+𝑑⋅𝑖
.