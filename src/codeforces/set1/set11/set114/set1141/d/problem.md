There are 𝑛
 left boots and 𝑛
 right boots. Each boot has a color which is denoted as a lowercase Latin letter or a question mark ('?'). Thus, you are given two strings 𝑙
 and 𝑟
, both of length 𝑛
. The character 𝑙𝑖
 stands for the color of the 𝑖
-th left boot and the character 𝑟𝑖
 stands for the color of the 𝑖
-th right boot.

A lowercase Latin letter denotes a specific color, but the question mark ('?') denotes an indefinite color. Two specific colors are compatible if they are exactly the same. An indefinite color is compatible with any (specific or indefinite) color.

For example, the following pairs of colors are compatible: ('f', 'f'), ('?', 'z'), ('a', '?') and ('?', '?'). The following pairs of colors are not compatible: ('f', 'g') and ('a', 'z').

Compute the maximum number of pairs of boots such that there is one left and one right boot in a pair and their colors are compatible.

Print the maximum number of such pairs and the pairs themselves. A boot can be part of at most one pair.
