Vasya has written some permutation 𝑝1,𝑝2,…,𝑝𝑛
 of integers from 1
 to 𝑛
, so for all 1≤𝑖≤𝑛
 it is true that 1≤𝑝𝑖≤𝑛
 and all 𝑝1,𝑝2,…,𝑝𝑛
 are different. After that he wrote 𝑛
 numbers 𝑛𝑒𝑥𝑡1,𝑛𝑒𝑥𝑡2,…,𝑛𝑒𝑥𝑡𝑛
. The number 𝑛𝑒𝑥𝑡𝑖
 is equal to the minimal index 𝑖<𝑗≤𝑛
, such that 𝑝𝑗>𝑝𝑖
. If there is no such 𝑗
 let's let's define as 𝑛𝑒𝑥𝑡𝑖=𝑛+1
.

In the evening Vasya went home from school and due to rain, his notebook got wet. Now it is impossible to read some written numbers. Permutation and some values 𝑛𝑒𝑥𝑡𝑖
 are completely lost! If for some 𝑖
 the value 𝑛𝑒𝑥𝑡𝑖
 is lost, let's say that 𝑛𝑒𝑥𝑡𝑖=−1
.

You are given numbers 𝑛𝑒𝑥𝑡1,𝑛𝑒𝑥𝑡2,…,𝑛𝑒𝑥𝑡𝑛
 (maybe some of them are equal to −1
). Help Vasya to find such permutation 𝑝1,𝑝2,…,𝑝𝑛
 of integers from 1
 to 𝑛
, that he can write it to the notebook and all numbers 𝑛𝑒𝑥𝑡𝑖
, which are not equal to −1
, will be correct.

### ideas
1. j = next[i] != -1, 那么在i...j中间的所有数字都比p[i]小
2. 那么所有i+1...j中间的-1都至多是j
3. 