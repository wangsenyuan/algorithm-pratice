You are given an array 𝑎
 of size 𝑛
. You will do the following process to calculate your penalty:

Split array 𝑎
 into two (possibly empty) subsequences†
 𝑠
 and 𝑡
 such that every element of 𝑎
 is either in 𝑠
 or 𝑡‡
.
For an array 𝑏
 of size 𝑚
, define the penalty 𝑝(𝑏)
 of an array 𝑏
 as the number of indices 𝑖
 between 1
 and 𝑚−1
 where 𝑏𝑖<𝑏𝑖+1
.
The total penalty you will receive is 𝑝(𝑠)+𝑝(𝑡)
.
If you perform the above process optimally, find the minimum possible penalty you will receive.

†
 A sequence 𝑥
 is a subsequence of a sequence 𝑦
 if 𝑥
 can be obtained from 𝑦
 by the deletion of several (possibly, zero or all) elements.

‡
 Some valid ways to split array 𝑎=[3,1,4,1,5]
 into (𝑠,𝑡)
 are ([3,4,1,5],[1])
, ([1,1],[3,4,5])
 and ([],[3,1,4,1,5])
 while some invalid ways to split 𝑎
 are ([3,4,5],[1])
, ([3,1,4,1],[1,5])
 and ([1,3,4],[5,1])
.

### ideas
1. 如果 a[i] < a[i+1], 显然不能把它们放在一起，放在一起，肯定贡献1，但是放在两边，只有有可能减少1
2. 贪心的处理就好了吗？
3. 似乎很难证明 
4. 把数组a变成b, 
5. b[i] = 1 if a[i] > a[i-1]
6. 如果放在一组内，那么就是sum
7. 但是如果在某个1的前面断开，要么 a[i]跟在 a[i-1]的后面, 要么a[i]跟在a[j]的后面
8. 且 a[j] > a[i]