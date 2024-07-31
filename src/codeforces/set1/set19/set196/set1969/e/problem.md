You are given an integer array 𝑎
 of length 𝑛
. A subarray of 𝑎
 is one of its contiguous subsequences (i. e. an array [𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟]
 for some integers 𝑙
 and 𝑟
 such that 1≤𝑙<𝑟≤𝑛
). Let's call a subarray unique if there is an integer that occurs exactly once in the subarray.

You can perform the following operation any number of times (possibly zero): choose an element of the array and replace it with any integer.

Your task is to calculate the minimum number of aforementioned operation in order for all the subarrays of the array 𝑎
 to be unique.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤3⋅105
).

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤𝑛
).

Additional constraint on the input: the sum of 𝑛
 over all test cases doesn't exceed 3⋅105
.

Output
For each test case, print a single integer — the minimum number of aforementioned operation in order for all the subarrays of the array 𝑎
 to be unique.

### ideas
1. 对于l...r. 如果这区间内存在任意一个数字，只出现了一次，那么这个区间就是unique的
2. 有点不清楚咋处理，贪心？动态规划？找规律？
3. 假设通过某种方式，使得区间[l...r] unique了，花费了x次操作，那么这些操作，同样会影响到其他的区间
4. 还有一个就是，每次操作，始终是可以让这个数字变成unique的，使得所有包含这个位置的区间，unique
5. 假设调整了i, 那么结果 = f(l...i) + change(i) or not + f(i+1, r)
6. 这个是n*n。还需要优化

###
When we replace an element, we can always choose an integer that is not present in the array. So, if we replace the 𝑖
-th element, every subarray containing it becomes unique; and the problem can be reformulated as follows: consider all non-unique subarrays of the array, and calculate the minimum number of elements you have to choose so that, for every non-unique subarray, at least one of its elements is chosen.

We can use the following greedy strategy to do it: go through the array from left to right, maintaining the index 𝑠
 of the last element we replaced. When we consider the 𝑖
-th element of the array, if there is a non-unique subarray [𝑙,𝑟]
 with 𝑙>𝑠
, we replace the 𝑖
-th element, otherwise we don't replace anything.

Why is it optimal? Essentially, this greedy approach always finds a non-unique subarray [𝑙,𝑟]
 with the lowest value of 𝑟
, and replaces the 𝑟
-th element. We obviously have to replace at least one element from the subarray [𝑙,𝑟]
; but replacing the 𝑟
-th element is optimal since we picked the lowest value of 𝑟
, so every non-unique subarray which contains any element from [𝑙,𝑟]
 also contains the 𝑟
-th element.

Okay, but we need to make this greedy solution work fast. When we consider the 𝑖
-th element, how do we check that there's a non-unique subarray starting after the element 𝑠
 and ending at the 𝑖
-th element? Suppose we go from the 𝑖
-th element to the left and maintain a counter; when we meet an element for the first time, we increase this counter; when we meet an element for the second time, we decrease this counter. If this counter is equal to 0
, then the current subarray is non-unique: every element appears at least twice. Otherwise, at least one element has exactly one occurrence.

Suppose we maintain an array 𝑡
 where for each integer present in the original array, we put 1
 in the last position we've seen this element, and −1
 in the second-to-last position we've seen this element (i. e. for every element, we consider its two last occurrences among the first 𝑖
 positions in the array, put 1
 in the last of them, and −1
 in the second-to-last of them). Then, if we go from 𝑖
 to 𝑙
 and maintain the counter in the same way as we described in the previous paragraph, the value of this counter will be equal to the sum of the corresponding segment in this array 𝑡
.

So, we want to check if there's a segment in the array 𝑡
 such that its left border is greater than 𝑠
 (the last position where we made a replacement), the right border is 𝑖
, and the sum is 0
. We can show that the sum on any segment ending in the 𝑖
-th position is currently non-negative; so, we actually want to find the segment with the minimum sum. We can store a segment tree that, for every position 𝑙
 from 𝑠+1
 to 𝑖
, maintains the sum on segment [𝑙,𝑖]
; then changing an element is just performing the query "add on segment", and finding the minimum sum is just performing the query "minimum on segment". This allows us to get a solution with complexity of 𝑂(𝑛log𝑛)
.