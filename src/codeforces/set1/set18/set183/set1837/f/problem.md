Berland Intercollegiate Contest has just finished. Monocarp and Polycarp, as the jury, are going to conduct an
editorial. Unfortunately, the time is limited, since they have to finish before the closing ceremony.

There were 𝑛
problems in the contest. The problems are numbered from 1
to 𝑛
. The editorial for the 𝑖
-th problem takes 𝑎𝑖
minutes. Monocarp and Polycarp are going to conduct an editorial for exactly 𝑘
of the problems.

The editorial goes as follows. They have a full problemset of 𝑛
problems before them, in order. They remove 𝑛−𝑘
problems without changing the order of the remaining 𝑘
problems. Then, Monocarp takes some prefix of these 𝑘
problems (possibly, an empty one or all problems). Polycarp takes the remaining suffix of them. After that, they go to
different rooms and conduct editorials for their problems in parallel. So, the editorial takes as much time as the
longer of these two does.

Please, help Monocarp and Polycarp to choose the problems and the split in such a way that the editorial finishes as
early as possible. Print the duration of the editorial.

### ideas

1. binary search on the min sum
2. use a heap for the prefix, take the least values until sum > expect
3. check whether exists a (k - L) suffix (sum <= expect) using AVL tree