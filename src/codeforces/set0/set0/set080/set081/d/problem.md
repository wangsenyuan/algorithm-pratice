Polycarp loves not only to take pictures, but also to show his photos to friends. On his personal website he has recently installed a widget that can display n photos with the scroll option. At each moment of time the widget displays exactly one photograph with the option showing the previous/next one. From the first photo, you can switch to the second one or to the n-th one, from the second photo you can switch to the third one or to the first one, etc. Thus, navigation is performed in a cycle.

Polycarp's collection consists of m photo albums, the i-th album contains ai photos. Polycarp wants to choose n photos and put them on a new widget. To make watching the photos interesting to the visitors, he is going to post pictures so that no two photos from one album were neighboring (each photo will have exactly two neighbors, the first photo's neighbors are the second and the n-th one).

Help Polycarp compile a photo gallery. Select n photos from his collection and put them in such order that no two photos from one album went one after the other.

### ideas
1. 考虑n = 3, m = 2
2. 如果n是偶数，比如4， 那么1, 3, 5, ... n - 1（放置最多的那个，然后其他的就比较好处理了）
3. 如果n是奇数，就比较难处理了 1, 3, 5, .... n - 1 (放置最多的那个)