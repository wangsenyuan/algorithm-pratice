Nikolay has only recently started in competitive programming, but already qualified to the finals of one prestigious olympiad. There going to be 𝑛
 participants, one of whom is Nikolay. Like any good olympiad, it consists of two rounds. Tired of the traditional rules, in which the participant who solved the largest number of problems wins, the organizers came up with different rules.

Suppose in the first round participant A took 𝑥
-th place and in the second round — 𝑦
-th place. Then the total score of the participant A is sum 𝑥+𝑦
. The overall place of the participant A is the number of participants (including A) having their total score less than or equal to the total score of A. Note, that some participants may end up having a common overall place. It is also important to note, that in both the first and the second round there were no two participants tying at a common place. In other words, for every 𝑖
 from 1
 to 𝑛
 exactly one participant took 𝑖
-th place in first round and exactly one participant took 𝑖
-th place in second round.

Right after the end of the Olympiad, Nikolay was informed that he got 𝑥
-th place in first round and 𝑦
-th place in the second round. Nikolay doesn't know the results of other participants, yet he wonders what is the minimum and maximum place he can take, if we consider the most favorable and unfavorable outcome for him. Please help Nikolay to find the answer to this question.

### ideas
1. 要得到最小的名次，x+y要最小。
2. 那么除了x和y的，1...n, 1.。。。n要交差组合
3. 1+n, 2 + n - 1... n + 1 这样
4. 如果 1 + n >= x + y 那么可以获得1的位置
5. 否则 1 + n < x + y, 那么使用n来匹配1就太浪费了，应该使用(1, 1)
6. i + n >= x + y => i = x + y - n（i就是A所能获得最小排名) 
7. 反过来，如何获得最大排名(假设 x <= y)
8. 不能把1个过早的用掉，这样子n就太大了，所以应该反过来
9. n + i <= x + y => i <= x + y - n
10. 如果 i < 0, 那么也就是说没有数字可以和n匹配，不妨将n给使用(n, n)使用掉
11. 1 + i <= x + y => i <= x + y - 1
12. 