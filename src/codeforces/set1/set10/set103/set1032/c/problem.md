Little Paul wants to learn how to play piano. He already has a melody he wants to start with. For simplicity he
represented this melody as a sequence 𝑎1,𝑎2,…,𝑎𝑛
of key numbers: the more a number is, the closer it is to the right end of the piano keyboard.

Paul is very clever and knows that the essential thing is to properly assign fingers to notes he's going to play. If he
chooses an inconvenient fingering, he will then waste a lot of time trying to learn how to play the melody by these
fingers and he will probably not succeed.

Let's denote the fingers of hand by numbers from 1
to 5
. We call a fingering any sequence 𝑏1,…,𝑏𝑛
of fingers numbers. A fingering is convenient if for all 1≤𝑖≤𝑛−1
the following holds:

if 𝑎𝑖<𝑎𝑖+1
then 𝑏𝑖<𝑏𝑖+1
, because otherwise Paul needs to take his hand off the keyboard to play the (𝑖+1)
-st note;
if 𝑎𝑖>𝑎𝑖+1
then 𝑏𝑖>𝑏𝑖+1
, because of the same;
if 𝑎𝑖=𝑎𝑖+1
then 𝑏𝑖≠𝑏𝑖+1
, because using the same finger twice in a row is dumb. Please note that there is ≠
, not =
between 𝑏𝑖
and 𝑏𝑖+1
.
Please provide any convenient fingering or find out that there is none.

Input
The first line contains a single integer 𝑛
(1≤𝑛≤105
) denoting the number of notes.

The second line contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤2⋅105
) denoting the positions of notes on the keyboard.

Output
If there is no convenient fingering, print −1
. Otherwise, print 𝑛
numbers 𝑏1,𝑏2,…,𝑏𝑛
, each from 1
to 5
, denoting a convenient fingering, separated by spaces.

### ideas

1. interesting problem
2. dp[i][j] 表示到i为止，使用finger j play at i (yes or no)