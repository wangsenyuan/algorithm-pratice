Vova's family is building the Great Vova Wall (named by Vova himself). Vova's parents, grandparents, grand-grandparents
contributed to it. Now it's totally up to Vova to put the finishing touches.

The current state of the wall can be respresented by a sequence 𝑎
of 𝑛
integers, with 𝑎𝑖
being the height of the 𝑖
-th part of the wall.

Vova can only use 2×1
bricks to put in the wall (he has infinite supply of them, however).

Vova can put bricks only horizontally on the neighbouring parts of the wall of equal height. It means that if for some
𝑖
the current height of part 𝑖
is the same as for part 𝑖+1
, then Vova can put a brick there and thus increase both heights by 1. Obviously, Vova can't put bricks in such a way
that its parts turn out to be off the borders (to the left of part 1
of the wall or to the right of part 𝑛
of it).

Note that Vova can't put bricks vertically.

Vova is a perfectionist, so he considers the wall completed when:

all parts of the wall has the same height;
the wall has no empty spaces inside it.
Can Vova complete the wall using any amount of bricks (possibly zero)?

### ideas

1. 先不考虑性能，有个自然的想法，从最低的点开始，如果这个位置长度为偶数，那么就可以加上去
2. 直到遇到奇数长度的间隔，或者是到顶
3. stack？
