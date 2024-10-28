You got a box with a combination lock. The lock has a display showing n digits. There are two buttons on the box, each button changes digits on the display. You have quickly discovered that the first button adds 1 to all the digits (all digits 9 become digits 0), and the second button shifts all the digits on the display one position to the right (the last digit becomes the first one). For example, if the display is currently showing number 579, then if we push the first button, the display will show 680, and if after that we push the second button, the display will show 068.

You know that the lock will open if the display is showing the smallest possible number that can be obtained by pushing the buttons in some order. The leading zeros are ignored while comparing numbers. Now your task is to find the desired number.

### ideas
1. 最小的情况，就是看，如何得到0,1,2...类似这样的结果
2. 所有位都可以变成0，变成0后，把它shift到第一位，然后判断就可以了s