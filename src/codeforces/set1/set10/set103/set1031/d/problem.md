You are given a matrix of size 𝑛×𝑛
filled with lowercase English letters. You can change no more than 𝑘
letters in this matrix.

Consider all paths from the upper left corner to the lower right corner that move from a cell to its neighboring cell to
the right or down. Each path is associated with the string that is formed by all the letters in the cells the path
visits. Thus, the length of each string is 2𝑛−1
.

Find the lexicographically smallest string that can be associated with a path after changing letters in at most 𝑘
cells of the matrix.

A string 𝑎
is lexicographically smaller than a string 𝑏
, if the first different letter in 𝑎
and 𝑏
is smaller in 𝑎
.

### ideas

1. 如果 k >= 2n - 1 => all a
2. k < 2n - 1
3. if k = 0, 就是要在每个斜线上找到最小值，然后再找最小值这样
4. 假设 k > 0, 那么就应该在第一个地方要变成'a'， 在它变成'a'的地方，它其实可以到达这条斜线上所有的位置
5. 如果这条线上存在'a'， 那么就不要使用k，留下一个会是更好的选择
6. 也就是说，理论上s的前k个字符都是a，如果在里面存在a，继续保留完后，知道第一个不是a，且没有剩余k的地方