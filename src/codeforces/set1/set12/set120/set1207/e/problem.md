The jury picked an integer 𝑥
 not less than 0
 and not greater than 214−1
. You have to guess this integer.

To do so, you may ask no more than 2
 queries. Each query should consist of 100
 integer numbers 𝑎1
, 𝑎2
, ..., 𝑎100
 (each integer should be not less than 0
 and not greater than 214−1
). In response to your query, the jury will pick one integer 𝑖
 (1≤𝑖≤100
) and tell you the value of 𝑎𝑖⊕𝑥
 (the bitwise XOR of 𝑎𝑖
 and 𝑥
). There is an additional constraint on the queries: all 200
 integers you use in the queries should be distinct.

 ### ideas
 1. a[i] ^ x 
 2. 如果所有ai的前半部分都是0，那么 y = a[i] ^ x的前半部分= x的前半部分
 3. 或者更简单一点，就把x < 1<<7（128）找100个数就可以了吧，然后取结果的前7位