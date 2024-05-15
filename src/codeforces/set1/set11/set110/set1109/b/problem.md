Reading books is one of Sasha's passions. Once while he was reading one book, he became acquainted with an unusual character. The character told about himself like that: "Many are my names in many countries. Mithrandir among the Elves, Tharkûn to the Dwarves, Olórin I was in my youth in the West that is forgotten, in the South Incánus, in the North Gandalf; to the East I go not."

And at that moment Sasha thought, how would that character be called in the East? In the East all names are palindromes. A string is a palindrome if it reads the same backward as forward. For example, such strings as "kazak", "oo" and "r" are palindromes, but strings "abb" and "ij" are not.

Sasha believed that the hero would be named after one of the gods of the East. As long as there couldn't be two equal names, so in the East people did the following: they wrote the original name as a string on a piece of paper, then cut the paper minimum number of times 𝑘
, so they got 𝑘+1
 pieces of paper with substrings of the initial string, and then unite those pieces together to get a new string. Pieces couldn't be turned over, they could be shuffled.

In this way, it's possible to achive a string abcdefg from the string f|de|abc|g using 3
 cuts (by swapping papers with substrings f and abc). The string cbadefg can't be received using the same cuts.

More formally, Sasha wants for the given palindrome 𝑠
 find such minimum 𝑘
, that you can cut this string into 𝑘+1
 parts, and then unite them in such a way that the final string will be a palindrome and it won't be equal to the initial string 𝑠
. It there is no answer, then print "Impossible" (without quotes).

### ideas
1. 如果 n时奇数的， 那么 cut后，最中间的字母，还是原有的字母
2. 看起来挺混乱的，但是，考虑一个点是，大部分情况，s中的字符其实是可以保留不变的
3. 因为它们本来就已经匹配了，如果对它们进行处理，会带来混乱
4. 然后考虑是否能够在足够小的cut后，得到新的回文串
5. 假设s[i]  = s[j], s[u] = s[v], 且s[u] != s[v]
6. 那么在最多8次cut后，就可以得到新的回文串 (i, j, u, v 的前后cut一次)
7. 如果不存在这样的情况，那么也不存在新的回文串
8. 然后考虑是否有更好的情况如果 (i, u)是连接在一起的，那么，可以在最多4次cut出来 （而且这个始终是存在的）
9. 如果i = 0 (j = n - 1) 那么在2次后
10. 还有一种就是检查1次的，中间某个地方cut后，是否能组成回文