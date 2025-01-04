You are given 𝑛
 pairs of integers (𝑎1,𝑏1),(𝑎2,𝑏2),…,(𝑎𝑛,𝑏𝑛)
. All of the integers in the pairs are distinct and are in the range from 1
 to 2⋅𝑛
 inclusive.

Let's call a sequence of integers 𝑥1,𝑥2,…,𝑥2𝑘
 good if either

𝑥1<𝑥2>𝑥3<…<𝑥2𝑘−2>𝑥2𝑘−1<𝑥2𝑘
, or
𝑥1>𝑥2<𝑥3>…>𝑥2𝑘−2<𝑥2𝑘−1>𝑥2𝑘
.
You need to choose a subset of distinct indices 𝑖1,𝑖2,…,𝑖𝑡
 and their order in a way that if you write down all numbers from the pairs in a single sequence (the sequence would be 𝑎𝑖1,𝑏𝑖1,𝑎𝑖2,𝑏𝑖2,…,𝑎𝑖𝑡,𝑏𝑖𝑡
), this sequence is good.

What is the largest subset of indices you can choose? You also need to construct the corresponding index sequence 𝑖1,𝑖2,…,𝑖𝑡
.

### ideas
1. 确定一个顺序， 比如 x1 > x2 < x3  > ... 
2. 然后选择那些在一对中满足>条件的pair
3. 然后根据第一个数字排序 x1 < x3 < x5... 
4. 这样就ok了