Polycarp has to solve exactly 𝑛
 problems to improve his programming skill before an important programming competition. But this competition will be held very soon, most precisely, it will start in 𝑘
 days. It means that Polycarp has exactly 𝑘
 days for training!

Polycarp doesn't want to procrastinate, so he wants to solve at least one problem during each of 𝑘
 days. He also doesn't want to overwork, so if he solves 𝑥
 problems during some day, he should solve no more than 2𝑥
 problems during the next day. And, at last, he wants to improve his skill, so if he solves 𝑥
 problems during some day, he should solve at least 𝑥+1
 problem during the next day.

More formally: let [𝑎1,𝑎2,…,𝑎𝑘]
 be the array of numbers of problems solved by Polycarp. The 𝑖
-th element of this array is the number of problems Polycarp solves during the 𝑖
-th day of his training. Then the following conditions must be satisfied:

sum of all 𝑎𝑖
 for 𝑖
 from 1
 to 𝑘
 should be 𝑛
;
𝑎𝑖
 should be greater than zero for each 𝑖
 from 1
 to 𝑘
;
the condition 𝑎𝑖<𝑎𝑖+1≤2𝑎𝑖
 should be satisfied for each 𝑖
 from 1
 to 𝑘−1
.
Your problem is to find any array 𝑎
 of length 𝑘
 satisfying the conditions above or say that it is impossible to do it.

 ### ideas
 1. 1,2,3...k 这个是最小的数
 2. 如果 n < sum(1...k) => no answer
 3. n >= ... 
 4. 