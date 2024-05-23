An array 𝑎
 of length 𝑚
 is considered good if there exists an integer array 𝑏
 of length 𝑚
 such that the following conditions hold:

∑𝑖=1𝑚𝑎𝑖=∑𝑖=1𝑚𝑏𝑖
;
𝑎𝑖≠𝑏𝑖
 for every index 𝑖
 from 1
 to 𝑚
;
𝑏𝑖>0
 for every index 𝑖
 from 1
 to 𝑚
.
You are given an array 𝑐
 of length 𝑛
. Each element of this array is greater than 0
.

You have to answer 𝑞
 queries. During the 𝑖
-th query, you have to determine whether the subarray 𝑐𝑙𝑖,𝑐𝑙𝑖+1,…,𝑐𝑟𝑖
 is good.


### ideas
1. 如果a都等于1, 那显然没有答案
2. 如果a = []int{1, 2, 1}
3. 似乎也没有答案
4. 如果 a = []int{1, 2, 2, 1} 有答案
5. 首先我们让b都等于1，剩下的数字 = sum(a) - n
6. 如果这个数 >= n, 似乎始终是有答案的
7. 换个角度，考虑sum(a) - n, 如果这个数 < n, (假设m), 那么意味着至少有 n - m个0，如果a中0的个数 > n / 2
8. 那么就没有答案