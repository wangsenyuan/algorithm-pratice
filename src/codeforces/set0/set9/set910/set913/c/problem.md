A New Year party is not a New Year party without lemonade! As usual, you are expecting a lot of guests, and buying
lemonade has already become a pleasant necessity.

Your favorite store sells lemonade in bottles of n different volumes at different costs. A single bottle of type i has
volume 2i - 1 liters and costs ci roubles. The number of bottles of each type in the store can be considered infinite.

You want to buy at least L liters of lemonade. How many roubles do you have to spend?

### ideas

1. 一个type i的相当于两个type i - 1
2. 也就是说，如果要买一个完整的type i， 那么相当于购买两个type i - 1,
3. 所以如果 c[i] > 2 * c[i-1] 那么，可以根据上面的推论，让c[i] = 2 * c[i-1]
4. 然后就可以从高往低的处理

### solution

Note that if 2·ai ≤ ai + 1, then it doesn't make sense to buy any bottles of type i + 1 — it won't ever be worse to buy
two bottles of type i instead. In this case let's assume that we actually have an option to buy a bottle of type i + 1
at the cost of 2·ai and replace ai + 1 with min(ai + 1, 2·ai). Let's do this for all i from 1 to n - 1 in increasing
order.

Now for all i it's true that 2·ai ≥ ai + 1. Note that now it doesn't make sense to buy more than one bottle of type i if
i<n. Indeed, in this case it won't ever be worse to buy a bottle of type i + 1 instead of two bottles of type i. From
now on, we'll only search for options where we buy at most one bottle of every type except the last one.

Suppose that we had to buy exactly L liters of lemonade, as opposed to at least L. Note that in this case the last n - 1
bits of L uniquely determine which bottles of types less than n we have to buy. Indeed, if L is odd, then we have to buy
a bottle of type 0, otherwise we can't do that. By the same line of thought, it's easy to see that bit j in the binary
representation of L is responsible for whether we should buy a bottle of type j. Finally, we have to buy exactly ⌊ L /
2n - 1⌋ bottles of type n.

But what to do with the fact that we're allowed to buy more than L liters? Suppose we buy M>L liters. Consider the
highest bit j in which M and L differ. Since M>L, the j-th bit in M is 1, and the j-th bit in L is 0. But then all bits
lower than the j-th in M are 0 in the optimal answer, since these bits are responsible for the "extra" bottles — those
for which we spend money for some reason, but without which we would still have M>L.

Thus, here is the overall solution. Loop over the highest bit j in which M differs from L. Form the value of M, taking
bits higher than the j-th from L, setting the j-th bit in M to 1, and bits lower than the j-th to 0. Calculate the
amount of money we have to pay to buy exactly M liters of lemonade. Take the minimum over all j.

The complexity of the solution is O(n) or O(n2), depending on the implementation.