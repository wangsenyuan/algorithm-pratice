You are given an array 𝑎
of 𝑛
positive integers. In one operation, you must pick some (𝑖,𝑗)
such that 1≤𝑖<𝑗≤|𝑎|
and append |𝑎𝑖−𝑎𝑗|
to the end of the 𝑎
(i.e. increase 𝑛
by 1
and set 𝑎𝑛
to |𝑎𝑖−𝑎𝑗|
). Your task is to minimize and print the minimum value of 𝑎
after performing 𝑘
operations.

### thoughts

1. 如果一个数字不用来计算最终结果，就不用处理它
2. 先排序，
3. 没想法啊～～
4. 如果只有一次操作，那必然是相邻的两个数进行处理
5. 如果 k > 2, 那么就是0
6. 因为第二次操作可以选和第一次操作一样的数，那么再一次操作就是0
7. got~
8. k = 2 时，有点不大对～
9. 