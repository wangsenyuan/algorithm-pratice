You all know that the Library of Bookland is the largest library in the world. There are dozens of thousands of books in the library.

Some long and uninteresting story was removed...

The alphabet of Bookland is so large that its letters are denoted by positive integers. Each letter can be small or large, the large version of a letter x is denoted by x'. BSCII encoding, which is used everywhere in Bookland, is made in that way so that large letters are presented in the order of the numbers they are denoted by, and small letters are presented in the order of the numbers they are denoted by, but all large letters are before all small letters. For example, the following conditions hold: 2 < 3, 2' < 3', 3' < 2.

A word x1, x2, ..., xa is not lexicographically greater than y1, y2, ..., yb if one of the two following conditions holds:

a ≤ b and x1 = y1, ..., xa = ya, i.e. the first word is the prefix of the second word;
there is a position 1 ≤ j ≤ min(a, b), such that x1 = y1, ..., xj - 1 = yj - 1 and xj < yj, i.e. at the first position where the words differ the first word has a smaller letter than the second word has.
For example, the word "3' 7 5" is before the word "2 4' 6" in lexicographical order. It is said that sequence of words is in lexicographical order if each word is not lexicographically greater than the next word in the sequence.

Denis has a sequence of words consisting of small letters only. He wants to change some letters to large (let's call this process a capitalization) in such a way that the sequence of words is in lexicographical order. However, he soon realized that for some reason he can't change a single letter in a single word. He only can choose a letter and change all of its occurrences in all words to large letters. He can perform this operation any number of times with arbitrary letters of Bookland's alphabet.

Help Denis to choose which letters he needs to capitalize (make large) in order to make the sequence of words lexicographically ordered, or determine that it is impossible.

Note that some words can be equal.

### ideas
1. v 表示不大些, v + m 表示大写，就看能不能做到2-sat