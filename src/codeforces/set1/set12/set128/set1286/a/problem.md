Vadim loves decorating the Christmas tree, so he got a beautiful garland as a present. It consists of 𝑛
 light bulbs in a single row. Each bulb has a number from 1
 to 𝑛
 (in arbitrary order), such that all the numbers are distinct. While Vadim was solving problems, his home Carp removed some light bulbs from the garland. Now Vadim wants to put them back on.


Vadim wants to put all bulb back on the garland. Vadim defines complexity of a garland to be the number of pairs of adjacent bulbs with numbers with different parity (remainder of the division by 2
). For example, the complexity of 1 4 2 3 5 is 2
 and the complexity of 1 3 5 7 6 4 2 is 1
.

No one likes complexity, so Vadim wants to minimize the number of such pairs. Find the way to put all bulbs back on the garland, such that the complexity is as small as possible.

### ideas
1. dp