K1o0n gave you an array 𝑎
 of length 𝑛
, consisting of numbers 1,2,…,𝑛
. Accept it? Of course! But what to do with it? Of course, calculate MEOW(𝑎)
.

Let MEX(𝑆,𝑘)
 be the 𝑘
-th positive (strictly greater than zero) integer in ascending order that is not present in the set 𝑆
. Denote MEOW(𝑎)
 as the sum of MEX(𝑏,|𝑏|+1)
, over all distinct subsets 𝑏
 of the array 𝑎
.

Examples of MEX(𝑆,𝑘)
 values for sets:

MEX({3,2},1)=1
, because 1
 is the first positive integer not present in the set;
MEX({4,2,1},2)=5
, because the first two positive integers not present in the set are 3
 and 5
;
MEX({},4)=4
, because there are no numbers in the empty set, so the first 4
 positive integers not present in it are 1,2,3,4
.

### ideas
1. 计算的是1, 2, ... n的数组a的所有subset的 mex(b, sz(b) + 1)
2. 比如 mex({}, 1) = 1
3.     mex({2, 3}, 3) = 5 (1, 2, 3, 4, 5)
4.     mex({1, 2}, 3) = 5
5.     mex({4, 5}, 3) = 3
6.     mex({1, 4}, 3) = 5
7.     mex({1, 5}, 3) = 4
8.     mex({4, 5, 6}, 4) = 7 (1, 2, 3, ... 7)
9.     mex({1, 5, 6}, 4) = 7
10.    mex({1, 2, 6}, 4) = 7 (3, 4, 5, 7)
11.    mex({1, 2, 7}, 4) = 6
12.    mex({1, 6, 7}, 4) = 5 (2, 3, 4, 5)
13. mex({set}, sz + 1) 最小也是sz+1, sz + 2, sz + ... 2 * sz + 1
14. 计算贡献？
15. 当 f(set, sz + 1) = sz + 1 的时候， 1....sz + 1不能出现，其他的 sz + 2 .... n 可以随便出现 nCr(n - (sz + 1), sz)
16. 当 f = sz + 2 时只要set中包含了1....sz + 1 中的任何一个，但是不包括sz+2时，nCr(sz + 1, 1) * (n - (sz + 2), sz - 1)
17. f = sz + i 时， nCr(sz + i - 1, i - 1) * nCr(n - (sz + i), sz - (i - 1)) 个
18. 当set的大小时 1, 2, 3, 4, .... n时， j = sz + 1, .... 2 * sz + 1 时
19. 这个是n * n 的