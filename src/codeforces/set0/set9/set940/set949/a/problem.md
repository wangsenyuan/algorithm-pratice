Oleg writes down the history of the days he lived. For each day he decides if it was good or bad. Oleg calls a non-empty
sequence of days a zebra, if it starts with a bad day, ends with a bad day, and good and bad days are alternating in it.
Let us denote bad days as 0 and good days as 1. Then, for example, sequences of days 0, 010, 01010 are zebras, while
sequences 1, 0110, 0101 are not.

Oleg tells you the story of days he lived in chronological order in form of string consisting of 0 and 1. Now you are
interested if it is possible to divide Oleg's life history into several subsequences, each of which is a zebra, and the
way it can be done. Each day must belong to exactly one of the subsequences. For each of the subsequences, days forming
it must be ordered chronologically. Note that subsequence does not have to be a group of consecutive days.

### ideas

1. 当前遇到1，如果前面没有active 的0 => -1
2. else 就append到0的后面，并减少一个 active
3. 当前如果遇到一个0，那么如果前面有1结尾的，就append上去，并更新即可
4. else 新开一个