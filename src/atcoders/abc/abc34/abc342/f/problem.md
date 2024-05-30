You will play a game against a dealer. The game uses a 
D-sided die (dice) that shows an integer from 
1 to 
D with equal probability, and two variables 
x and 
y initialized to 
0. The game proceeds as follows:

You may perform the following operation any number of times: roll the die and add the result to 
x. You can choose whether to continue rolling or not after each roll.

Then, the dealer will repeat the following operation as long as 
y<L: roll the die and add the result to 
y.

If 
x>N, you lose. Otherwise, you win if 
y>N or 
x>y and lose if neither is satisfied.

Determine the probability of your winning when you act in a way that maximizes the probability of winning.

### ideas
1. 考虑L = 1 的情况（L比较小的情况）
2. 且 y = L, 如果 L >= N (似乎很容易获胜)
3. 这种情况下只要保证 x < N 即可
4. 必须保证 x <= N 
5.  这个概率是多少？
6.  

### solution

First, let us find the probability 
p[i] that 
y=i. This can be found with dynamic programming (DP).

Specifically, start with 
p[0]=1, and for each 
i=0,1,…,l−1, add 
D
p[i]
​
  to 
p[i+1],…,p[i+D]. This can be done fast with cumulative sum trick. Or, one can regard it as receiving DP to process it using cumulative sum.

Then, let us find the probability of winning 
q[i] when stopping at 
x=i. This can be found as cumulative sums of 
p[i], which was found above.

Finally, find the probability of winning 
r[i] when 
x=i. Corresponding to whether or not to roll the die, we have

r[i]=max(q[i], 
D
1
​
  
j=i+1
∑
i+D
​
 r[j]),

so 
r can be found in descending order of 
i. The resulting 
r[0] is the answer.