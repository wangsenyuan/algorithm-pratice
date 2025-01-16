You are given a text consisting of 𝑛
 space-separated words. There is exactly one space character between any pair of adjacent words. There are no spaces before the first word and no spaces after the last word. The length of text is the number of letters and spaces in it. 𝑤𝑖
 is the 𝑖
-th word of text. All words consist only of lowercase Latin letters.

Let's denote a segment of words 𝑤[𝑖..𝑗]
 as a sequence of words 𝑤𝑖,𝑤𝑖+1,…,𝑤𝑗
. Two segments of words 𝑤[𝑖1..𝑗1]
 and 𝑤[𝑖2..𝑗2]
 are considered equal if 𝑗1−𝑖1=𝑗2−𝑖2
, 𝑗1≥𝑖1
, 𝑗2≥𝑖2
, and for every 𝑡∈[0,𝑗1−𝑖1]
 𝑤𝑖1+𝑡=𝑤𝑖2+𝑡
. For example, for the text "to be or not to be" the segments 𝑤[1..2]
 and 𝑤[5..6]
 are equal, they correspond to the words "to be".

An abbreviation is a replacement of some segments of words with their first uppercase letters. In order to perform an abbreviation, you have to choose at least two non-intersecting equal segments of words, and replace each chosen segment with the string consisting of first letters of the words in the segment (written in uppercase). For example, for the text "a ab a a b ab a a b c" you can replace segments of words 𝑤[2..4]
 and 𝑤[6..8]
 with an abbreviation "AAA" and obtain the text "a AAA b AAA b c", or you can replace segments of words 𝑤[2..5]
 and 𝑤[6..9]
 with an abbreviation "AAAB" and obtain the text "a AAAB AAAB c".

What is the minimum length of the text after at most one abbreviation?

### ideas
1. 是不是越长越好？好像也不是
2. 比如 aaab  aaab  a b a b
3. 显然缩写aaab 要更好一些
4. hash + binary search吗？
5. 至少两个，不是正好2个