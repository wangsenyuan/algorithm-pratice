Iahub helps his grandfather at the farm. Today he must milk the cows. There are n cows sitting in a row, numbered from 1
to n from left to right. Each cow is either facing to the left or facing to the right. When Iahub milks a cow, all the
cows that see the current cow get scared and lose one unit of the quantity of milk that they can give. A cow facing left
sees all the cows with lower indices than her index, and a cow facing right sees all the cows with higher indices than
her index. A cow that got scared once can get scared again (and lose one more unit of milk). A cow that has been milked
once cannot get scared and lose any more milk. You can assume that a cow never loses all the milk she can give (a cow
gives an infinitely amount of milk).

Iahub can decide the order in which he milks the cows. But he must milk each cow exactly once. Iahub wants to lose as
little milk as possible. Print the minimum amount of milk that is lost.

### ideas

1. 考虑两个位置i, j, i < j
2. 如果a[i] = 1, a[j] = 1, 也就是它们都朝向右边，如果先操作i,对j没有影响，反过来，j会影响到i
3. 如果a[i] = 1, a[j] = 0, 它们相互看着，无论怎样，都会影响到
4. 如果a[i] = 0, a[j] = 1, 无论怎样，都不会影响到
5. 如果a[i] = 0, a[j] = 0, 它们朝向左边，应该先处理j
6. 综合，就是a[i] = a[j]时，需要特殊处理，否则不用处理