You are given an integer sequence 
A=(A 
1
​
 ,A 
2
​
 ,…,A 
N
​
 ) of length 
N.
You can perform the following operation any number of times, possibly zero:

Choose an integer pair 
(i,j) satisfying 
1≤i<j≤N, and replace 
A 
i
​
  with 
A 
i
​
 +1 and 
A 
j
​
  with 
A 
j
​
 −1.
Determine whether it is possible to make 
A a non-decreasing sequence through the operations.

You are given 
T test cases. Solve each of them.

### ideas
1. 考虑前缀和，pref[i] = a[1] + .. + a[i]
2. 交换(i, j) => 增加pref[i]（以及后面的），减小pref[j]（以及后面的）
3. 原数组非递减，表明pref[i]-pref[i-1]的差分非递减
4. 考虑i， pref[i-2] - pref[i-1] > pref[i] - pref[i-1] 
5. 这里必须增加pref[i]（或者减小pref[i-1])，貌似减小pref[i-1]更合理
6. 大概有点想法了
7. 假设前面能够得到的最小的差分是x,在得到x的时候，需要从后面borrorw一些数（这些数需要还给后面）
8. 