Vasya is an administrator of a public page of organization "Mouse and keyboard" and his everyday duty is to publish news
from the world of competitive programming. For each news he also creates a list of hashtags to make searching for a
particular topic more comfortable. For the purpose of this problem we define hashtag as a string consisting of lowercase
English letters and exactly one symbol '#' located at the beginning of the string. The length of the hashtag is defined
as the number of symbols in it without the symbol '#'.

The head administrator of the page told Vasya that hashtags should go in lexicographical order (take a look at the notes
section for the definition).

Vasya is lazy so he doesn't want to actually change the order of hashtags in already published news. Instead, he decided
to delete some suffixes (consecutive characters at the end of the string) of some of the hashtags. He is allowed to
delete any number of characters, even the whole string except for the symbol '#'. Vasya wants to pick such a way to
delete suffixes that the total number of deleted symbols is minimum possible. If there are several optimal solutions, he
is fine with any of them.

### ideas

1. 先假设只有两个字符串x, y
2. 如果 x > y, 要删除最少的情况下，保证 pref(x) < pref(y)
3. 似乎只能贪心的处理，就是删除x的后缀，直到满足条件
4. 但是这样子，不就非常简单了吗？