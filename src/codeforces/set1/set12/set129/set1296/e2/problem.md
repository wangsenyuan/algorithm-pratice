You are given a string 𝑠
 consisting of 𝑛
 lowercase Latin letters.

You have to color all its characters the minimum number of colors (each character to exactly one color, the same letters can be colored the same or different colors, i.e. you can choose exactly one color for each index in 𝑠
).

After coloring, you can swap any two neighboring characters of the string that are colored different colors. You can perform such an operation arbitrary (possibly, zero) number of times.

The goal is to make the string sorted, i.e. all characters should be in alphabetical order.

Your task is to find the minimum number of colors which you have to color the given string in so that after coloring it can become sorted by some sequence of swaps. Note that you have to restore only coloring, not the sequence of swaps.

### ideas
1. 干死脑细胞的题目
2. 结果string是知道的
3. x = sorted(s)
4. 如果s[i] 要移动到位置j, 那么它必须在[j...i]中间可以被移动
5. 所以这个中间的color，都不能和i的color一致（否则它就移动不了）包括位置j的颜色
6. 