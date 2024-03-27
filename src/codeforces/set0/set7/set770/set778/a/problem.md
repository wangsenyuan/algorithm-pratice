Little Nastya has a hobby, she likes to remove some letters from word, to obtain another word. But it turns out to be
pretty hard for her, because she is too young. Therefore, her brother Sergey always helps her.

Sergey gives Nastya the word t and wants to get the word p out of it. Nastya removes letters in a certain order (one
after another, in this order strictly), which is specified by permutation of letters' indices of the word t: a1... a|t|.
We denote the length of word x as |x|. Note that after removing one letter, the indices of other letters don't change.
For example, if t ="nastya" and a =[4, 1, 5, 3, 2, 6] then removals make the following sequence of words "nastya"  "
nastya"  "nastya"  "nastya"  "nastya"  "nastya"  "nastya".

Sergey knows this permutation. His goal is to stop his sister at some point and continue removing by himself to get the
word p. Since Nastya likes this activity, Sergey wants to stop her as late as possible. Your task is to determine, how
many letters Nastya can remove before she will be stopped by Sergey.

It is guaranteed that the word p can be obtained by removing the letters from word t.

### ideas

1. binary search
2. 