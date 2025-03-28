Vasya is a born Berland film director, he is currently working on a new blockbuster, "The Unexpected". Vasya knows from his own experience how important it is to choose the main characters' names and surnames wisely. He made up a list of n names and n surnames that he wants to use. Vasya haven't decided yet how to call characters, so he is free to match any name to any surname. Now he has to make the list of all the main characters in the following format: "Name1 Surname1, Name2 Surname2, ..., Namen Surnamen", i.e. all the name-surname pairs should be separated by exactly one comma and exactly one space, and the name should be separated from the surname by exactly one space. First of all Vasya wants to maximize the number of the pairs, in which the name and the surname start from one letter. If there are several such variants, Vasya wants to get the lexicographically minimal one. Help him.

An answer will be verified a line in the format as is shown above, including the needed commas and spaces. It's the lexicographical minimality of such a line that needs to be ensured. The output line shouldn't end with a space or with a comma.

### ideas
1. 应该首先处理A开头的
2. 但是这里有个问题，就是A不能把B的给匹配掉（优先保证相同字符开始的规则）
3. 所以，这里A是去找那些匹配剩下的部分