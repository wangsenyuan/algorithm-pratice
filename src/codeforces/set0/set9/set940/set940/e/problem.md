Since you are the best Wraith King, Nizhniy Magazin «Mir» at the centre of Vinnytsia is offering you a discount.

You are given an array a of length n and an integer c.

The value of some array b of length k is the sum of its elements except for the smallest. For example, the value of the
array [3, 1, 6, 5, 2] with c = 2 is 3 + 6 + 5 = 14.

Among all possible partitions of a into contiguous subarrays output the smallest possible sum of the values of these
subarrays.

### ideas

1. 考虑一个长位k的arr，它的floor(k/c)个元素被舍弃掉
2. 如果arr是排好序的，如何partition？
3. 如果单个子数组的长度 < c, 那么这个子数组内没有被删除的元素
4. 当子数组的长度 >= c时，才能减少一次
5. 然后最好长度是c的倍数，否则多余的部分就会被浪费掉
6. dp[i]表示到i为止且分后，最小sum
7. dp[i] = dp[j-1] + sum(...) - min(arr[j...i]) when i - j >= c else - 0
8. 这里没有必要去检查 i - j > 2 * c 的部分， 那部分答案已经包含在dp[j]中
9. dp[i] = dp[j-1] + sum[i] - sum[j-1] - 最小值
10. sum[i]是确定的， dp[j-1] - sum[j-1] - ? 的最小值
11. 考虑2个为止 j1, j2, j1 < j2
12. x1是 j1+1...i 中间的最小值
13. x2是 j2+1...i中间的最小值
14. x1 <= x2
15. x2是确定的，但是x1怎么弄上去，出现问题了
16. 因为x1其实是在变化的，假设j1的x1，一开始 = a[j1+1],
17. a[j1+2] >= a[j1+1] 不需要更新，但是如果 a[j1+2] < a[j1+1]那么就需要更新
18. 但是如果 a[j1+3] < a[j1+2] > a[j1+1] 要怎么样更新？
19. 麻了～～
20. 但是这里删除x1而不是x2，肯定不是一个好选择，
21. 更优的选择是放弃x1, 直接删除x2, 可以通过，将 j1+1...j-1分为一组(这里它们不够c, 所以不删除)
22. 再将[j...i]分为一组，从而删除x2
23. 所以上面的 dp[i] = sum[i] + min(dp[j-1] - sum[j-1]) - x2 (j处于 i - 2 * c+1 ... i - c+1中间)
24. 