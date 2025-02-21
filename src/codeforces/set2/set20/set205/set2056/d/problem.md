An array 𝑏
 of 𝑚
 integers is called good if, when it is sorted, 𝑏⌊𝑚+12⌋=𝑏⌈𝑚+12⌉
. In other words, 𝑏
 is good if both of its medians are equal. In particular, ⌊𝑚+12⌋=⌈𝑚+12⌉
 when 𝑚
 is odd, so 𝑏
 is guaranteed to be good if it has an odd length.

You are given an array 𝑎
 of 𝑛
 integers. Calculate the number of good subarrays∗
 in 𝑎
.

∗
An array 𝑥
 is a subarray of an array 𝑦
 if 𝑥
 can be obtained from 𝑦
 by the deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤104
). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤105
) — the length of the array.

The second line of each test case contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤10
) — the given array.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 105
.

### ideas
1. 突破口难道在 a[i] <= 10
2. 当x是median时，为cnt[y < x] < n / 2 的最小值x
3. 假设可以迭代l...r
4. cnt[r][y < x] - cnt[l][y<x] < (r - l)/ 2
5. (r - l) 为偶数
6. (r - l) / 2 = r / 2 - l / 2
7. (4 - 2) / 2 = 
8. (5 - 3) / 2 = 2 - 1 = 1
9. cnt[r][y<x] - r / 2 < cnt[l][y < x] - l / 2
10. 那就简单了
11. 