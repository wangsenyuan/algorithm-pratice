There are 𝑛
 slimes placed in a line. The slimes are numbered from 1
 to 𝑛
 in order from left to right. The size of the 𝑖
-th slime is 𝑎𝑖
.

Every second, the following happens: exactly one slime eats one of its neighbors and increases its size by the eaten neighbor's size. A slime can eat its neighbor only if it is strictly bigger than this neighbor. If there is no slime which is strictly bigger than one of its neighbors, the process ends.

For example, suppose 𝑛=5
, 𝑎=[2,2,3,1,4]
. The process can go as follows:

first, the 3
-rd slime eats the 2
-nd slime. The size of the 3
-rd slime becomes 5
, the 2
-nd slime is eaten.
then, the 3
-rd slime eats the 1
-st slime (they are neighbors since the 2
-nd slime is already eaten). The size of the 3
-rd slime becomes 7
, the 1
-st slime is eaten.
then, the 5
-th slime eats the 4
-th slime. The size of the 5
-th slime becomes 5
, the 4
-th slime is eaten.
then, the 3
-rd slime eats the 5
-th slime (they are neighbors since the 4
-th slime is already eaten). The size of the 3
-rd slime becomes 12
, the 5
-th slime is eaten.
For each slime, calculate the minimum number of seconds it takes for this slime to be eaten by another slime (among all possible ways the process can go), or report that it is impossible.

### ideas
1. 如果一开始大家都是一样a，所有的答案都是-1
2. 如果不是1样大，这个过程似乎似乎始终能够完成
3. 但是对于一个具体的i，它却不一定能够被吃点
4. 考虑i被前面的吃点，假设i-1 > a[i]那么可以直接完成
5. 否则 a[i-1] <= a[i], 如果 a[i-2] != a[i-1]，那么就可以合并它们
6. 直到要么 sum > a[i]， 要么和a[j] == sum
7. 但是即使 a[j] = sum, 其实也可以从a[j]放过来吃到i-1
8. 所以，似乎只有一种情况 i = 2, 且 a[0] = a[1], 那么它们无法对i来吃
9. 否则只要前面size 超过3，始终有办法把前面给组合起来 