You are given a book with 𝑛
 chapters.

Each chapter has a specified list of other chapters that need to be understood in order to understand this chapter. To understand a chapter, you must read it after you understand every chapter on its required list.

Currently you don't understand any of the chapters. You are going to read the book from the beginning till the end repeatedly until you understand the whole book. Note that if you read a chapter at a moment when you don't understand some of the required chapters, you don't understand this chapter.

Determine how many times you will read the book to understand every chapter, or determine that you will never understand every chapter no matter how many times you read the book.


### ideas
1. 和简单的dag还不一样
2. 就是假设a依赖b，如果b在a的前面，那么就是在一次内读完，否则就要在两次后读完