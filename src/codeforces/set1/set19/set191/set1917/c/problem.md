You have an array of integers 𝑎1,𝑎2,…,𝑎𝑛
 of length 𝑛
. On the 𝑖
-th of the next 𝑑
 days you are going to do exactly one of the following two actions:

Add 1
 to each of the first 𝑏𝑖
 elements of the array 𝑎
 (i.e., set 𝑎𝑗:=𝑎𝑗+1
 for each 1≤𝑗≤𝑏𝑖
).
Count the elements which are equal to their position (i.e., the 𝑎𝑗=𝑗
). Denote the number of such elements as 𝑐
. Then, you add 𝑐
 to your score, and reset the entire array 𝑎
 to a 0
-array of length 𝑛
 (i.e., set [𝑎1,𝑎2,…,𝑎𝑛]:=[0,0,…,0]
).
Your score is equal to 0
 in the beginning. Note that on each day you should perform exactly one of the actions above: you cannot skip a day or perform both actions on the same day.

What is the maximum score you can achieve at the end?

Since 𝑑
 can be quite large, the sequence 𝑏
 is given to you in the compressed format:

You are given a sequence of integers 𝑣1,𝑣2,…,𝑣𝑘
. The sequence 𝑏
 is a concatenation of infinitely many copies of 𝑣
: 𝑏=[𝑣1,𝑣2,…,𝑣𝑘,𝑣1,𝑣2,…,𝑣𝑘,…]
.


### ideas
1. d <= 1e9, n <= 2000
2. 这个题目没法模拟，因为每天有两种选择，每种选择都会对后续的结果产生影响
3. 所以要找规律
4. 似乎是不会的，因为后面的始终是在前缀上添加，有可能把前缀里面已经可以得分的给替换掉
5. 且还有一个就是，如果a[i]在时刻x，a[j]在时刻y, i < j, 那么肯定有 x <= y
6. 也就是说，只要查找a[1]在什么时刻能够满足条件，a[2]在什么时刻得分
7. a[i]在什么时刻得分。相同的放在一起
8. 得分后，再检查从0开始什么时候能满足即可