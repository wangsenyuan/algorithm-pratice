Let 
N be a positive integer. You are given a string 
s of length 
N−1, consisting of < and >.

Find the number of permutations 
(p 
1
​
 ,p 
2
​
 ,…,p 
N
​
 ) of 
(1,2,…,N) that satisfy the following condition, modulo 
10 
9
 +7:

For each 
i (
1≤i≤N−1), 
p 
i
​
 <p 
i+1
​
  if the 
i-th character in 
s is <, and 
p 
i
​
 >p 
i+1
​
  if the 
i-th character in 
s is >.
Constraints
N is an integer.
2≤N≤3000
s is a string of length 
N−1.
s consists of < and >.


### ideas
此问题使用通常称为“插入 DP ” 的 DP 方法。
这个解释有下面的DP表。
d p ​​[ i ] [ j ]=（确定直到第i个值的位置关系时，为第i个值确定的值是第j个最小数的情况数）​​​

如果将此 DP 表的状态比作放置一个球，则如下图所示。
这次我就根据这些条件来解释一下。
