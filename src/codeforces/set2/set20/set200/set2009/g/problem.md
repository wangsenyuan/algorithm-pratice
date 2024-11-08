For an arbitrary array 𝑏
, Yunli can perform the following operation any number of times:

Select an index 𝑖
. Set 𝑏𝑖=𝑥
 where 𝑥
 is any integer she desires (𝑥
 is not limited to the interval [1,𝑛]
).
Denote 𝑓(𝑏)
 as the minimum number of operations she needs to perform until there exists a consecutive subarray∗
 of length at least 𝑘
 in 𝑏
.

Yunli is given an array 𝑎
 of size 𝑛
 and asks you 𝑞
 queries. In each query, you must output ∑𝑟𝑗=𝑙+𝑘−1𝑓([𝑎𝑙,𝑎𝑙+1,…,𝑎𝑗])
.

### ideas
1. 在给定的序列中， f(arr) 怎么快速计算呢？
2. 如果 arr[i] = arr[i-1] + 1
3. 那么有 arr[i] -  i = arr[i-1] - (i-1)
4. 对于每个位置i (i >= k), 计算f(i) = 在区间(i-k + 1, i)中的多数值
5. 不考虑l, f(i) = min(f(i), f(i-1)) (这个不大对)
6. 因为对于r来讲， 从哪个l出发，是有关系的
7. L[i]表示f(i)取值的左下标？
8. 如果L[r] >= l, 那么f(r) * (r - L[r] + k + 1)就可以添加到答案中
9. 然后更新r到L[r] + k - 1
10. 如果L[r] < l, 那么有问题了
11. L[r][0]表示一个k单位的最小值, L[r][1]表示2个k单位的距离
12. fp[r][0] 表示一个k单位长度的结果， f(r)[i]是 1<<i 个k单位的结果
13. f(r)(i) = f(r)(i-1) + f(r - k * (1 << (i-1))(i-1) 
14. 第2部分没问题， 但是第一部分有问题（因为对于r来说, 前面(i-1)个单位的，最小值，也可以被后面部分用到
15. 所以，必须知道各个单位内的最小值
16. 