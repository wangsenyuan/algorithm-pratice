Berland scientists know that the Old Berland language had exactly n words. Those words had lengths of l1, l2, ..., ln letters. Every word consisted of two letters, 0 and 1. Ancient Berland people spoke quickly and didn’t make pauses between the words, but at the same time they could always understand each other perfectly. It was possible because no word was a prefix of another one. The prefix of a string is considered to be one of its substrings that starts from the initial symbol.


### ideas
1. sort l first
2. 把它们当成一棵树，每层的个数不能超过 1 << l