Rudolph is a scientist who studies alien life forms. There is a room in front of Rudolph with 𝑛
different objects scattered around. Among the objects there is exactly one amazing creature — a mimic that can turn into
any object. He has already disguised himself in this room and Rudolph needs to find him by experiment.

The experiment takes place in several stages. At each stage, the following happens:

Rudolf looks at all the objects in the room and writes down their types. The type of each object is indicated by a
number; there can be several objects of the same type.
After inspecting, Rudolph can point to an object that he thinks is a mimic. After that, the experiment ends. Rudolph
only has one try, so if he is unsure of the mimic's position, he does the next step instead.
Rudolf can remove any number of objects from the room (possibly zero). Then Rudolf leaves the room and at this time all
objects, including the mimic, are mixed with each other, their order is changed, and the mimic can transform into any
other object (even one that is not in the room).
After this, Rudolf returns to the room and repeats the stage. The mimic may not change appearance, but it can not remain
a same object for more than two stages in a row.
Rudolf's task is to detect mimic in no more than five stages.

### thoughts

1. 首先在第一次stage，不能删除任何一个
2. 然后如果没有变化，要再等一轮，
3. 一旦有变化，必然是一个减少了，一个增加了，保留增加的那些
4. 删除所有其他的
5. 然后再等最多两轮即可