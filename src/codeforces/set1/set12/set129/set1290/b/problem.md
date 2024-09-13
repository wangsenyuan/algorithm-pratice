Let's call two strings 𝑠
 and 𝑡
 anagrams of each other if it is possible to rearrange symbols in the string 𝑠
 to get a string, equal to 𝑡
.

Let's consider two strings 𝑠
 and 𝑡
 which are anagrams of each other. We say that 𝑡
 is a reducible anagram of 𝑠
 if there exists an integer 𝑘≥2
 and 2𝑘
 non-empty strings 𝑠1,𝑡1,𝑠2,𝑡2,…,𝑠𝑘,𝑡𝑘
 that satisfy the following conditions:

If we write the strings 𝑠1,𝑠2,…,𝑠𝑘
 in order, the resulting string will be equal to 𝑠
;
If we write the strings 𝑡1,𝑡2,…,𝑡𝑘
 in order, the resulting string will be equal to 𝑡
;
For all integers 𝑖
 between 1
 and 𝑘
 inclusive, 𝑠𝑖
 and 𝑡𝑖
 are anagrams of each other.
If such strings don't exist, then 𝑡
 is said to be an irreducible anagram of 𝑠
. Note that these notions are only defined when 𝑠
 and 𝑡
 are anagrams of each other.

For example, consider the string 𝑠=
 "gamegame". Then the string 𝑡=
 "megamage" is a reducible anagram of 𝑠
, we may choose for example 𝑠1=
 "game", 𝑠2=
 "gam", 𝑠3=
 "e" and 𝑡1=
 "mega", 𝑡2=
 "mag", 𝑡3=
 "e":


On the other hand, we can prove that 𝑡=
 "memegaga" is an irreducible anagram of 𝑠
.

You will be given a string 𝑠
 and 𝑞
 queries, represented by two integers 1≤𝑙≤𝑟≤|𝑠|
 (where |𝑠|
 is equal to the length of the string 𝑠
). For each query, you should find if the substring of 𝑠
 formed by characters from the 𝑙
-th to the 𝑟
-th has at least one irreducible anagram.

### ideas
1. 对于s，存在一个anagram t， (如果不存在，那么也就不存在不可规约的t)
2. 但是这个t, 没法reducable from s
3. 比如 s = abb, t = bba, （t是s的 anagram, 但是没法进行匹配)
4. 如果s由3个字符 abc, 那么 t = cba 肯定存在不可规约的t
5. 只需要在前两个字符出现的位置，截断，并把第三个字符放在第一个位置；
6. 现在只需要考虑s有两个字符组成，aabb   bbaa 如果有两个字符，也肯定存在不可规约的t
7. 所以答案就是看s是否由同一个字符组成吗？
8. 不大对
9. s = abba, (t = aabb 不是irreducible, t = bbaa, 也不是， )
10. 所以就是必须存在3个以上，如果是两个字符的，其中的一个字符必须占多数
11. 比如 s = abbaa 那么 bbaaa就是irreduciable的