Let's call a list of positive integers 𝑎0,𝑎1,...,𝑎𝑛−1
 a power sequence if there is a positive integer 𝑐
, so that for every 0≤𝑖≤𝑛−1
 then 𝑎𝑖=𝑐𝑖
.

Given a list of 𝑛
 positive integers 𝑎0,𝑎1,...,𝑎𝑛−1
, you are allowed to:

Reorder the list (i.e. pick a permutation 𝑝
 of {0,1,...,𝑛−1}
 and change 𝑎𝑖
 to 𝑎𝑝𝑖
), then
Do the following operation any number of times: pick an index 𝑖
 and change 𝑎𝑖
 to 𝑎𝑖−1
 or 𝑎𝑖+1
 (i.e. increment or decrement 𝑎𝑖
 by 1
) with a cost of 1
.
Find the minimum cost to transform 𝑎0,𝑎1,...,𝑎𝑛−1
 into a power sequence.

### ideas
 1. c 太大的时候，其实就没啥意义了
 2. 因为 c(n-1) 太大了
 3. 也就是说当 n > 50 的时候， c 只能等于1