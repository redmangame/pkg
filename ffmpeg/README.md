package main

import (
	"path/filepath"

	ffmpegv1 "github.com/redmangame/fyneguiffmpeg/pkg/ffmpeg"
)

func main() {
	videoPath := `C:\Users\hyx\Desktop\fa8a9b845de6c7013f997d0362ba1a25_raw.mp4`
	outPath := videoPath[:len(videoPath)-len(filepath.Ext(videoPath))] + "_transcode_" + filepath.Ext(videoPath)
	ffmpegv1.Transcoding(videoPath, outPath)

}
