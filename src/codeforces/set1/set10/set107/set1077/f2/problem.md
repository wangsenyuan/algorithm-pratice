### ideas

1. dp[i][j][k] 表示到i为止，前面已经选择了j个，且最近的在i的前面k个位置（k = 0时，表示在i处进行了选择)
2. 现在不能用上面的dp了，上面的dp是n*k*x的，现在就会到1e9了
3. 显然要降低一个维度
4. 这里有个贪心的想法是，在相同的位置i, 如果在已选择j次的时候，它的gain是w， 那么j+1次，至少要比这个好
5. 这里可以有个队列，就是当前i状态迁移时，只需要考虑前面k-1个长度的窗口，然后取其中的最大值

### solution

Let's use dynamic programming described in the previous tutorial to solve this problem too. But its complexity is 𝑂(
𝑛𝑘𝑥)
so we have to improve some part of the solution.

Let's see how we do transitions in this dp: for 𝑝∈[𝑖−𝑘;𝑖−1]
𝑑𝑝𝑖,𝑗=𝑚𝑎𝑥(𝑑𝑝𝑖,𝑗,𝑑𝑝𝑝,𝑗+1+𝑎𝑖)
. What can we do to optimize it? 𝑎𝑖
is the constant and we have to take the maximum value among 𝑑𝑝𝑖−𝑘,𝑗+1,𝑑𝑝𝑖−𝑘+1,𝑗+1,…,𝑑𝑝𝑖−1,𝑗+1
. You will say "segment tree"! I say no. Not a segment tree. Not a sparse table. Not a cartesian tree or some other
logarithmic data structures. If you want to spend a lot of time to fit such solution in time and memory limits — okay,
it is your choice. I prefer the queue with supporting the maximum on it.

The last part of this tutorial will be a small guide about how to write and use the queue with supporting the maximum on
it.

The first part of understanding this data structure is the stack with the maximum. How do we support the stack with the
maximum on it? That's pretty easy: let's maintain the stack of pairs, when the first value of pair is the value in the
stack and the second one is the maximum on the stack if this element will be the topmost. Then when we push some value
𝑣𝑎𝑙
in it, the first element of pair will be 𝑣𝑎𝑙
and the second one will be 𝑚𝑎𝑥(𝑣𝑎𝑙,𝑠.𝑡𝑜𝑝().𝑠𝑒𝑐𝑜𝑛𝑑)
(if 𝑠
is our stack and 𝑡𝑜𝑝()
is the topmost element). When we pop the element we don't need any special hacks to do it. Just pop it. And the maximum
on the stack is always 𝑠.𝑡𝑜𝑝().𝑠𝑒𝑐𝑜𝑛𝑑
.

Okay, the second part of understanding this data structure is the queue on two stacks. Let's maintain two stacks 𝑠1
and 𝑠2
and try to implement the queue using it. We will push elements only to 𝑠2
and pop elements only from 𝑠1
. Then how to maintain the queue using such stacks? The push is pretty easy — just push it in 𝑠2
. The main problem is pop. If 𝑠1
is not empty then we have to pop it from 𝑠1
. But what do we do if 𝑠1
is empty? No problems: let's just transfer elements of 𝑠2
to 𝑠1
(pop from 𝑠2
, push to 𝑠1
) in order from top to bottom. And don't forget to pop the element after this transfer!

Okay, if we will join these two data structures, we can see that we obtain exactly what we want! Just two stacks with
maximums! That's pretty easy to understand and implement it.

The last part of the initial solution is pretty easy — just apply this data structure (in fact, 𝑥+1
data structures) to do transitions in our dynamic programming. The implementation of this structure can be found in the
authors solution.

Total complexity of the solution is 𝑂(𝑛𝑥)
.