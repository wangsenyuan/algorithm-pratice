Ujan has a lot of numbers in his boxes. He likes order and balance, so he decided to reorder the numbers.

There are k
 boxes numbered from 1
 to k
. The i
-th box contains ni
 integer numbers. The integers can be negative. All of the integers are distinct.

Ujan is lazy, so he will do the following reordering of the numbers exactly once. He will pick a single integer from each of the boxes, k
 integers in total. Then he will insert the chosen numbers — one integer in each of the boxes, so that the number of integers in each box is the same as in the beginning. Note that he may also insert an integer he picked from a box back into the same box.

Ujan will be happy if the sum of the integers in each box is the same. Can he achieve this and make the boxes perfectly balanced, like all things should be?

### ideas
1. sum = all integers sum
2. sum % k = 0, 所以每个box最后的期望sum是确定的
3. sum[i]表示第i个box一开始的结果, sum[i] - x + y = sum
4. x - y = sum[i] - sum 也就是说拿走的和insert的数字的差是确定的
5. 但是这里存在环的情况，就是从box 1中取出数字x，从box 2中取出数字y， 3.。。z
6. 然后把z放入1，把x放入2，把y放入3，也是一种合理的状态
7. dp[state] = 表示是否能够得到state表示的集合的一种排列 = dp[sub1] and dp[sub2] ?
8. 其中sub1，就当作这种环形的排列进行处理
9. 那么这个sub1怎么判断呢？不可能每个数字都判断一遍吧？
10. 1 << 15 * k * 5000 ?
11. 1024 * 32 * 15 * 5000 不大行呐
12. 这里可以先建一下边（因为每个数字都是唯一的）
13. 假设在集合i中可以移除x，那么y是确定的，然后用y,这样子就可以吧这两个集合给连接起来（用x的有向边）
14. 然后沿着这样的边，移动一圈后，回到原来的节点，那么它们就是一条合理的安排
15. 