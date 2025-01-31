You are given sequences of length 
N, 
A=(A 
1
​
 ,A 
2
​
 ,…,A 
N
​
 ) and 
B=(B 
1
​
 ,B 
2
​
 ,…,B 
N
​
 ).
You are also given 
Q queries to process in order.

There are three types of queries:

1 l r x : Add 
x to each of 
A 
l
​
 ,A 
l+1
​
 ,…,A 
r
​
 .
2 l r x : Add 
x to each of 
B 
l
​
 ,B 
l+1
​
 ,…,B 
r
​
 .
3 l r : Print the remainder of 
i=l
∑
r
​
 (A 
i
​
 ×B 
i
​
 ) when divided by 
998244353.

### ideas
1. (a + x) * (b + y) = a * b + a * y + b * x + x * y