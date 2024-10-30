# Convert mp4 to gif with High Quality

```bash
mkdir frames
ffmpeg -i video.mp4 frames/%04d.png
gifski -o output.gif frames/*.png
```

-r: fps
-W: width

ref:
https://stackoverflow.com/questions/42980663/ffmpeg-high-quality-animated-gif
