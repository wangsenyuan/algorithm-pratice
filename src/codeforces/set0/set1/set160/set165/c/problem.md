A string is binary, if it consists only of characters "0" and "1".

String v is a substring of string w if it has a non-zero length and can be read starting from some position in string w. For example, string "010" has six substrings: "0", "1", "0", "01", "10", "010". Two substrings are considered different if their positions of occurrence are different. So, if some string occurs multiple times, we should consider it the number of times it occurs.

You are given a binary string s. Your task is to find the number of its substrings, containing exactly k characters "1".

### ideas
1. two pointers?