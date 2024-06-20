You have a one-dimensional puzzle, all the elements of which need to be put in one row, connecting with each other. All the puzzle elements are completely white and distinguishable from each other only if they have different shapes.

Each element has straight borders at the top and bottom, and on the left and right it has connections, each of which can be a protrusion or a recess. You cannot rotate the elements.

You can see that there are exactly 4
 types of elements. Two elements can be connected if the right connection of the left element is opposite to the left connection of the right element.

All possible types of elements.
The puzzle contains 𝑐1,𝑐2,𝑐3,𝑐4
 elements of each type. The puzzle is considered complete if you have managed to combine all elements into one long chain. You want to know how many ways this can be done.


 ### ideas
 1. 这里有多组关系, 
 2. 1 -> 2 = 4
 3. 1 -> 3 = 1
 4. 2 -> 1 = 3
 5. 2 -> 4 = 2
 6. 3 -> 2 = 2
 7. 3 -> 3 = 3
 8. 4 -> 1 = 1
 9.  4 -> 4 = 4
 10. 箭头 表示从左到右， = 表示生成的新的图形
 11. 数字小是可以模拟的 (a, b, c, d) 表示当前的状态，
 12. 但是这里都太大了。没法模拟。
 13. 所以还得想其他的办法
 14. 2和2不能直接链接，1和1也不能直接链接, 3和4不能链接
 15. 1 -> 3 .... 3 -> 2
 16. 2 -> 4 ..... 4 -> 1
 17. 好像只有这两种模式
 18. 1， 2， 1， 2, 然后1， 2中间必须是3， 2， 1中间必须是4