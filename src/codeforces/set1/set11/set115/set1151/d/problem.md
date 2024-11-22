During a break in the buffet of the scientific lyceum of the Kingdom of Kremland, there was formed a queue of 𝑛
 high school students numbered from 1
 to 𝑛
. Initially, each student 𝑖
 is on position 𝑖
. Each student 𝑖
 is characterized by two numbers — 𝑎𝑖
 and 𝑏𝑖
. Dissatisfaction of the person 𝑖
 equals the product of 𝑎𝑖
 by the number of people standing to the left of his position, add the product 𝑏𝑖
 by the number of people standing to the right of his position. Formally, the dissatisfaction of the student 𝑖
, which is on the position 𝑗
, equals 𝑎𝑖⋅(𝑗−1)+𝑏𝑖⋅(𝑛−𝑗)
.

The director entrusted Stas with the task: rearrange the people in the queue so that minimize the total dissatisfaction.

Although Stas is able to solve such problems, this was not given to him. He turned for help to you.

### ideas
1. a * (j - 1) + b * (n - j)
2. a * j - a + b * n - b * j
3. (a - b) * j - a + b * n 
4. a - b 越小，前面人越多越好？