Jayden has an array 𝑎
 which is initially empty. There are 𝑛
 operations of two types he must perform in the given order.

Jayden appends an integer 𝑥
 (1≤𝑥≤𝑛
) to the end of array 𝑎
.
Jayden appends 𝑥
 copies of array 𝑎
 to the end of array 𝑎
. In other words, array 𝑎
 becomes [𝑎,𝑎,…,𝑎⏟𝑥]
. It is guaranteed that he has done at least one operation of the first type before this.
Jayden has 𝑞
 queries. For each query, you must tell him the 𝑘
-th element of array 𝑎
. The elements of the array are numbered from 1
.

Input
Each test consists of multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤5000
) — the number of test cases. The description of the test cases follows.

The first line of each test case contains two integers 𝑛
 and 𝑞
 (1≤𝑛,𝑞≤105
) — the number of operations and the number of queries.

The following 𝑛
 lines describe the operations. Each line contains two integers 𝑏
 and 𝑥
 (𝑏∈{1,2}
), where 𝑏
 denotes the type of operation. If 𝑏=1
, then 𝑥
 (1≤𝑥≤𝑛
) is the integer Jayden appends to the end of the array. If 𝑏=2
, then 𝑥
 (1≤𝑥≤109
) is the number of copies Jayden appends to the end of the array.

The next line of each test case contains 𝑞
 integers 𝑘1,𝑘2,…,𝑘𝑞
 (1≤𝑘𝑖≤min(1018,𝑐)
), which denote the queries, where 𝑐
 is the size of the array after finishing all 𝑛
 operations.

It is guaranteed that the sum of 𝑛
 and the sum of 𝑞
 over all test cases does not exceed 105
.


### ideas
1. 先观察一个例子
2. 1 （1, 2), (2, 1) 
3. 1, 2, 1, 2
4. 感觉可以倒过来考虑，假设最后的完整数组是知道的
5. 然后要查第k个数
6. 如果k正好等于数组的size， 那它就是最后一个数
7. 这里分两种情况， 如果最后一个数是操作1得到的，那么就可以直接获取到
8. 如果最后一次操作是操作2， 那么 k <= k / x, 然后递归
9. 如果k 〈 size
10. 如果最后一次是append， 那么舍弃掉最后一个继续处理
11. 如果最后一次是copy，那么应该是 k % (size / x)的位置
12. 但是这样的问题是，对于一个k，要处理q次才行
13. 至少有了一个方向
14. 还有一个点，就是最后所有的数字，只可能是从操作1加入的
15. 所以把操作1加入的数字，作为其他数字的root
16. 如果能很快的找到这些root，那么就能解决问题
17. 考虑一个append后面的copy的操作，
18. 假设append后的数组的size = a
19. copy x后的size  = x * a
20. 那么可以确定的一点是，所有a的倍数位置的数，都应该是append的那个数，直到append一个新的数后
21. 假设k就是目前<= 最大的那个数组的size 
22. 这个数组中的数字，似乎满足这样一个规律
23. 假设依次通过操作1加入了数字 x[1], x[2], x[3], ...
24. 且y[1], y[2], y[3], ... y[i]是x[i]加入后的数组的大小
25. 那么从后往前，如果 pos % y[i] = 0, 那么
26. 也就是找到最的，能够整除 pos的下标
27. 但是pos还是太大了，有1e18这么大，开方也有1e9
28. 但似乎可以反过来吗？
29. 就是用当前的pos去更新后面的答案？
30. 这样子是 log(1e18)吗？
31. 试试看吧