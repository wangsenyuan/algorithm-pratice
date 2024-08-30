You are given an array 𝑎
 consisting of 𝑛
 integers. Indices of the array start from zero (i. e. the first element is 𝑎0
, the second one is 𝑎1
, and so on).

You can reverse at most one subarray (continuous subsegment) of this array. Recall that the subarray of 𝑎
 with borders 𝑙
 and 𝑟
 is 𝑎[𝑙;𝑟]=𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
.

Your task is to reverse such a subarray that the sum of elements on even positions of the resulting array is maximized (i. e. the sum of elements 𝑎0,𝑎2,…,𝑎2𝑘
 for integer 𝑘=⌊𝑛−12⌋
 should be maximum possible).

You have to answer 𝑡
 independent test cases.

### ideas
1. 假设reverse了区间(l...r), 1， 2， 3， 4， 5， 6
2. 考虑4中情况, l是偶数，r是偶数， 比如上面的[2, 4], 显然这个转换不会改变所有下标的parity；
3. 同样的，对于l是奇数，r是偶数，仍然是成立的
4. 所以，剩下l是奇数，r是偶数/ l是偶数，r是奇数的情况
5. 先考虑r是偶数的情况
6. 考虑even[i]/odd[i]分别表示i（偶数/奇数）前面的前缀和
7. 那么r是偶数时， 
8. 在转变前 = even[r]
9. 转变后 even[l-1] + odd[r] - odd[l-1],也就是取 odd[r] + even[l-1] - odd[l-1]的最大值
10. 如果 l是偶数，r是奇数
11. 转变前 = even[r] = even[r-1]
12. 转变后 = odd[r] - odd[l-1] + even[l-1] (仍然是成立的)
13. got