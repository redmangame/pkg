/*
 * @Author: redmangame redmangame@163.com
 * @Date: 2024-04-26 14:13:43
 * @LastEditors: redmangame redmangame@163.com
 * @LastEditTime: 2024-04-26 14:44:10
 * @FilePath: \fyneguiffmpeg\pkg\ffmpeg\getDuration.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package ffmpegv1

import (
	"os/exec"
	"strconv"
	"strings"
)

/*
filePath := "video.mp4" // 替换为你的视频或音频文件路径
duration, err := getMediaDuration(filePath)
if err != nil {
	fmt.Println("获取媒体文件时长出错:", err)
	return
}

fmt.Printf("媒体文件 %s 的时长为 %.2f 秒\n", filePath, duration)
*/

func GetMediaDuration(filePath string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	durationStr := strings.TrimSpace(string(output))
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
