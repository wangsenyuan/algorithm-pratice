You are given an array 𝑎
 consisting of 𝑛
 integers.

Let the function 𝑓(𝑏)
 return the minimum number of operations needed to make an array 𝑏
 a palindrome. The operations you can make are:

choose two adjacent elements 𝑏𝑖
 and 𝑏𝑖+1
, remove them, and replace them with a single element equal to (𝑏𝑖+𝑏𝑖+1)
;
or choose an element 𝑏𝑖>1
, remove it, and replace it with two positive integers 𝑥
 and 𝑦
 (𝑥>0
 and 𝑦>0
) such that 𝑥+𝑦=𝑏𝑖
.
For example, from an array 𝑏=[2,1,3]
, you can obtain the following arrays in one operation: [1,1,1,3]
, [2,1,1,2]
, [3,3]
, [2,4]
, or [2,1,2,1]
.

Calculate (∑1≤𝑙≤𝑟≤𝑛𝑓(𝑎[𝑙..𝑟]))
, where 𝑎[𝑙..𝑟]
 is the subarray of 𝑎
 from index 𝑙
 to index 𝑟
, inclusive. In other words, find the sum of the values of the function 𝑓
 for all subarrays of the array 𝑎
.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤2000
).

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤105
).

Additional constraint on the input: the sum of 𝑛
 over all test cases does not exceed 2000
.

### ideas
1. 单个数字 f(arr) = 0
2. f(arr, 2) = 0 if already good, else 1 (一次操作后，比如分解大的那个数字, 或者把两个数合并成一个)
3. f(arr, l...r) = f(arr, n - 2) + 0 (如果 arr[l] = arr[r]) or ? 
4.               or f(arr, l, r - 1) + 2? 假设a[l], a[l+1], .... a[r-1], a[r]
5.     a[l].... a[r-1] 变成回文的最少次数是x, 现在，希望把 a[l].... a[r] 变成回文
6. a, b, c, d  如果  a + b == c + d, 那么需要额外的两次操作 a[l+2][r-2] + 2
7. a[l].... a[r-1] 变成回文的话，感觉和差有关系？
8. 如果 f(l, r) = f(l + 1, r - 1) + 0 (如果 a[l] = a[r])
9.             = 那和(l...r-1)有什么关系？如果能从 a[l] 里面分解出一个 a[r], 且剩余的
10.  a + b = c + d, good， else a + b > c + d, 那么合并 + 分解
11.  但感觉这样子不是最优策略 
12.  a, b, c 如果 a != c, 如果 a + b = c or a = b + c => 1, else 2 (从大的数字中分解一个出来，然后和b融合)
13. 一个感觉就是f(l,r) 貌似和 f(l+1, r), f(l, r - 1), f(l+1, r - 1)的关系不大
14. f(l, r) = f(l + 1, r - 1) + 0 if a[l] == a[r] else 把大的那个数分解出来的，弄到最中间的距离？


### solution
合并操作，每次可以减少一个数字。

分解操作，比如最左边是 
2
2，最右边是 
5
5，我们分解大的那个数，得到 
3
3 和 
2
2。这等价于去掉最左边的数，并修改最右边的数。所以每次操作也可以减少一个数字。

所以两种操作效果是相同的，不妨只考虑合并操作。

长为 
m
m 的子数组 
b
b，操作 
m
−
1
m−1 次可以变成一个数，此时一定是回文的。

能否减少操作次数呢？

比如 
b
=
[
2
,
2
,
1
,
3
]
b=[2,2,1,3]，左右都可以合并成 
4
4，最后得到 
[
4
,
4
]
[4,4]，这样只需操作 
2
2 次而不是 
3
3 次，减少了 
1
1 次操作。

又比如 
b
=
[
2
,
2
,
9
,
1
,
3
]
b=[2,2,9,1,3]，左右都可以合并成 
4
4，最后得到 
[
4
,
9
,
4
]
[4,9,4]，这样只需操作 
2
2 次而不是 
4
4 次，减少了 
2
2 次操作。

一般地，如果前缀 
[
0
,
i
]
[0,i] 与后缀 
[
j
,
m
−
1
]
[j,m−1] 的元素和相同，那么可以减少 
1
1 次操作。注意前后缀可以相交，在 
b
=
[
2
,
2
,
9
,
1
,
3
]
b=[2,2,9,1,3] 这个例子中，我们有 
2
+
2
=
1
+
3
2+2=1+3，同时还有 
2
+
2
+
9
=
9
+
1
+
3
2+2+9=9+1+3，所以可以减少 
2
2 次操作。

考虑贡献法，如果有两个非空子数组 
[
i
1
,
j
1
]
,
[
i
2
,
j
2
]
[i 
1
​
 ,j 
1
​
 ],[i 
2
​
 ,j 
2
​
 ] 的元素和相同，那么对于下标从 
i
1
i 
1
​
  到 
j
2
j 
2
​
  的子数组 
b
b 来说，可以减少 
1
1 次操作，对答案的贡献就是 
−
1
−1。

所以只需统计非空子数组和的个数 
cnt
s
cnt 
s
​
 。遍历到子数组 
[
i
,
j
]
[i,j] 时，设其元素和为 
s
s，那么它与之前统计过的 
cnt
s
cnt 
s
​
  个子数组和相同，对答案的贡献就是 
j
−
i
−
cnt
s
j−i−cnt 
s
​
 。其中 
j
−
i
j−i 是不考虑相同前后缀时，长为 
j
−
i
+
1
j−i+1 的子数组所需的操作次数。