In the last mission, MDCS has successfully shipped 𝑁
 AI robots to Mars. Before they start exploring, system initialization is required so they are arranged in a line. Every robot can be described with three numbers: position (𝑥𝑖
), radius of sight (𝑟𝑖
) and IQ (𝑞𝑖
).

Since they are intelligent robots, some of them will talk if they see each other. Radius of sight is inclusive, so robot can see other all robots in range [𝑥𝑖−𝑟𝑖,𝑥𝑖+𝑟𝑖]
. But they don't walk to talk with anybody, but only with robots who have similar IQ. By similar IQ we mean that their absolute difference isn't more than 𝐾
.

Help us and calculate how many pairs of robots are going to talk with each other, so we can timely update their software and avoid any potential quarrel.


### ideas
1. sort by x first
2. 只有两个robot能相互看到时，才会talk(且IQ不超过K)
3. 相互看到这个比较好处理，（每个robot都只关心左边的，这样可以避免重复）
4. i能看到的访问是 x[i] - r[i], 那么所有此时都active的，都可以看到
5. active定义为，robot j的视野还包含x[i]
6. 现在的问题就是IQ的差，怎么处理？
7. 貌似要用AVL树了。用IQ作为key
8. 按照x排序不对，对于例子中的情况，虽然1号是active的，它可以看到2号，但是2号看不到它
9. 应该按照iq排序，只需要处理iq - k范围内，超过这个访问的，可以直接删除掉
10. 然后怎么处理是个问题
11. 假设i，j能够相互看到, 且 x[i] < x[j]
12. x[j] <= x[i] + r[i]
13. x[i] >= x[j] - r[j]
14. x[j] <= x[i] + r[j]
15. x[j] <= min(x[i] + r[i], x[i] + r[j]) = x[i] + min(r[i], r[j])
16. x[j] - x[i] <= min(r[i], r[j]) (两点的距离 <= 最短的半径)
17. 按照半径降序处理？当前的r肯定是最小的，那么在它的访问内的，肯定就可以相互看到