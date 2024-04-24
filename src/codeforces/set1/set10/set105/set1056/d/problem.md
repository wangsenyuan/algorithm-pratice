There is one apple tree in Arkady's garden. It can be represented as a set of junctions connected with branches so that
there is only one way to reach any junctions from any other one using branches. The junctions are enumerated from 1
to 𝑛
, the junction 1
is called the root.

A subtree of a junction 𝑣
is a set of junctions 𝑢
such that the path from 𝑢
to the root must pass through 𝑣
. Note that 𝑣
itself is included in a subtree of 𝑣
.

A leaf is such a junction that its subtree contains exactly one junction.

The New Year is coming, so Arkady wants to decorate the tree. He will put a light bulb of some color on each leaf
junction and then count the number happy junctions. A happy junction is such a junction 𝑡
that all light bulbs in the subtree of 𝑡
have different colors.

Arkady is interested in the following question: for each 𝑘
from 1
to 𝑛
, what is the minimum number of different colors needed to make the number of happy junctions be greater than or equal
to 𝑘
?

### ideas

1. 问题是，对于k，需要的最小的不同的color，使得happy的子树的个数 >= k
2. 对于子树x，如果它的leaf的count = w，那么要它happy，至少需要w个color