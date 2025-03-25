You have 𝑛
 videos on your watchlist on the popular platform YooCube. The 𝑖
-th video lasts 𝑑𝑖
 minutes.

YooCube has recently increased the frequency of their ads. Ads are shown only between videos. After finishing a video, an ad is shown if either of these two conditions is true:

three videos have been watched since the last ad;
at least 𝑘
 minutes have passed since the end of the last ad.
You want to watch the 𝑛
 videos in your watchlist. Given that you have just watched an ad, and that you can choose the order of the 𝑛
 videos, what is the minimum number of ads that you are forced to watch? You can start a new video immediately after the previous video or ad ends, and you don't have to watch any ad after you finish.


### ideas
1. 假设目前已经看了2个（且free 时间没有超过k)
2. 那么下一个最好是选择是时长最长的video
3. 但前提是2个没有超时
4. 还有那种单个超时的（这种最好和不超时的匹配）
5. 