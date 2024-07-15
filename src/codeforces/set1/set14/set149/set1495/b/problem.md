On a weekend, Qingshan suggests that she and her friend Daniel go hiking. Unfortunately, they are busy high school students, so they can only go hiking on scratch paper.

A permutation 𝑝
 is written from left to right on the paper. First Qingshan chooses an integer index 𝑥
 (1≤𝑥≤𝑛
) and tells it to Daniel. After that, Daniel chooses another integer index 𝑦
 (1≤𝑦≤𝑛
, 𝑦≠𝑥
).

The game progresses turn by turn and as usual, Qingshan moves first. The rules follow:

If it is Qingshan's turn, Qingshan must change 𝑥
 to such an index 𝑥′
 that 1≤𝑥′≤𝑛
, |𝑥′−𝑥|=1
, 𝑥′≠𝑦
, and 𝑝𝑥′<𝑝𝑥
 at the same time.
If it is Daniel's turn, Daniel must change 𝑦
 to such an index 𝑦′
 that 1≤𝑦′≤𝑛
, |𝑦′−𝑦|=1
, 𝑦′≠𝑥
, and 𝑝𝑦′>𝑝𝑦
 at the same time.
The person who can't make her or his move loses, and the other wins. You, as Qingshan's fan, are asked to calculate the number of possible 𝑥
 to make Qingshan win in the case both players play optimally.

Input
The first line contains a single integer 𝑛
 (2≤𝑛≤105
) — the length of the permutation.

The second line contains 𝑛
 distinct integers 𝑝1,𝑝2,…,𝑝𝑛
 (1≤𝑝𝑖≤𝑛
) — the permutation.

### ideas
1. 先理解问题, x -> x1, 必须 abs(x1 - x) = 1, and p[x1] < p[x]
2. x是下标，alice新的位置，是左边或者右边（满足相差1），且新位置 p[x1] < p[x], 且不能是bob选的位置
3. 对于例子 [1 2 5 4 3], 假设 alice 选择 x= 3, bob 选择 y = 4
4. 那么 alice let x = 2, 满足条件，距离1，且 p[2] < p[3]
5. bob选择 y = 3, 相差距离 = 1， 且 p[y1] > p[y]
6. alice 选择 x = 1, bob can't move (因为p[2], p[4] < p[3])
7. 如果 bob选择 y = 5, 那么alice第一步就会走到 x = 4 的位置， bob lose
8. 如果没有限制的情况下， alice是从山顶向山谷走， bob从山谷向山顶走
9. alice肯定选择山顶，假设不是，如果他选择的是山谷，很简单，bob选任何位置都可以，alice第一步就无法行动
10. 如果alice选择的是山坡，那么bob只要选择山坡下一个位置，那么alice也无法移动
11. 所以，反推出alice只能选择山顶
12. 假设alice选择了一个山顶，向两边分别移动x, y的距离
13. 如果bob要阻止alice。如果存在一个山谷，它的上行距离至少是max(x, y), alice没有赢的机会
14. 否则，这个山顶就是alice赢的一个地方
15. 如果alice选了一个山顶，bob肯定选这山顶对应的山谷，且是那个长的山谷