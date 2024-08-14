The weight of a sequence is defined as the number of unordered pairs of indexes (𝑖,𝑗)
 (here 𝑖<𝑗
) with same value (𝑎𝑖=𝑎𝑗
). For example, the weight of sequence 𝑎=[1,1,2,2,1]
 is 4
. The set of unordered pairs of indexes with same value are (1,2)
, (1,5)
, (2,5)
, and (3,4)
.

You are given a sequence 𝑎
 of 𝑛
 integers. Print the sum of the weight of all subsegments of 𝑎
.

A sequence 𝑏
 is a subsegment of a sequence 𝑎
 if 𝑏
 can be obtained from 𝑎
 by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end.

 ### ideas
 1. 排序，似乎会影响结果
 2. f(arr) = 其中每个distinct 数字的 freq[x] 决定的 = freq[x] * (freq[x] - 1) / 2
 3. g(i) = 到i为止的arr[:i]的统计（包括它的子串）
 4. g(i) = g(i-1) + arr[i] 的贡献
 5. 假设arr[i] = x, 它出现了2次, 一次是在i, 另外一次是j（j < i), 那么 arr[i]的贡献 = j+1（包括(j...i)这个子串的所有子串)
 6. dp[x] = x的贡献 g(i) = g(i-1) + dp[x]
 7. dp[x] += (i+1)