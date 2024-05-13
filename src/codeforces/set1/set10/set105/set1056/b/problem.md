Arkady and his friends love playing checkers on an 𝑛×𝑛
 field. The rows and the columns of the field are enumerated from 1
 to 𝑛
.

The friends have recently won a championship, so Arkady wants to please them with some candies. Remembering an old parable (but not its moral), Arkady wants to give to his friends one set of candies per each cell of the field: the set of candies for cell (𝑖,𝑗)
 will have exactly (𝑖2+𝑗2)
 candies of unique type.

There are 𝑚
 friends who deserve the present. How many of these 𝑛×𝑛
 sets of candies can be split equally into 𝑚
 parts without cutting a candy into pieces? Note that each set has to be split independently since the types of candies in different sets are different.

Input
The only line contains two integers 𝑛
 and 𝑚
 (1≤𝑛≤109
, 1≤𝑚≤1000
) — the size of the field and the number of parts to split the sets into.


### ideas
1. (i * i + j * j) % m = 0
2. 但是 n 有1e9这么大，所以没法迭代处理，找规律
3. 如果 i % m = 0 and j % m = 0, 那么 那么上式肯定也可以整除m
4. 如果 i * i % m + j * j % m = 0/3 那么也是ok的
5. 所以，就是找 1...n 中 i * i % m = 0,1,2...m-1的个数
6. 先算出 i % m = x 的个数， 然后再计算 i * i % m 的个数
7. 