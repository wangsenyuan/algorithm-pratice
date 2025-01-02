The legend of the foundation of Vectorland talks of two integers 𝑥
 and 𝑦
. Centuries ago, the array king placed two markers at points |𝑥|
 and |𝑦|
 on the number line and conquered all the land in between (including the endpoints), which he declared to be Arrayland. Many years later, the vector king placed markers at points |𝑥−𝑦|
 and |𝑥+𝑦|
 and conquered all the land in between (including the endpoints), which he declared to be Vectorland. He did so in such a way that the land of Arrayland was completely inside (including the endpoints) the land of Vectorland.

Here |𝑧|
 denotes the absolute value of 𝑧
.

Now, Jose is stuck on a question of his history exam: "What are the values of 𝑥
 and 𝑦
?" Jose doesn't know the answer, but he believes he has narrowed the possible answers down to 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
. Now, he wants to know the number of unordered pairs formed by two different elements from these 𝑛
 integers such that the legend could be true if 𝑥
 and 𝑦
 were equal to these two values. Note that it is possible that Jose is wrong, and that no pairs could possibly make the legend true.


### ideas
1. 假设 x > 0, y > 0
2. y + x > y 没有问题
3. y - x < x => y < 2 * x要成立
4. 也就是说，那么这个在 x < 0, y > 0 时呢？
5. abs(x) <= y
6. abs(x) .... y
7. -x < y
8. y-x > y > 0
9. abs(y + x) = y + x
10. y + x <= -x => y <= -2 *x
11. y >= -x and y <= -2 * x
12. 如果 abs(x) > y呢
13. -x > y => y + x < 0, y - x > -x (这个肯定是右端点)
14. y+x < 0 abs(y+x) = -x - y < y
15. -x < 2 * y
16. 所以，感觉应该变成绝对值，然后按照正数处理？