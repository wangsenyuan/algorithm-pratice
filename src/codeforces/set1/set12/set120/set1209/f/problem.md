Koala Land consists of 𝑚
bidirectional roads connecting 𝑛
cities. The roads are numbered from 1
to 𝑚
by order in input. It is guaranteed, that one can reach any city from every other city.

Koala starts traveling from city 1
. Whenever he travels on a road, he writes its number down in his notebook. He doesn't put spaces between the numbers,
so they all get concatenated into a single number.

Before embarking on his trip, Koala is curious about the resulting number for all possible destinations. For each
possible destination, what is the smallest number he could have written for it?

Since these numbers may be quite large, print their remainders modulo 109+7
. Please note, that you need to compute the remainder of the minimum possible number, not the minimum possible
remainder.

### ideas

1. 理解错了～～～
2. 假设到目前为止，在访问x，那么对于所有后续的节点，x都会出现在相同的位置
3. 又理解错了，边权 = 它的编号
4. 