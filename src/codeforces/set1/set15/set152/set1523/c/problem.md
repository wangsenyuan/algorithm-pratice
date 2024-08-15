William is a huge fan of planning ahead. That is why he starts his morning routine by creating a nested list of upcoming errands.

A valid nested list is any list which can be created from a list with one item "1" by applying some operations. Each operation inserts a new item into the list, on a new line, just after one of existing items 𝑎1.𝑎2.𝑎3.⋯.𝑎𝑘
 and can be one of two types:

Add an item 𝑎1.𝑎2.𝑎3.⋯.𝑎𝑘.1
 (starting a list of a deeper level), or
Add an item 𝑎1.𝑎2.𝑎3.⋯.(𝑎𝑘+1)
 (continuing the current level).
Operation can only be applied if the list does not contain two identical items afterwards. And also, if we consider every item as a sequence of numbers, then the sequence of items should always remain increasing in lexicographical order. Examples of valid and invalid lists that are shown in the picture can found in the "Notes" section.
When William decided to save a Word document with the list of his errands he accidentally hit a completely different keyboard shortcut from the "Ctrl-S" he wanted to hit. It's not known exactly what shortcut he pressed but after triggering it all items in the list were replaced by a single number: the last number originally written in the item number.

William wants you to help him restore a fitting original nested list.