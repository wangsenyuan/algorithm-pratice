Once Bob saw a string. It contained so many different letters, that the letters were marked by numbers, but at the same time each letter could be met in the string at most 10 times. Bob didn't like that string, because it contained repeats: a repeat of length x is such a substring of length 2x, that its first half coincides character by character with its second half. Bob started deleting all the repeats from the string. He does it as follows: while it's possible, Bob takes the shortest repeat, if it is not unique, he takes the leftmost one, and deletes its left half and everything that is to the left of this repeat.

You're given the string seen by Bob. Find out, what it will look like after Bob deletes all the repeats in the way described above.

### ideas
1. 遇到一个重复的段，就把它前半部分，及左边的都删除掉
2. 如果 a, b, c, d, c, d => 删除后就剩余 c, d
3. It's guaranteed that each letter can be met in the string at most 10 times.
4. 这个信息看样子比较重要
5. got， 记录x前面的位置（最多时10个），然后记录这个位置后面到当前位置，是否是一个repeat
6. 用双hash即可