This problem is interactive, and the judge is adaptive. See the notes for details.
Also, the parameters in the problem statement are fixed at 
N=1000, 
M=10, 
Q=950.

There are 
N coins numbered 
1,2,…,N.
Exactly 
M of these coins are counterfeit.

An appraiser can, in one appraisal, determine whether two coins are of the same type or different types. Specifically:

If the two coins are both genuine or both counterfeit, they are judged to be of the same type.
Otherwise, they are judged to be of different types.
Identify all the counterfeit coins using at most 
Q appraisals.

### ideas
1. n = 1000, m = 10, Q = 950
2. 随便找出两个来，? x, y, 大概率它们是相同的（且都是真实的）
3. 所以不能随便拿两个出来
4. 两个一组（1， 2）(3, 4) => 在500次内
5. 存在几种可能性， 所有的回答是相同的（最坏的情况）
6. 存在其中几种（x, y)不同 <= m
7. 不大行
8. (1, 2), (1, 3).... (1, 11), (1, 20)
9. 如果相同的数量 >= 10, 那么1肯定是真实的
10. 否则它肯定是虚假的（且和它一致的也是虚假的）
11. 在1是真实的情况下，有可能找出部分虚假的
12. 然后再从21开始，，，，一共是分了1000 / 20 = 50组
13. 每组问19次（2.。。。20）所以刚好是 950次
14. 但是存在一种情况，还没有想清楚,就是虚假的，刚好在一个分组内
15. 比如就在第一个分组内，和1相同的等于9，但是恰好是另外10个是虚假的
16. 这时候就麻烦了
17. （1， 2), .... (1, 20), (1, 21)
18. 20次查询，如果相同的 <= 9（那么1和它相同的，肯定是虚假的）
19. 否则的话，和1相同的，至少是11
20. 那么这样子一共查询了 1, 22, 43, 64....
21. 1000/21 * 21 + 1
22. 48 * 21 = 48 + 960 = 2008
23. 47 * 21 = 47 + 940 = 987
24. 47 * 20 + 13 = 940 + 13 = 953