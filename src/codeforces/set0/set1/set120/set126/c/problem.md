After years of hard work scientists invented an absolutely new e-reader display. The new display has a larger resolution, consumes less energy and its production is cheaper. And besides, one can bend it. The only inconvenience is highly unusual management. For that very reason the developers decided to leave the e-readers' software to programmers.

The display is represented by n × n square of pixels, each of which can be either black or white. The display rows are numbered with integers from 1 to n upside down, the columns are numbered with integers from 1 to n from the left to the right. The display can perform commands like "x, y". When a traditional display fulfills such command, it simply inverts a color of (x, y), where x is the row number and y is the column number. But in our new display every pixel that belongs to at least one of the segments (x, x) - (x, y) and (y, y) - (x, y) (both ends of both segments are included) inverts a color.

For example, if initially a display 5 × 5 in size is absolutely white, then the sequence of commands (1, 4), (3, 5), (5, 1), (3, 3) leads to the following changes:


You are an e-reader software programmer and you should calculate minimal number of commands needed to display the picture. You can regard all display pixels as initially white.


### ideas
1. 完全没啥想法～
2. 如果位置(x, y) 被涂成了黑色
3. 有多种选择，假设某个(x, y)后面的点(x, y0) 作为中点
4. 或者某个和y同一列的点(x0, y)作为中点
5. 两个中点 (x1, y1), (x2, y2)
6. 那么它们的重叠点的颜色（是白色）
7. 感觉要从右上到左下处理，如果某个地方和它的预期颜色不一样，就必须在这里操作一下，假设是(x, y)
8. 那么从(x, x) .... (x, y), (y, y) 到 (x, y) 
9. 但是这个貌似只对右上的有效，对左下没有效果
10. 所以要从外到内的处理，这样子之前处理的结果，会对现在产生影响，但是现在处理的结果，不会对之前的处理造成影响
11. 似乎也不用，右上的不会越过中轴线，左下的也不会。
12. 所以单独处理就可以了
13. 中轴线单独处理即可