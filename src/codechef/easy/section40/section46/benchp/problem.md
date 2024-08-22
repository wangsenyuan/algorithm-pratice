In the gym, Chef prefers to lift at least 
W
W grams during a bench press and if that's impossible, Chef considers his workout to be incomplete and feels bad.

The rod weighs 
W
r
W 
r
​
  grams and there are 
N
N other weights lying on the floor that weigh 
w
1
,
w
2
,
.
.
.
,
w
N
w 
1
​
 ,w 
2
​
 ,...,w 
N
​
  grams. To maintain balance and to ensure that there is no unnecessary load due to torque, it's important that the weights added to the left side are symmetric to the weights added to the right side. It is not required to use all of the weights. It is also not required to use any weights at all, if Chef feels satisfied lifting only the rod.

For example:

1
1 
2
2 
2
2 
1
1 
∣
∣Rod Center
∣
∣ 
1
1 
1
1 
1
1 
3
3 is a wrong configuration, but

1
1 
2
2 
3
3 
1
1 
∣
∣Rod Center
∣
∣ 
1
1 
3
3 
2
2 
1
1 is a right configuration.

Find whether Chef will be able to collect the required weights to feel satisfied.

###Input

The first line contains an integer 
T
T, the number of test cases. Then the test cases follow.
Each test case contains two lines of input.
The first line contains three space-separated integers 
N
,
W
,
W
r
N,W,W 
r
​
 .
The second line contains 
N
N space-separated integers 
w
1
,
w
2
,
…
,
w
N
w 
1
​
 ,w 
2
​
 ,…,w 
N
​
 .
###Output For each test case, output the answer in a single line: "YES" if Chef will be satisfied after his workout and "NO" if not (without quotes).

You may print each character of each string in uppercase or lowercase (for example, the strings "yEs", "yes", "Yes" and "YES" will all be treated as identical).