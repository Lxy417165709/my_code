## 视频->图片
ffmpeg -i inputfile.avi -r 1 image-%3d.jpeg

## 图片->视频
ffmpeg -r 30 -i video_img/kk_%05d.jpg video.mp4



对于音频提取，可以使用-b:a 128k 指定音频的码率是128kb/s，-ar 44k 指定音频的采样频率为44kHz，完整命令如下：
ffmpeg -i video.mp4 -ab 320k audio.mp3
https://yingfeng.me/archives/1183


ffmpeg -r 30 -i video_img/kk_%05d.jpg -i kk.mp3 video.mp4
加音乐