For a given sequence of 𝑛
integers 𝑎
, a triple (𝑖,𝑗,𝑘)
is called magic if:

1≤𝑖,𝑗,𝑘≤𝑛
.
𝑖
, 𝑗
, 𝑘
are pairwise distinct.
there exists a positive integer 𝑏
such that 𝑎𝑖⋅𝑏=𝑎𝑗
and 𝑎𝑗⋅𝑏=𝑎𝑘
.
Kolya received a sequence of integers 𝑎
as a gift and now wants to count the number of magic triples for it. Help him with this task!

Note that there are no constraints on the order of integers 𝑖
, 𝑗
and 𝑘
.

### thoughts

1. b > 1, a[i] < a[j] < a[k]
2. b = 1, a[i] = a[j] = a[k]
3. 排序
4. 分成两种情况，相等和不相等
5. 相等的比较好处理 C(m, 3)
6. 不相等的情况， 可以计算 prev[x] 表示当前数字钱x的个数
7. suf[x]表示当前数字后x的个数
8. 迭代d from 2, 3,4.... sqrt(max of a)
9. 还要再分析一下