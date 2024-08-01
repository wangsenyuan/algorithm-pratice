We define the MAD
 (Maximum Appearing Duplicate) in an array as the largest number that appears at least twice in the array. Specifically, if there is no number that appears at least twice, the MAD
 value is 0
.

For example, MAD([1,2,1])=1
, MAD([2,2,3,3])=3
, MAD([1,2,3,4])=0
.

You are given an array 𝑎
 of size 𝑛
. Initially, a variable 𝑠𝑢𝑚
 is set to 0
.

The following process will be executed in a sequential loop until all numbers in 𝑎
 become 0
:

Set 𝑠𝑢𝑚:=𝑠𝑢𝑚+∑𝑛𝑖=1𝑎𝑖
;
Let 𝑏
 be an array of size 𝑛
. Set 𝑏𝑖:= MAD([𝑎1,𝑎2,…,𝑎𝑖])
 for all 1≤𝑖≤𝑛
, and then set 𝑎𝑖:=𝑏𝑖
 for all 1≤𝑖≤𝑛
.
Find the value of 𝑠𝑢𝑚
 after the process.

 ### ideas
 1. b[i] = mad(a[...i])
 2. b[i] >= b[i-1]
 3. just bruteforce?
 4. 还需要更优的结果
 5. 一次操作后，必然满足 a[i] <= a[i+1] <= ... <= a[n]
 6. 如果 a[i] = a[i+1] 也就是在a[i]变成0前，它始终是存在的
 7. 但是如果 A[i] != a[i+1] 呢？
 8. 那么显然在下一次a[i],的时候，就不会存在了
 9. 如果连续m个x，出现了，那么它的贡献 = m * x + (m - 1) * x + (m - 2) * x ...
 10. 