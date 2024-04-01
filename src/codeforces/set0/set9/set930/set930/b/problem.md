Vasya and Kolya play a game with a string, using the following rules. Initially, Kolya creates a string s, consisting of
small English letters, and uniformly at random chooses an integer k from a segment [0, len(s) - 1]. He tells Vasya this
string s, and then shifts it k letters to the left, i. e. creates a new string t = sk + 1sk + 2... sns1s2... sk. Vasya
does not know the integer k nor the string t, but he wants to guess the integer k. To do this, he asks Kolya to tell him
the first letter of the new string, and then, after he sees it, open one more letter on some position, which Vasya can
choose.

Vasya understands, that he can't guarantee that he will win, but he wants to know the probability of winning, if he
plays optimally. He wants you to compute this probability.

Note that Vasya wants to know the value of k uniquely, it means, that if there are at least two cyclic shifts of s that
fit the information Vasya knowns, Vasya loses. Of course, at any moment of the game Vasya wants to maximize the
probability of his win.

### ideas

1. 看不懂～～～
2. 对于每一个字母(26个)，如果它在第一位，是否存在一个位置使得能够确定这个字母来自哪里？
3. 那么就需要看这些字母后面的位置，
4. 假设字母x的位置组成了arr， 对于其中的每个i, 是否存在一个l s[(i + l) % n] 都不相同
5. 首先如果arr的size > 26， 肯定是判断不出来的，因为即使存在这样的一个l，其中至少有两个字母是一样的
6. 所以arr 的size <= 26, 然后迭代l，如果存在，就是可以获胜的
7. 不是按字符的，要按照（i, j)来判断

### solution

Idea. Let's consider all possible c1 that will be first in t. Then, let's consider all possible numbers of second letter
that Vasya will ask about — this will be d. If pair of letters (c1, c2) occurs only once at d distance, than if c2 opens
second time, Vasya will be able to determine shift.

Solution. Let's loop through all letters at d distance from all c1 letters and for each symbol c2 we will calculate
number of such letters. This can be done in O(cnt(c1)), where cnt(c1) is number of letters c1 in initial string. Now, if
we fix such d after opening c1, that maximizes number of unique pairs(we will name it p) (c1, c2) at d distance, this
will be optimal d, and conditional probability of victory in situation of fixed c1 equals p / cnt(c1).

Now we only need to sum up conditional probabilities for different c1. Probability of c1 equals cnt(c1)/ n, thus answer
is .