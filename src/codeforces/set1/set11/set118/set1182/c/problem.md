You are given 𝑛
 words, each of which consists of lowercase alphabet letters. Each word contains at least one vowel. You are going to choose some of the given words and make as many beautiful lyrics as possible.

Each lyric consists of two lines. Each line consists of two words separated by whitespace.

A lyric is beautiful if and only if it satisfies all conditions below.

The number of vowels in the first word of the first line is the same as the number of vowels in the first word of the second line.
The number of vowels in the second word of the first line is the same as the number of vowels in the second word of the second line.
The last vowel of the first line is the same as the last vowel of the second line. Note that there may be consonants after the vowel.
Also, letters "a", "e", "o", "i", and "u" are vowels. Note that "y" is never vowel.

For example of a beautiful lyric,

"hello hellooowww"
"whatsup yowowowow"

is a beautiful lyric because there are two vowels each in "hello" and "whatsup", four vowels each in "hellooowww" and "yowowowow" (keep in mind that "y" is not a vowel), and the last vowel of each line is "o".
For example of a not beautiful lyric,

"hey man"
"iam mcdic"

is not a beautiful lyric because "hey" and "iam" don't have same number of vowels and the last vowels of two lines are different ("a" in the first and "i" in the second).
How many beautiful lyrics can you write from given words? Note that you cannot use a word more times than it is given to you. For example, if a word is given three times, you can use it at most three times.

### ideas
1. 所有的单词可以按照 元音的数量/最后一个元音字符 分组
2. 相同元音数量+最优一个元音字符 / 2 是它的上限
3. 