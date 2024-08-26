## Problem Analysis and Organization

**Problem:**
* Given two sequences `A` and `B` of length `N`.
* Find the number of pairs `(l, r)` satisfying:
  * `1 ≤ l ≤ r ≤ N`
  * `min(A[l], A[l+1], ..., A[r]) + (B[l] + B[l+1] + ... + B[r]) ≤ S`


### ideas
1. 对于l...r区间，如果a[i]是最小值，且 sum(b[l...r]) + a[i] <= S 
2. 固定r的时候， pref[r+1] - pref[l] + a[i] <= S
3. -pref[l] + a[i] <= S - pref[r+1] = X
4. 也就是左边有多少(l, i)的组合 <= X (且这里 l <= i)
5. 在l也确定的时候（假设），那么a[i]是最大的最小值满足 -pref[l] + a[i] <= X 的下标
6. 这时候的贡献是 i - l + 1
7. 所以对于给定i来说，l是最小的-pref[l] + a[i] <= X的下标
8. 在区间[l...r]中间， a[i]组成的是一个递增序列，且最后一个肯定是a[r]
9. 对于a[r]的l是可以计算出来的 r - l[r] + 1个
10. i - l[i] + 1, j - l[j] + 1, ... r - l[r] + 1
11. 关键就在与 l[i], ... l[j]是不确定的 但是有个点，就是从r到r+1的时候，很多是不变的
12. （有一部分它的l[i]会比当前的l小）