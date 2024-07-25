On a number axis directed from the left rightwards, n marbles with coordinates x1, x2, ..., xn are situated. Let's assume that the sizes of the marbles are infinitely small, that is in this task each of them is assumed to be a material point. You can stick pins in some of them and the cost of sticking in the marble number i is equal to ci, number ci may be negative. After you choose and stick the pins you need, the marbles will start to roll left according to the rule: if a marble has a pin stuck in it, then the marble doesn't move, otherwise the marble rolls all the way up to the next marble which has a pin stuck in it and stops moving there. If there is no pinned marble on the left to the given unpinned one, it is concluded that the marble rolls to the left to infinity and you will pay an infinitely large fine for it. If no marble rolled infinitely to the left, then the fine will consist of two summands:

the sum of the costs of stuck pins;
the sum of the lengths of the paths of each of the marbles, that is the sum of absolute values of differences between their initial and final positions.
Your task is to choose and pin some marbles in the way that will make the fine for you to pay as little as possible.


### ideas
1. 显然这个是分了很多段
2. 每段的成本 = c[i] + sum(x[j] - x[i]) i是这段的起点