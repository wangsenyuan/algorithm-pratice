Denis was very sad after Nastya rejected him. So he decided to walk through the gateways to have some fun. And luck smiled at him! When he entered the first courtyard, he met a strange man who was selling something.

Denis bought a mysterious item and it was... Random permutation generator! Denis could not believed his luck.

When he arrived home, he began to study how his generator works and learned the algorithm. The process of generating a permutation consists of 𝑛
 steps. At the 𝑖
-th step, a place is chosen for the number 𝑖
 (1≤𝑖≤𝑛)
. The position for the number 𝑖
 is defined as follows:

For all 𝑗
 from 1
 to 𝑛
, we calculate 𝑟𝑗
  — the minimum index such that 𝑗≤𝑟𝑗≤𝑛
, and the position 𝑟𝑗
 is not yet occupied in the permutation. If there are no such positions, then we assume that the value of 𝑟𝑗
 is not defined.
For all 𝑡
 from 1
 to 𝑛
, we calculate 𝑐𝑜𝑢𝑛𝑡𝑡
  — the number of positions 1≤𝑗≤𝑛
 such that 𝑟𝑗
 is defined and 𝑟𝑗=𝑡
.
Consider the positions that are still not occupied by permutation and among those we consider the positions for which the value in the 𝑐𝑜𝑢𝑛𝑡
 array is maximum.
The generator selects one of these positions for the number 𝑖
. The generator can choose any position.
Let's have a look at the operation of the algorithm in the following example:


Let 𝑛=5
 and the algorithm has already arranged the numbers 1,2,3
 in the permutation. Consider how the generator will choose a position for the number 4
:

The values of 𝑟
 will be 𝑟=[3,3,3,4,×]
, where ×
 means an indefinite value.
Then the 𝑐𝑜𝑢𝑛𝑡
 values will be 𝑐𝑜𝑢𝑛𝑡=[0,0,3,1,0]
.
There are only two unoccupied positions in the permutation: 3
 and 4
. The value in the 𝑐𝑜𝑢𝑛𝑡
 array for position 3
 is 3
, for position 4
 it is 1
.
The maximum value is reached only for position 3
, so the algorithm will uniquely select this position for number 4
.
Satisfied with his purchase, Denis went home. For several days without a break, he generated permutations. He believes that he can come up with random permutations no worse than a generator. After that, he wrote out the first permutation that came to mind 𝑝1,𝑝2,…,𝑝𝑛
 and decided to find out if it could be obtained as a result of the generator.

Unfortunately, this task was too difficult for him, and he asked you for help. It is necessary to define whether the written permutation could be obtained using the described algorithm if the generator always selects the position Denis needs.

### ideas
1. 如果每次都重新计算r[j], 然后计算count[i], 这个太慢了
2. 我们考虑选择了位置p[i]，来放置i，怎么影响r[j]?
3. 所以，它左边的（包括它自己）r[j] = i 的位置，全部要更新到新的，右边的还没有设置值的最近的位置k
4. count[k] += count[j] - 1
5. 所以，这里一个是找右边，最近的k，已经更新k以后，找出最大的count[k]，用来放置i+1
6. 