You are given two strings 𝑎
 and 𝑏
 consisting of lowercase English letters, both of length 𝑛
. The characters of both strings have indices from 1
 to 𝑛
, inclusive.

You are allowed to do the following changes:

Choose any index 𝑖
 (1≤𝑖≤𝑛
) and swap characters 𝑎𝑖
 and 𝑏𝑖
;
Choose any index 𝑖
 (1≤𝑖≤𝑛
) and swap characters 𝑎𝑖
 and 𝑎𝑛−𝑖+1
;
Choose any index 𝑖
 (1≤𝑖≤𝑛
) and swap characters 𝑏𝑖
 and 𝑏𝑛−𝑖+1
.
Note that if 𝑛
 is odd, you are formally allowed to swap 𝑎⌈𝑛2⌉
 with 𝑎⌈𝑛2⌉
 (and the same with the string 𝑏
) but this move is useless. Also you can swap two equal characters but this operation is useless as well.

You have to make these strings equal by applying any number of changes described above, in any order. But it is obvious that it may be impossible to make two strings equal by these swaps.

In one preprocess move you can replace a character in 𝑎
 with another character. In other words, in a single preprocess move you can choose any index 𝑖
 (1≤𝑖≤𝑛
), any character 𝑐
 and set 𝑎𝑖:=𝑐
.

Your task is to find the minimum number of preprocess moves to apply in such a way that after them you can make strings 𝑎
 and 𝑏
 equal by applying some number of changes described in the list above.

Note that the number of changes you make after the preprocess moves does not matter. Also note that you cannot apply preprocess moves to the string 𝑏
 or make any preprocess moves after the first change is made.