You are given a one-based array consisting of 𝑁
 integers: 𝐴1,𝐴2,⋯,𝐴𝑁
. Initially, the value of each element is set to 0
.

There are 𝑀
 operations (numbered from 1
 to 𝑀
). Operation 𝑖
 is represented by ⟨𝐿𝑖,𝑅𝑖,𝑋𝑖⟩
. If operation 𝑖
 is executed, all elements 𝐴𝑗
 for 𝐿𝑖≤𝑗≤𝑅𝑖
 will be increased by 𝑋𝑖
.

You have to answer 𝑄
 independent queries. Each query is represented by ⟨𝐾,𝑆,𝑇⟩
 which represents the following task. Choose a range [𝑙,𝑟]
 satisfying 𝑆≤𝑙≤𝑟≤𝑇
, and execute operations 𝑙,𝑙+1,…,𝑟
. The answer to the query is the maximum value of 𝐴𝐾
 after the operations are executed among all possible choices of 𝑙
 and 𝑟
.

Input
The first line consists of two integers 𝑁
 𝑀
 (1≤𝑁,𝑀≤100000
).

Each of the next 𝑀
 lines consists of three integers 𝐿𝑖
 𝑅𝑖
 𝑋𝑖
 (1≤𝐿𝑖≤𝑅𝑖≤𝑁;−100000≤𝑋𝑖≤100000
).

The following line consists of an integer 𝑄
 (1≤𝑄≤100000
).

Each of the next 𝑄
 lines consists of three integers 𝐾
 𝑆
 𝑇
 (1≤𝐾≤𝑁;1≤𝑆≤𝑇≤𝑀
).

Output
For each query, output in a single line, an integer which represent the answer of the query.

### ideas
1. 感觉很难呐
2. 单独的操作1，可以用range update
3. 把query 按照 T升序排
4. 然后把operation和query一起处理，在处理完operation i后，处理T = i的query
5. 构建一个segment forest
6. 每个segment tree表示操作i后的状态
7. 每个node维护两个信息，它在操作i后的value，在它前面最大的值，以及对应的操作是什么
8. 上面的不大行
9. 这里的关键是要知道k在操作区间[l...r]的最大值是什么
10. 二维的segment tree？
11. 但是如何更新呢？这里更新的时候，是把一个区间的所有的值都更新掉
12. 如果每个元素都更新， 肯定会成 n * m 的复杂性
13. 假设不去更新，要怎么样知道这个信息呢？
14. 在每个节点上，把更新记录下来。但不真的更新
15. 在查询的时候，把更新不断的push下去？但是只要push下去了，在最终的节点上，就发生了更新了
16. 所以，还是不大对
17. 对于区间[s...t],如果最大值出现在区间[l...r]， 那么可以肯定的一点是 
18. [s...l-1] < 0
19. [r+1....t] < 0 否则的话，只要把前缀或者后缀加入进去，就能获得更大的值
20. 所以，从上面的分析可以得出这样一个结论，就是要找到最小的前缀[s...l)
21. 和最小的后缀 (r....t]
22. 剩下的部分就是最大值
23. 不对～～～～
24. 还是没有找到题眼～～～
25. 假设有个数据结构，维护了足够的信息。在查询k的时候，如果能够确定k的操作和当前节点的操作是一致的
26. 那么就不需要一路访问到k; 
27. 比如操作 [1...5], [2...5], [2...4], 那么在查询k=3的时候，那么只要到达包含区间 [2...4]的节点时
28. 就ok了
29. 如果是要查询k=5， 那只需要到区间 [3...5] (因为它已经包含了操作1，2，且不包含操作3)
30. 假设发现node的左子节点为空，且需要go left， 那么是不是查询当前节点即可呢？
31. 好像是的呐
32. 但是如何维护这个数据结构呢？另外一个segment tree？
33. 在push的时候，创建出来？
34. 搞不定， 不搞了～

### solution

First, let’s discuss how to solve this problem if all of the K are same.
We can construct an auxiliary array B with size M which contains operations that can affect AK.
If K is not between Li and Ri
, then the value of array Bi
is 0. Otherwise, the value of array Bi
is
Xi
. By using array B, a query can then be answered by returning “the maximum subarray sum in
a given range” from this array. This query can be answered using a segment tree data structure.
Now, to solve the problem for any K, we can sort the queries based on K. Initially, we need to
construct a segment tree for the auxiliary array when K = 1 and answer all queries with K =
1. When transitioning to next K, we can avoid rebuilding the segment tree by only updating the
affected indices. If operation i has Ri = K, then we need to update Bi
to 0. Another case is if
operation i has Li = K + 1, then we need to update Bi
to Xi
. Each operation i will only update the
auxiliary array at most twice.
The total time complexity for this problem is O((Q + M) log M).