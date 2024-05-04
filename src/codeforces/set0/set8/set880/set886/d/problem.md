A substring of some string is called the most frequent, if the number of its occurrences is not less than number of
occurrences of any other substring.

You are given a set of strings. A string (not necessarily from this set) is called good if all elements of the set are
the most frequent substrings of this string. Restore the non-empty good string with minimum length. If several such
strings exist, restore lexicographically minimum string. If there are no good strings, print "NO" (without quotes).

A substring of a string is a contiguous subsequence of letters in the string. For example, "ab", "c", "abc" are
substrings of string "abc", while "ac" is not a substring of that string.

The number of occurrences of a substring in a string is the number of starting positions in the string where the
substring occurs. These occurrences could overlap.

String a is lexicographically smaller than string b, if a is a prefix of b, or a has a smaller letter at the first
position where a and b differ.

### ideas

1. 如果word[i]完全包含在word[j]中，那么可以忽略i,
2. 如果word[i]和word[j]有部分重叠，则没有答案 =》 因为 它们重叠的部分，更加frequent
3. 然后就是它们之间没有关系，那么就按照顺序排即可
4. 第2点，可以直接找重叠的字符即可，有重叠的字符，然后检查是否包含，
5. 第二点确实在大部分情况下成立，但是如果重叠的部分，是一个的前缀，另外一个的后缀，那么这种情况也是存在的
6. 以字符开始的只能有一个，否则也是不行的

### solution

If some string is the most frequent then all its substrings are the most frequent too.
If string ab or similar is the most frequent then letter a is always followed by letter b and b always follow a.
Let's consider directed graph on letters where edge a → b exists only if ab is the most frequent. If there is cycle in
such graph then good string doesn't exist.
So such graph can be represented as several non-intersecting paths. All strings which correspond to paths must occur in
non-empty good string. So if we print them in lexicographical order then we will get the answer.