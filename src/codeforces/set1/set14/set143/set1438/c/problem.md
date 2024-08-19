Artem is building a new robot. He has a matrix 𝑎
 consisting of 𝑛
 rows and 𝑚
 columns. The cell located on the 𝑖
-th row from the top and the 𝑗
-th column from the left has a value 𝑎𝑖,𝑗
 written in it.

If two adjacent cells contain the same value, the robot will break. A matrix is called good if no two adjacent cells contain the same value, where two cells are called adjacent if they share a side.

Artem wants to increment the values in some cells by one to make 𝑎
 good.

More formally, find a good matrix 𝑏
 that satisfies the following condition —

For all valid (𝑖,𝑗
), either 𝑏𝑖,𝑗=𝑎𝑖,𝑗
 or 𝑏𝑖,𝑗=𝑎𝑖,𝑗+1
.
For the constraints of this problem, it can be shown that such a matrix 𝑏
 always exists. If there are several such tables, you can output any of them. Please note that you do not have to minimize the number of increments.

 ### ideas
 1. 这里有两个冲突点，一个是在改变前， a[i][j] = a[i+1][j] ...
 2. 另外一个是改变后, a[i][j] = b[i+1][j]
 3. 如果是改造前，两个是相同的，那么它们必须是不同的颜色（不能同时改变）
 4. 如果是改造后，两个值是相同的，那么它们必须是相同的颜色（必须同时改变）
 5. 所以在相同的两个cell间连线，然后进行着色
 6. 如果出现了奇数长度的环 => false
 7. 似乎还是缺点东西，上面的过程是把值相同的连在了一起，然后隔一位增加1
 8. 但是没法很好的处理，增加1后，和临近数相同的情况
 9. 而且选择哪一些进行变大，对结果会有不一样的影响
 10. 基本上，还是把那些不变的和变化的，分成两类
 11. 相同值之间连接线，但是在增加后，也会出现相同值的，要怎么处理呢？
 12. 是不是在它们中间连接一个虚拟的点
 13. 但是还是很难处理啊。假设在u ->(w) -> v 但是如果u没有变化，那么它对v是没有影响的
 14. 只有当u变化的时候，需要v也同时变化
 15. 假设红色表示要+1, 那么只有当u是红色色，v才需要也是红色（+1）
 16. 当u是蓝色时，v可以是红色，也可以是蓝色
 17. u => v
 18. 也就是说要在u和v间建立一个u => v 的路径（u为true，v也为true）(u + 1 = v)
 19. 在 u => ~v and ~u => v的路径（表示u为true时，v为false， 或者u为false， v为true）的路径（不能同时为true)
 20. 然后看是否存在冲突