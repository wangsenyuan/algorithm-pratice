Let's denote the score of an array 𝑏
 with 𝑘
 elements as ∑𝑘𝑖=1(∑𝑖𝑗=1𝑏𝑗)
. In other words, let 𝑆𝑖
 denote the sum of the first 𝑖
 elements of 𝑏
. Then, the score can be denoted as 𝑆1+𝑆2+…+𝑆𝑘
.

Skibidus is given 𝑛
 arrays 𝑎1,𝑎2,…,𝑎𝑛
, each of which contains 𝑚
 elements. Being the sigma that he is, he would like to concatenate them in any order to form a single array containing 𝑛⋅𝑚
 elements. Please find the maximum possible score Skibidus can achieve with his concatenated array!

Formally, among all possible permutations∗
 𝑝
 of length 𝑛
, output the maximum score of 𝑎𝑝1+𝑎𝑝2+⋯+𝑎𝑝𝑛
, where +
 represents concatenation†
.


### ideas
1. 如果 sum(row[i]) > sum(row[j])
2. 似乎 i在前面更优。这样子row[i]出现的次数更多
3. 如果 sum(row[i]) = sum(row[j]) 呢？
4. 好像有点复杂呢
5. 假设row[i]排在第一个位置
6. row[i][0] * (n * m) + row[i][1] * (n * m - 1) + ... row[i][m-1] * (n * m - (m - 1))
7. let L = n * m 
8. row[i][0] * L + row[i][1] * L - row[i][1] + ... + row[i][m-1] * L - row[i][m-1] * (m - 1)
9. sum(row[i]) * L - row[i][0] * 0 - row[i][1] * 1 - .... - row[i][m-1] * (m - 1)
10. 如果两个要比较
11.  
