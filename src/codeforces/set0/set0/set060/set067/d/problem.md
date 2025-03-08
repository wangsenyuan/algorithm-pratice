Professor Phunsuk Wangdu has performed some experiments on rays. The setup for n rays is as follows.

There is a rectangular box having exactly n holes on the opposite faces. All rays enter from the holes of the first side and exit from the holes of the other side of the box. Exactly one ray can enter or exit from each hole. The holes are in a straight line.


Professor Wangdu is showing his experiment to his students. He shows that there are cases, when all the rays are intersected by every other ray. A curious student asked the professor: "Sir, there are some groups of rays such that all rays in that group intersect every other ray in that group. Can we determine the number of rays in the largest of such groups?".

Professor Wangdu now is in trouble and knowing your intellect he asks you to help him.

### ideas
1. 上面从左往右处理，当前为i，那么和i能相交的只能前面的，且它们在下面一排，在i的后面
2. 这些比较好找，但是，还需要的一个特性是，这些都必须相互交叉
3. 假设知道某个j相互交叉的数量, f(j)
4. f(j) = j前面，且和j交叉，且它们都相互交叉的数量
5. j在下排中，肯定是最前面那个
6. 如果i和j相交，那么i肯定和g(j)相交
7. f(i) = f(j) + 1
8. 也就是在i的前面找到那些同时在i的后面的j，它们的f(j)中的最大值
9. 