You are given an array consisting of 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 and an integer 𝑥
. It is guaranteed that for every 𝑖
, 1≤𝑎𝑖≤𝑥
.

Let's denote a function 𝑓(𝑙,𝑟)
 which erases all values such that 𝑙≤𝑎𝑖≤𝑟
 from the array 𝑎
 and returns the resulting array. For example, if 𝑎=[4,1,1,4,5,2,4,3]
, then 𝑓(2,4)=[1,1,5]
.

Your task is to calculate the number of pairs (𝑙,𝑟)
 such that 1≤𝑙≤𝑟≤𝑥
 and 𝑓(𝑙,𝑟)
 is sorted in non-descending order. Note that the empty array is also considered sorted.

 ### ideas 
 1. f(l, r) 表示删除了 l <= a[?] <= r 的元素后，剩余的数组
 2. 定义 g(u) 表示保留1....u后，剩余的元素组成的数组，且这个数组是递增的，g(u) = v, 是这个数组最后的下标
 3. 这个数组肯定是递增的, g(0) = 0, g(1) <= g(2) <= .. <= g(i) ...
 4. f(u) 表示保留 u...x后，数组是递增的，且f(u) = v表示最小的数组
 5. 对于u来说，如果它是l，那么删除的 r, 要满足 f(r) >= g(l) 