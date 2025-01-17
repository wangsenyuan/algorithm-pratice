You are given integer sequences of length 
N: 
A=(A 
1
​
 ,A 
2
​
 ,⋯,A 
N
​
 ) and 
B=(B 
1
​
 ,B 
2
​
 ,⋯,B 
N
​
 ), and an integer 
K.

You can perform the following operation zero or more times.

Choose integers 
i and 
j (
1≤i,j≤N). Here, 
∣i−j∣≤K must hold. Then, change the value of 
A 
i
​
  to 
A 
j
​
 .
Determine whether it is possible to make 
A identical to 
B.

There are 
T test cases for each input.


### ideas
1. 如果 a[i] = b[i], 似乎是可以不用变的（有一种情况要考虑）
2. 如果a[i] != b[i], 需要改变a[i] = some a[j]
3. 但是这个改变不一定是一次的，比如存在某个a[j] = b[i], 但是 j - i > k
4. 那么存在一个option，是先改变某个中间的位置a[x] = a[j], 然后把a[i] = a[x]
5. 所以理论上，可以将所有距离内的a[j] 变成 a[i]
6. 但是这里有个问题，就是a[x]被变成a[j]后，就不匹配b[x]了
7. 如果 a[x] != b[x], 那么使用它作为中继是没有关系的（因为它肯定要从其他位置变化来），但是如果a[x] = b[x], 那么就有点麻烦了
8. 感觉是一个图
9. 如果 i必须从最近的j那里得到需要的数, 那么