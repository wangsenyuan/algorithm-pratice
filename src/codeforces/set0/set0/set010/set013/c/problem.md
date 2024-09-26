Little Petya likes to play very much. And most of all he likes to play the following game:

He is given a sequence of N integer numbers. At each step it is allowed to increase the value of any number by 1 or to decrease it by 1. The goal of the game is to make the sequence non-decreasing with the smallest number of steps. Petya is not good at math, so he asks for your help.

The sequence a is called non-decreasing if a1 ≤ a2 ≤ ... ≤ aN holds, where N is the length of the sequence.

### ideas
1. 左端正常应该是要减，右端正常应该是加
2. n <= 5000 (所以一个n * n 的算法，应该也可以搞定)
3. 这个序列中，是不是一定有一个数是不变的？
4. 好像是存在的，把这些线连起来，然后把最后的结果连起来
5. 这两条线肯定有一个交点，这个交点如果在整数位置，那么就是这个位置
6. 如果不在，假设这个焦点在某个点的右侧，那么它左侧的相当于减小了，它右侧的要增加
7. 那么这时不妨就让左、右两测都不变，会得到更好的结果
8. 上面的分析时不对的
9. 