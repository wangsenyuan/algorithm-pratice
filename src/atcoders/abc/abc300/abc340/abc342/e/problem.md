In the country of AtCoder, there are 
N stations: station 
1, station 
2, 
…, station 
N.

You are given 
M pieces of information about trains in the country. The 
i-th piece of information 
(1≤i≤M) is represented by a tuple of six positive integers 
(l 
i
​
 ,d 
i
​
 ,k 
i
​
 ,c 
i
​
 ,A 
i
​
 ,B 
i
​
 ), which corresponds to the following information:

For each 
t=l 
i
​
 ,l 
i
​
 +d 
i
​
 ,l 
i
​
 +2d 
i
​
 ,…,l 
i
​
 +(k 
i
​
 −1)d 
i
​
 , there is a train as follows:
The train departs from station 
A 
i
​
  at time 
t and arrives at station 
B 
i
​
  at time 
t+c 
i
​
 .
No trains exist other than those described by this information, and it is impossible to move from one station to another by any means other than by train.
Also, assume that the time required for transfers is negligible.

Let 
f(S) be the latest time at which one can arrive at station 
N from station 
S.
More precisely, 
f(S) is defined as the maximum value of 
t for which there is a sequence of tuples of four integers 
((t 
i
​
 ,c 
i
​
 ,A 
i
​
 ,B 
i
​
 )) 
i=1,2,…,k
​
  that satisfies all of the following conditions:

t≤t 
1
​
 
A 
1
​
 =S,B 
k
​
 =N
B 
i
​
 =A 
i+1
​
  for all 
1≤i<k,
For all 
1≤i≤k, there is a train that departs from station 
A 
i
​
  at time 
t 
i
​
  and arrives at station 
B 
i
​
  at time 
t 
i
​
 +c 
i
​
 .
t 
i
​
 +c 
i
​
 ≤t 
i+1
​
  for all 
1≤i<k.
If no such 
t exists, set 
f(S)=−∞.

Find 
f(1),f(2),…,f(N−1).