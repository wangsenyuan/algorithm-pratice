During her tantrums the princess usually smashes some collectable porcelain. Every furious shriek is accompanied with one item smashed.

The collection of porcelain is arranged neatly on n shelves. Within each shelf the items are placed in one row, so that one can access only the outermost items — the leftmost or the rightmost item, not the ones in the middle of the shelf. Once an item is taken, the next item on that side of the shelf can be accessed (see example). Once an item is taken, it can't be returned to the shelves.

You are given the values of all items. Your task is to find the maximal damage the princess' tantrum of m shrieks can inflict on the collection of porcelain.

### ideas
1. dp[i][j] = 在第i行，使用j次操作，所能达到的最大value