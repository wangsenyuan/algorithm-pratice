The array 𝑎1,𝑎2,…,𝑎𝑚
is initially filled with zeroes. You are given 𝑛
pairwise distinct segments 1≤𝑙𝑖≤𝑟𝑖≤𝑚
. You have to select an arbitrary subset of these segments (in particular, you may select an empty set). Next, you do
the following:

For each 𝑖=1,2,…,𝑛
, if the segment (𝑙𝑖,𝑟𝑖)
has been selected to the subset, then for each index 𝑙𝑖≤𝑗≤𝑟𝑖
you increase 𝑎𝑗
by 1
(i. e. 𝑎𝑗
is replaced by 𝑎𝑗+1
). If the segment (𝑙𝑖,𝑟𝑖)
has not been selected, the array does not change.
Next (after processing all values of 𝑖=1,2,…,𝑛
), you compute max(𝑎)
as the maximum value among all elements of 𝑎
. Analogously, compute min(𝑎)
as the minimum value.
Finally, the cost of the selected subset of segments is declared as max(𝑎)−min(𝑎)
.
Please, find the maximum cost among all subsets of segments.

### thoughts

1. 最优的答案，是不是存在令 min(a) = 0的情况?
2. 考虑min(a) = x, 是某个位置得到的，那么必然有x段覆盖了它
3. 那么这几段存在两种情况，一种是它们为max(a)没有贡献
4. 这类可以直接删除，从而得到更优的结果
5. 一种是它们对max(a)有贡献，那么移除掉它们后，不会产生更差的结果，
6. 那么就是找到某个位置前面，或者后面的最大值即可