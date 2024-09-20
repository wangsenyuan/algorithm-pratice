You are given an integer 𝑛
. You choose 𝑛
 cells (𝑥1,𝑦1),(𝑥2,𝑦2),…,(𝑥𝑛,𝑦𝑛)
 in the grid 𝑛×𝑛
 where 1≤𝑥𝑖≤𝑛
 and 1≤𝑦𝑖≤𝑛
.

Let 
 be the set of distinct Manhattan distances between any pair of cells. Your task is to maximize the size of 
. Examples of sets and their construction are given in the notes.

If there exists more than one solution, you are allowed to output any.

Manhattan distance between cells (𝑥1,𝑦1)
 and (𝑥2,𝑦2)
 equals |𝑥1−𝑥2|+|𝑦1−𝑦2|
.

### ideas
1. 这个上限是多少， n个节点，那么最多可以有(n + 1) * n / 2 个不同的距离
2. n = 2是，显然是成立的
3. n = 3时， 3 * 4 / 2 = 6， 从示例看，应该是不成立的  
4. 肯定有一个特殊的构型可以达到最大值