Dreamoon likes coloring cells very much.

There is a row of 𝑛
 cells. Initially, all cells are empty (don't contain any color). Cells are numbered from 1
 to 𝑛
.

You are given an integer 𝑚
 and 𝑚
 integers 𝑙1,𝑙2,…,𝑙𝑚
 (1≤𝑙𝑖≤𝑛
)

Dreamoon will perform 𝑚
 operations.

In 𝑖
-th operation, Dreamoon will choose a number 𝑝𝑖
 from range [1,𝑛−𝑙𝑖+1]
 (inclusive) and will paint all cells from 𝑝𝑖
 to 𝑝𝑖+𝑙𝑖−1
 (inclusive) in 𝑖
-th color. Note that cells may be colored more one than once, in this case, cell will have the color from the latest operation.

Dreamoon hopes that after these 𝑚
 operations, all colors will appear at least once and all cells will be colored. Please help Dreamoon to choose 𝑝𝑖
 in each operation to satisfy all constraints.

 ### ideas
 1. 每次操作，都是要color一段连续的长度为l的段，让这段颜色为i
 2. 所以，感觉应该先处理长的？还是先处理短的？感觉应该处理长的，长的那一段，只要保证两头有一个i的颜色即可，中间的可以被短的给替代掉
 3. 所以，按照l将许排列
 4. 保证目前开始的位置，不会被覆盖掉，
 5. 但是，要让所有的都被color到，还缺一点
 6. 搞错了，这个顺序是固定的，不能重新排序，假设i，i+1, 如果先处理（i+1)， 再处理i，且i在i+1的内部，那么i的颜色就被覆盖掉了
 7. 还是要重新考虑
 8. 这里还挺麻烦的，即要保证每个颜色都出现了，还需要保证，整个区间都被图画到了
 9. 