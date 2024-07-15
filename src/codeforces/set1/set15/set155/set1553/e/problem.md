An identity permutation of length 𝑛
 is an array [1,2,3,…,𝑛]
.

We performed the following operations to an identity permutation of length 𝑛
:

firstly, we cyclically shifted it to the right by 𝑘
 positions, where 𝑘
 is unknown to you (the only thing you know is that 0≤𝑘≤𝑛−1
). When an array is cyclically shifted to the right by 𝑘
 positions, the resulting array is formed by taking 𝑘
 last elements of the original array (without changing their relative order), and then appending 𝑛−𝑘
 first elements to the right of them (without changing relative order of the first 𝑛−𝑘
 elements as well). For example, if we cyclically shift the identity permutation of length 6
 by 2
 positions, we get the array [5,6,1,2,3,4]
;
secondly, we performed the following operation at most 𝑚
 times: pick any two elements of the array and swap them.
You are given the values of 𝑛
 and 𝑚
, and the resulting array. Your task is to find all possible values of 𝑘
 in the cyclic shift operation.


 ### ideas
 1. 操作1，可以看2 * n次swap ?
 2. 操作1，不可能改变两个元素的相对位置。所以改变位置的，只有操作2
 3. 先假设没有操作1， a, b, c, d 假设交换了 a, d 或者 a, c, 对原数组的影响是什么？
 4. 如果没有操作2， 1, 2, 3, 4 => (k = 3) 2, 3, 4, 1; 可以直接通过a[0] = n - k + 1 确定 (k = n - a[0] + 1)
 5. 但是因为有操作2，所以，a[0], 可能是被交换走了
 6. 如果m=1，那么只能交换一次， 那么必然有 n - 2 个位置要符合原来的顺序
 7. 如果m=2, 那么可以交换2次，n-3, n - 4 个位置, 1, 2, 3 (2, 1, 3) => (2, 3, 1)
 8. 。。。
 9. m = i, 可以交换i次， n - (i + 1)个位置保持不变（最少i+1个位置发生了变化), 最少 n - 2 * i 个位置(最多 2 * i 个元素)
 10. 那么剩下的元素就组成了 lis?
1.  

### solution

Let's decrease all numbers by 1
 and start the numeration from 0
, because cyclic shifts are very easy to describe this way. Let's observe for 𝑛=4
:

𝑘=0
. 𝑝=[0,1,2,3]
. So, 𝑝𝑖=𝑖
.
𝑘=1
. 𝑝=[3,0,1,2]
. So, 𝑝𝑖=(𝑖−1)mod𝑛
.
...
Continuing this process, we verify that indeed, 𝑝𝑖=(𝑖−𝑘)mod𝑛
. Very simple!

Now suppose we have some value 0≤𝑘≤𝑛−1
 and we want to check if it is possible to obtain 𝑝
 from 𝑘
-th cyclic shift by doing at most 𝑚
 swaps. For this, we can calculate the minimum number of swaps and check it is not more than 𝑚
.

So, how to calculate the minimum number of swaps needed to transform a permutation 𝑎
 to another permutation 𝑏
? This is actually a well-known problem. The idea is, we build a graph with undirected edges (𝑎𝑖,𝑏𝑖)
. The minimum number of swaps will be equal to 𝑛−𝑐
, where 𝑐
 is equal to the number of connected components in the resulting graph.

Nice, now we can check if some 𝑘
 is good in 𝑂(𝑛)
 time. But we can't check all of them, right? Here comes the crucial observation:

Suppose you get a permutation 𝑎
 after a cyclic shift. Then you make at most 𝑚
 swaps and obtain 𝑏
. This means at most 2⋅𝑚
 numbers will be out of order! That is, there will be at least 𝑛−2⋅𝑚
 indices 𝑖
 such that 𝑎𝑖=𝑏𝑖
.

So can we calculate the number 𝑐𝑛𝑡𝑘
 — the count of integers in position for each cyclic shift 𝑘
? Yes, we can! For an arbitrary 𝑖
, there is exactly one 𝑘
 such that 𝑝𝑖=(𝑖−𝑘)mod𝑛
.

But wait, it means there are in total only 𝑛
 good positions because ∑𝑐𝑛𝑡𝑖=𝑛
! And we check only those 𝑘
 for which it is true that 𝑐𝑛𝑡𝑘≥𝑛−2⋅𝑚
. Remember that weird constraint 𝑚≤𝑛3
? Well, turns out there are at most 𝑛𝑛−2𝑛3=3
 different 𝑘
 to consider!

So we know we check at most 3
 different values and we know how to check in 𝑂(𝑛)
 time. That concludes the solution. The time and space complexities are 𝑂(𝑛)
.