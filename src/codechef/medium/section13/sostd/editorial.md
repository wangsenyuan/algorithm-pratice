Let’s maintain an array of pairs P[1...N]P[1...N] such that P_i = (A_i,B_i)P 
i
​	
 =(A 
i
​	
 ,B 
i
​	
 )

First of all, let’s sort our pairs according to AA. Let’s iterate through the pairs one by one. Suppose we are processing the i^{th}i 
th
  pair P_iP 
i
​	
 . Let’s think how this pair contributes to the final answer.

If our pairs are sorted then clearly A_iA 
i
​	
  would be the maximum AA for any subset of the set containing the first ii elements (but must contain the i_{th}i 
th
​	
  one for sure). So we will have A_iA 
i
​	
  included in 2^{i-1}2 
i−1
  multiplicands (and all of them are added to the final answer). So now we got rid of one part of the problem. We still have BB coefficients.

Let’s maintain a segment tree with MaxV=10^6MaxV=10 
6
  leaves (or you can compress numbers and have NN leaves but such a thing is not necessary). In the i^{th}i 
th
  leave we record 22 values (consider leaves from left to right).

The first value is S_iS 
i
​	
  the number of subsets having max_B = imax 
B
​	
 =i

The second Tot_iTot 
i
​	
  is simply i*Si∗S (we will see why we need it).

So a node covering [L,R][L,R] in the segment tree will have SS equal to the number of subsets such that max_B \in [L,R]max 
B
​	
 ∈[L,R].

For this node TotTot will be equal to S_L*L \, + S_{L+1}*(L+1) \,+ S_{L+2}*(L+2) \, + .... + \, + S_{R}*RS 
L
​	
 ∗L+S 
L+1
​	
 ∗(L+1)+S 
L+2
​	
 ∗(L+2)+....++S 
R
​	
 ∗R

In other words Tot = Tot_L+Tot_{L+1}+.....+Tot_RTot=Tot 
L
​	
 +Tot 
L+1
​	
 +.....+Tot 
R
​	
 

Now let’s get back. Consider the i_{th}i 
th
​	
  pair. For all subsets having max_B \leq B_imax 
B
​	
 ≤B 
i
​	
  we can add our pair to these subsets and the swagness of all of them (after adding our pair) will be A_i*B_iA 
i
​	
 ∗B 
i
​	
 . Counting the number of subsets is just an interval sum query on the segment tree since we are recording the number of subsets for each interval. Let’s denote this sum by LessSubCountLessSubCount.

Now what about the other case, if max_B>B_imax 
B
​	
 >B 
i
​	
  ? We will have something like

A_i*{y_1}*S_{y_1}+A_i*y_2*S_{y_2}...A 
i
​	
 ∗y 
1
​	
 ∗S 
y 
1
​	
 
​	
 +A 
i
​	
 ∗y 
2
​	
 ∗S 
y 
2
​	
 
​	
 ...

Where y_ky 
k
​	
  is some value (representing some BB coefficient) and y_k>B_iy 
k
​	
 >B 
i
​	
 .

S_{y_k}S 
y 
k
​	
 
​	
  denotes the number of subsets such that max_B=y_kmax 
B
​	
 =y 
k
​	
 . Take A_iA 
i
​	
  as a common factor so we have A_i*({y_1}*S_{y_1}+{y_2}*S_{y_2}+....)A 
i
​	
 ∗(y 
1
​	
 ∗S 
y 
1
​	
 
​	
 +y 
2
​	
 ∗S 
y 
2
​	
 
​	
 +....).

Doesn’t this look familiar? Well, it’s the sum of TotTot values. So basically we just need to calculate the sum of TotTot values over the interval [B+1, MaxV][B+1,MaxV]. That’s it, we can calculate the contribution of our current pair assuming the segment tree contains correct values.

Now how we should update our segment tree?

When processing a new pair (A_i,B_i)(A 
i
​	
 ,B 
i
​	
 ) we should add all the subsets containing this pair to our segment tree. Above we calculated the number of subsets which will have max_b=B_imax 
b
​	
 =B 
i
​	
  and we denoted it by LessSubCountLessSubCount. Let’s add this value to the number of subsets residing in the B_{i_{th}}B 
i 
th
​	
 
​	
  leaf.

For all subsets which have max_B>B_imax 
B
​	
 >B 
i
​	
  their number (count) is multiplied by 22, because for each one we can add our pair and still the value of max_Bmax 
B
​	
  won’t change. So our segment tree should support the opreation of multiplying a range by 22. This is a straightforward application of lazy propagation. If you are not familiar with the latter technique you should read about it as I won’t be covering it. Supporting multiplying a range by 22 is no harder at all than supporting adding a number to a range (which is the most basic application of lazy propagation).

A thing also that you should take care about, is resetting the segment tree after each test case. Resetting the whole segment tree is time-consuming. Each update and sum query runs in O(log \, MaxV)O(logMaxV). Each update modifies log\,MaxVlogMaxV nodes. So after each test case we will have only N*log\,MaxVN∗logMaxV nodes to reset.

Again, as I said we can compress the numbers and have NN leaves only in our segment tree.

