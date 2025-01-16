There are 
N buildings, building 
1, building 
2, 
…, building 
N, arranged in this order in a straight line from west to east. Building 
1 is the westernmost, and building 
N is the easternmost. The height of building 
i (1≤i≤N) is 
H 
i
​
 .

For a pair of integers 
(i,j) (1≤i<j≤N), building 
j can be seen from building 
i if the following condition is satisfied.

There is no building taller than building 
j between buildings 
i and 
j. In other words, there is no integer 
k (i<k<j) such that 
H 
k
​
 >H 
j
​
 .
You are given 
Q queries. In the 
i-th query, given a pair of integers 
(l 
i
​
 ,r 
i
​
 ) (l 
i
​
 <r 
i
​
 ), find the number of buildings to the east of building 
r 
i
​
  (that is, buildings 
r 
i
​
 +1, 
r 
i
​
 +2, 
…, 
N) that can be seen from both buildings 
l 
i
​
  and 
r 
i
​
 .

 ### ideas
 1. 找到一个位置的可以看到的building是比较容易的
 2. 5，2，3，6 
 3. (1, 2) = 2 
 4. (3, 4)都可以被1，2看到
 5. (1, 2)的lcp是4（a[4] = 6)，4能够看到的，1,2都可以看到
 6. 关键是 3要怎么处理
 7. 5， 4， 2， 3， 6 
 8. (1, 3) = 1 (智能看到5，a[5] = 6)
 9. a[3][i] 满足 (l+1....a[u][i] 中间没有遮挡的条件？)