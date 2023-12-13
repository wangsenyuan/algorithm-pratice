You are playing with a really long strip consisting of 1018
white cells, numbered from left to right as 0
, 1
, 2
and so on. You are controlling a special pointer that is initially in cell 0
. Also, you have a "Shift" button you can press and hold.

In one move, you can do one of three actions:

move the pointer to the right (from cell 𝑥
to cell 𝑥+1
);
press and hold the "Shift" button;
release the "Shift" button: the moment you release "Shift", all cells that were visited while "Shift" was pressed are
colored in black.
(Of course, you can't press Shift if you already hold it. Similarly, you can't release Shift if you haven't pressed it.)
Your goal is to color at least 𝑘
cells, but there is a restriction: you are given 𝑛
segments [𝑙𝑖,𝑟𝑖]
— you can color cells only inside these segments, i. e. you can color the cell 𝑥
if and only if 𝑙𝑖≤𝑥≤𝑟𝑖
for some 𝑖
.

What is the minimum number of moves you need to make in order to color at least 𝑘
cells black?

### thoughts

1. 假设移动了x的距离，那么在x的左侧的区间（包括包含x的区间）应该是最大的那些被使用，且它们的sum >= k
4. 每使用一个区间，会增加2，所以在当前区间来看，所有长度 < 2 的部分，都可以被当前移动距离为1的给替换掉

### solution

Let's look at what's happening in the task: in the end the pointer will move into some position 𝑝
and some segments on the prefix [0,𝑝]
will be colored. Note that it's optimal to stop only inside some segment (or 𝑙𝑖≤𝑝≤𝑟𝑖
for some 𝑖
), and if we colored 𝑥
segments (including the last segment [𝑙𝑖,𝑝]
that may be colored partially) the answer will be equal to 𝑝+2⋅𝑥
.

Let's prove that it's not optimal to skip segments with length 𝑙𝑒𝑛=𝑟[𝑖]−𝑙[𝑖]+1≥2
. By contradiction, suppose the optimal answer 𝑎
has a skipped segment [𝑙𝑖,𝑟𝑖]
. If we color that segment instead, we will make 2
more moves for pressing and releasing Shift, but we can make at least 𝑙𝑒𝑛
right move less. So the new answer 𝑎′=𝑎+2−𝑙𝑒𝑛≤𝑎
— the contradiction.

Now we are ready to write a solution. Let's iterate over 𝑖
— the last segment we will color (and therefore the segment where we stop). At first, let's imagine we colored the whole
segment [𝑙𝑖,𝑟𝑖]
as well. Let 𝑠
be the total length of all segments on prefix [0,𝑟𝑖]
that are longer than 1
and 𝑐
be the number of segments of length 1
on this prefix.

There are three cases:

if 𝑠+𝑐<𝑘
, stopping inside the 𝑖
-th segment is not enough;
if 𝑠<𝑘
but 𝑠+𝑐≥𝑘
, we will color all "long" segments plus several short ones. The current answer will be equal to 𝑟𝑖+2((𝑖−𝑐)+(𝑘−𝑠))
, where 𝑟𝑖
is where we stop, (𝑖−𝑐)
is the number of long segments and (𝑘−𝑠)
is the number of short segments we need;
if 𝑠≥𝑘
, then we don't need any short segments. More over, we can stop even before reaching 𝑟𝑖
. So, the current answer will be equal to 𝑟𝑖−(𝑠−𝑘)+2(𝑖−𝑐)
, where 𝑟𝑖−(𝑠−𝑘)
is the exact position to stop to get exactly 𝑘
black cells and (𝑖−𝑐)
is the number of long segments.
Note that 𝑖
is 1
-indexed in calculations above, and we can stop the first moment we met the situation 𝑠≥𝑘
.
The answer is the minimum among the answers we've got in the process. Since it's easy to update values 𝑠
and 𝑐
when we move from 𝑖
to 𝑖+1
, the total complexity is 𝑂(𝑛)
.