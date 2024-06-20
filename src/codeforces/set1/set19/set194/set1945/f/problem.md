As soon as everyone in the camp fell asleep, Kirill sneaked out of the tent and went to the Wise Oak to gather mushrooms.

It is known that there are 𝑛
 mushrooms growing under the Oak, each of which has magic power 𝑣𝑖
. Kirill really wants to make a magical elixir of maximum strength from the mushrooms.

The strength of the elixir is equal to the product of the number of mushrooms in it and the minimum magic power among these mushrooms. To prepare the elixir, Kirill will sequentially pick one mushroom growing under the Oak. Kirill can gather mushrooms in any order.

However, it's not that simple. The Wise Oak informed Kirill of a permutation of numbers 𝑝
 from 1
 to 𝑛
. If Kirill picks only 𝑘
 mushrooms, then the magic power of all mushrooms with indices 𝑝1,𝑝2,…,𝑝𝑘−1
 will become 0
. Kirill will not use mushrooms with zero magic power to prepare the elixir.

Your task is to help Kirill gather mushrooms in such a way that he can brew the elixir of maximum possible strength. However, Kirill is a little scared to stay near the oak for too long, so out of all the suitable options for gathering mushrooms, he asks you to find the one with the minimum number of mushrooms.

A permutation of length 𝑛
 is an array consisting of 𝑛
 different integers from 1
 to 𝑛
 in any order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears in the array twice) and [1,3,4]
 is also not a permutation (𝑛=3
, but 4
 appears in the array).


### ideas
1. 对于permutation p，如果选择了k个，那么前k-1个p[i]的magic power就会变成0，所以收益也就是0
2. 所以，当k = 1时，只能最大值
3.    当k=2时，只能选择除了p[0]外的最大值
4.    k= 3时，只能选择除了p[0],p[1]外的 两个最大的值中的最小值
5.    所以要用 avl树