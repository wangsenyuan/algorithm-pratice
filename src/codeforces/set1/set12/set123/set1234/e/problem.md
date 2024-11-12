Let's define 𝑝𝑖(𝑛)
 as the following permutation: [𝑖,1,2,…,𝑖−1,𝑖+1,…,𝑛]
. This means that the 𝑖
-th permutation is almost identity (i.e. which maps every element to itself) permutation but the element 𝑖
 is on the first position. Examples:

𝑝1(4)=[1,2,3,4]
;
𝑝2(4)=[2,1,3,4]
;
𝑝3(4)=[3,1,2,4]
;
𝑝4(4)=[4,1,2,3]
.
You are given an array 𝑥1,𝑥2,…,𝑥𝑚
 (1≤𝑥𝑖≤𝑛
).

Let 𝑝𝑜𝑠(𝑝,𝑣𝑎𝑙)
 be the position of the element 𝑣𝑎𝑙
 in 𝑝
. So, 𝑝𝑜𝑠(𝑝1(4),3)=3,𝑝𝑜𝑠(𝑝2(4),2)=1,𝑝𝑜𝑠(𝑝4(4),4)=1
.

Let's define a function 𝑓(𝑝)=∑𝑖=1𝑚−1|𝑝𝑜𝑠(𝑝,𝑥𝑖)−𝑝𝑜𝑠(𝑝,𝑥𝑖+1)|
, where |𝑣𝑎𝑙|
 is the absolute value of 𝑣𝑎𝑙
. This function means the sum of distances between adjacent elements of 𝑥
 in 𝑝
.

Your task is to calculate 𝑓(𝑝1(𝑛)),𝑓(𝑝2(𝑛)),…,𝑓(𝑝𝑛(𝑛))
.


### ideas
1. 脑子要冒烟了。
2. pos(p, xi) = xi在排列p中的位置，
3. 比如排列p(2, 5) = 2,1,3,4,5, 
4. pos(p, 3) = 3, pos(p, 2) = 1
5. 在给定排列p的情况下，如何计算 f(p) 呢？
6. x= [2 1 5 3 5], p1 = 1, 2, 3, 4, 5
7. f(p1) = abs(2 - 1) + abs(1 - 5) + abs(5 - 3) + abs(3 - 5)
8. p2 = 2, 1, 3, 4, 5
9. f(p2) = abs(1 - 2) + abs(2 - 5) + abs(5 - 3) + abs(3 - 5)
10. p3 = 3, 1, 2, 4, 5
11. f(p3) = abs(3 - 2) + abs(2 - 5) + abs(5 - 1) + abs(1 - 5)
12. 当从pi 变化到 pi+1时的变化是什么？
13. i+1移动了第一位，同时i移动了i+1的位置
14. 所以，原来用到i+1的，全部要替换成1，用到了i的位置，全部替换为i+1（从2开始比较好）
15. 其他的都不需要变
16. 这样子需要调整的，就只有3个数字
17. 不对，不仅仅是3个数字。
18. pos[x[i]], pos[x[i+1]]的差，大部分时间是不变的
19. 只有当pos[x[i]] 移动到1的时候，或者 pos[x[i+1]]移动到1的时候
20. 或者把某个数移动到他们中间的时候（比如2，移动到了1，3中间的时候）
21. 或者他们中间的某个数移动到头部的时候，比如把1，4中间的2移动到头部的时候
22. 那是不是意味着可以用计算，每一对它在哪些区间内贡献的值是多少？
23. 3 = pos[x[0]], 5 =  pos[x[1]] 
24. 在1,2移动到头部时，对它们没有影响，所以1,2收到的贡献时2
25. 当3移动到头部时，贡献是 5 - 1 = 4
26. 当5移动到头部时，贡献是 3 + 1 - 1 = 3  （5， 1， 2， 3， 4，6）
27. 当4移动到头部时，[4, 1, 2, 3, 5] 贡献 = 1
28. range update, point get