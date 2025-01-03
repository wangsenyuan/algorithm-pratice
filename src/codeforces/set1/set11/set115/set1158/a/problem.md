𝑛
 boys and 𝑚
 girls came to the party. Each boy presented each girl some integer number of sweets (possibly zero). All boys are numbered with integers from 1
 to 𝑛
 and all girls are numbered with integers from 1
 to 𝑚
. For all 1≤𝑖≤𝑛
 the minimal number of sweets, which 𝑖
-th boy presented to some girl is equal to 𝑏𝑖
 and for all 1≤𝑗≤𝑚
 the maximal number of sweets, which 𝑗
-th girl received from some boy is equal to 𝑔𝑗
.

More formally, let 𝑎𝑖,𝑗
 be the number of sweets which the 𝑖
-th boy give to the 𝑗
-th girl. Then 𝑏𝑖
 is equal exactly to the minimum among values 𝑎𝑖,1,𝑎𝑖,2,…,𝑎𝑖,𝑚
 and 𝑔𝑗
 is equal exactly to the maximum among values 𝑏1,𝑗,𝑏2,𝑗,…,𝑏𝑛,𝑗
.

You are interested in the minimum total number of sweets that boys could present, so you need to minimize the sum of 𝑎𝑖,𝑗
 for all (𝑖,𝑗)
 such that 1≤𝑖≤𝑛
 and 1≤𝑗≤𝑚
. You are given the numbers 𝑏1,…,𝑏𝑛
 and 𝑔1,…,𝑔𝑚
, determine this number.



### ideas
1. 这个题目有点绕
2. 就是假设boy1 给所有的女孩送糖果，a[1] 就是其中最少的那个数
3. 假设girl1收到了所有男孩的糖果，b[1]就是她收到的最多的那个数
4. 要计算所有男孩送出去的最少的糖果数
5. min(b[i]) >= max(a[j])
6. a[i]等于10，表示这个男孩给每个女孩送的至少有10个糖果，所以女孩收到的至少是10个
7. 