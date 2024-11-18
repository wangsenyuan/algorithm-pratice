Vova's family is building the Great Vova Wall (named by Vova himself). Vova's parents, grandparents, grand-grandparents
contributed to it. Now it's totally up to Vova to put the finishing touches.

The current state of the wall can be respresented by a sequence a
of n
integers, with ai
being the height of the i
-th part of the wall.

Vova can only use 2×1
bricks to put in the wall (he has infinite supply of them, however).

Vova can put bricks horizontally on the neighboring parts of the wall of equal height. It means that if for some i
the current height of part i
is the same as for part i+1
, then Vova can put a brick there and thus increase both heights by 1. Obviously, Vova can't put bricks in such a way
that its parts turn out to be off the borders (to the left of part 1
of the wall or to the right of part n
of it).

The next paragraph is specific to the version 1 of the problem.

Vova can also put bricks vertically. That means increasing height of any part of the wall by 2.

Vova is a perfectionist, so he considers the wall completed when:

all parts of the wall has the same height;
the wall has no empty spaces inside it.
Can Vova complete the wall using any amount of bricks (possibly zero)?

### ideas

1. 和d2不同的是，这次可以使用垂直的砖块，那么假设最后的高度是h，那么(h * w - sum) % 2 = 0
2. 只要满足这个条件就可以，
3. 如果sum是偶数，那么 h * w 也必须是偶数，如果w是偶数，那么肯定可以
4. 如果w是奇数，只要h是偶数，也是可以的
5. 也就是说如果sum是偶数，一定ok
6. 如果sum是奇数，那么h*w必须是奇数，如果h和w都是奇数（可以得到奇数）
7. 所以只要w是偶数，就不行 （因为h未定）
8. 如果sum是偶数，那么h*w必须是偶数（都可以，w时奇数/偶数，都ok）