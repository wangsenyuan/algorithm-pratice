Time Limit: 2 sec / Memory Limit: 1024 MB

Score : 
500 points

Problem Statement
This is an interactive task, where your and the judge's programs interact via Standard Input and Output.

You and the judge will follow the procedure below. The procedure consists of phases 
1 and 
2; phase 
1 is immediately followed by phase 
2.

(Phase 
1)

The judge decides an integer 
N between 
1 and 
10 
9
  (inclusive), which is hidden.
You print an integer 
M between 
1 and 
110 (inclusive).
You also print an integer sequence 
A=(A 
1
​
 ,A 
2
​
 ,…,A 
M
​
 ) of length 
M such that 
1≤A 
i
​
 ≤M for all 
i=1,2,…,M.
(Phase 
2)

The judge gives you an integer sequence 
B=(B 
1
​
 ,B 
2
​
 ,…,B 
M
​
 ) of length 
M. Here, 
B 
i
​
 =f 
N
 (i). 
f(i) is defined by 
f(i)=A 
i
​
  for all integers 
i between 
1 and 
M (inclusive), and 
f 
N
 (i) is the integer resulting from replacing 
i with 
f(i) 
N times.
Based on the given 
B, you identify the integer 
N that the judge has decided, and print 
N.
After the procedure above, terminate the program immediately to be judged correct.

### ideas
1. fn(i) = 在一个圈里面行进了n次后的结果
2. 假设a中有一个圈的长度是3，且通过 fn(i)和i的相对位置，就知道了 n % 3 = r 的余数
3. n % 2 = x2, n % 3 = x2, n % 5 = x5, ... n % p = xp
4. 怎么求解n呢？
5. n是10的9次方，也不能迭代
6. 似乎要用到逆元计算。是不是搞错了？
7. 特别的，如果n是一个很大的质数，那么使用比较小的质数，其实是计算不出来的
8. 