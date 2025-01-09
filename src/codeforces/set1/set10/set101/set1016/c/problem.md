Vasya's house is situated in a forest, and there is a mushroom glade near it. The glade consists of two rows, each of which can be divided into n consecutive cells. For each cell Vasya knows how fast the mushrooms grow in this cell (more formally, how many grams of mushrooms grow in this cell each minute). Vasya spends exactly one minute to move to some adjacent cell. Vasya cannot leave the glade. Two cells are considered adjacent if they share a common side. When Vasya enters some cell, he instantly collects all the mushrooms growing there.

Vasya begins his journey in the left upper cell. Every minute Vasya must move to some adjacent cell, he cannot wait for the mushrooms to grow. He wants to visit all the cells exactly once and maximize the total weight of the collected mushrooms. Initially, all mushrooms have a weight of 0. Note that Vasya doesn't need to return to the starting cell.

Help Vasya! Calculate the maximum total weight of mushrooms he can collect.

### ideas
1. 每个都必须访问一次（exactly once）
2. 所以可以分成两部分左边、右边
3. 左边必须是锯齿形，右边是长条形
4. 确定了分界线，就确定了每个cell的访问时刻
5. 但是每次重新计算右边的值，肯定会超时，所以要好好的想一下
6. 如果当前的位置是在 (1, i), 且从它开始往右，最后再回到(2, i)
7. 左边已经走过花费的时间 s = (i-1) * 2
8. s * a(1, i) + (s + 1) * a(1, i + 1) ..
9. s * (a(1, i) + a(1, i+1) + ...) + a(1, i + 1) + 2 * a(1, i + 2) ...
10. 让 x[1] = i * a[i]，然后在往右边移动的时候，不断的减去suf[i], 似乎能够达到快速计算的目的
11. 假设从第一行开始，已经计算了 i * a[i]的和，现在因为