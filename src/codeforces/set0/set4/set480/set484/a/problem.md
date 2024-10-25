Let us denote by the number of units in the binary notation of a non-negative integer x .

You are given several queries, each in the form of a pair of integers l and r . For each query, find x such that l  ≤  x  ≤  r , and is maximal. If there are several such numbers, then you need to find the smallest of them.

Input data
The first line contains an integer n  — the number of requests ( 1 ≤  n  ≤ 10000 ).

The next n lines contain two numbers l i ,  r i  — the parameters for the next request ( 0 ≤  l i  ≤  r i  ≤ 10 18 ).



### ideas
1. 考虑一种特殊的形式， r = 1111 那么 x = r
2. 如果不是，且hi(l) < hi(r), 那么 x = 1111 (比r的位数少1)
3. 如果 hi(l) = hi(r)呢？那么就是找到第一个i, l(i) < r(i), 然后这位后面的都是1，即可