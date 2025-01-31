Below, a permutation of 
(1,2,…,N) is simply called a permutation.

For two permutations 
A=(A 
1
​
 ,A 
2
​
 ,…,A 
N
​
 ),B=(B 
1
​
 ,B 
2
​
 ,…,B 
N
​
 ), let us define their similarity as the number of integers 
i between 
1 and 
N−1 such that:

(A 
i+1
​
 −A 
i
​
 )(B 
i+1
​
 −B 
i
​
 )>0.
Find the number, modulo a prime number 
P, of pairs of permutations 
(A,B) whose similarity is 
K.


### ideas
1. (a[i] - a[i+1]) * (b[i] - b[i+1]) > 0
2. 表示的是两个序列的变化趋势是一样的
3. 考虑最后两个数 (n - 1, n)
4. 假设把最后一个数去掉后，那么剩下的n-1个数，仍然是可以随意编排的（因为只用考虑相对顺序）
5. 所以，如果最后的两个数分别是(n, n-1) 那么如果 a是 (n, n - 1), 那么 b 就不能是 (n-1, n)
6. f(n) = f(n-1) * (n * n - 交叉的那些) 似乎不大对，因为交叉的那些，会收到(n-1)以什么结尾的限制
7. f(n)(i)(j) 表示以(i, j)结尾的情况
8. f(n)(i)(j) += f(n-1)(x)(y) 如果 (x - i) * (y - j) > 0 
9. 但是问题是这里是 pow(100, 5) = 1e10， TLE
10. 注意到 f(i)(j) += f(x)(y) 这个有点像矩形面积求和
11. 可以到1e6
12. 上面的分析，漏掉了一个很重要的东西，就是这个序列中，有k个相似度
13. 甚至k可以等于0
14. 首先k个相似度，那么就有k+1个数字，它们在a/b中的位置所组成的形状是一样的
15. 那么这样的数字的选择有 C(n, k) * C(n, k)中，
16. f(n)(k)表示有n个序列组成的，生成k个相似度
17. f(n)(k) = f(n-1)(k-1) 如果新的放在最后一个位置（它肯定会产生一个新的相似度）
18. 但是如果a里面的n放在了位置i， b里面的n放在了另外一个位置j， 这个对相似度的影响是怎样的？
19. 这样子还是有点乱。没想法
20. 对于两个长度都为n的序列，如果这个序列的相似度正好是n-1， 那么和其中的具体的元素是无关的，只和长度有关系
21. 先考虑这个简单的问题先。 
22. 考虑峰谷和峰顶，显然n在峰顶，1在峰谷
23. 如果一个峰顶元素，那么看n的位置在哪里。这样的有pow(2, n - 1) * pow(2, n - 1) = pow(2, n)种
24. 