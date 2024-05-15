/*
<<<<<<< HEAD
package main

import (
	"path/filepath"

	ffmpegv1 "github.com/redmangame/fyneguiffmpeg/pkg/ffmpeg"
)

func main() {
需要pid
	pidChan := make(chan int)
	var pid int
		go func() {
			videoPath := `C:\Users\hyx\Desktop\fa8a9b845de6c7013f997d0362ba1a25_raw.mp4`
			outPath := videoPath[:len(videoPath)-len(filepath.Ext(videoPath))] + "_transcode_" + filepath.Ext(videoPath)
			transcoding(videoPath, outPath, pidChan)
		}()
		go func() {
			// for {

			// }
			pid = <-pidChan
			log.Println("接收到的PID:", pid)
		}()





	不需要pid
		videoPath := `C:\Users\hyx\Desktop\fa8a9b845de6c7013f997d0362ba1a25_raw.mp4`
	outPath := videoPath[:len(videoPath)-len(filepath.Ext(videoPath))] + "_transcode_" + filepath.Ext(videoPath)
	ffmpegv1.Transcoding(videoPath, outPath)

}
*/
=======
 * @Author: redmangame redmangame@163.com
 * @Date: 2024-04-25 15:01:36
 * @LastEditors: redmangame redmangame@163.com
 * @LastEditTime: 2024-05-06 10:16:08
 * @FilePath: \fyneguiffmpeg\pkg\ffmpeg\transcodevideo.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
>>>>>>> remotes/origin/main
package ffmpegv1

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Transcoding(inPath, outPath string) {

	//videoPath := `C:\Users\hyx\Desktop\fa8a9b845de6c7013f997d0362ba1a25_raw.mp4`
	//outputPath := videoPath[:len(videoPath)-len(filepath.Ext(videoPath))] + "_transcode_" + filepath.Ext(videoPath)

	videoPath := inPath
	outputPath := outPath

	var done bool
	var pid int
	var err error
	var progress float64
	progressChan := make(chan float64)
<<<<<<< HEAD
	pidChan := make(chan int)
	defer close(progressChan)

	go func() {
		progress, pid, err = TranscodeVideo(videoPath, outputPath, progressChan, pidChan)
=======
	defer close(progressChan)

	go func() {
		progress, pid, err = TranscodeVideo(videoPath, outputPath, progressChan)
>>>>>>> remotes/origin/main
		fmt.Println("in main progress go run1", progress)
		if err != nil {
			fmt.Println("Error transcoding video:", err)
			progressChan <- -1 // Signal error
			return
		}
		//progressChan <- progress
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastProgress float64

outerLoop:
	for {
		select {
		case progress := <-progressChan:
			if progress >= 0 {
				lastProgress = progress
				fmt.Printf("Transcoding progress: %.2f%%\n", progress*100)
			} else if progress >= 1.00 {
				fmt.Printf("Transcoding progress: %.2f%%\n", progress*100)
				done = true
				break outerLoop // 使用带有标签的 break 终止外部循环
			} else {
				fmt.Println("Transcoding failed.")
				break outerLoop // 使用带有标签的 break 终止外部循环
			}
		case <-ticker.C:
			fmt.Printf("Transcoding progress: %.2f%% (no recent updates) %v\n", lastProgress*100, lastProgress)
			if lastProgress >= 1.00 {
				fmt.Printf("Transcoding >= 1.00 progress: %.2f%%\n", lastProgress*100)
				done = true
				break outerLoop // 使用带有标签的 break 终止外部循环
			}
<<<<<<< HEAD
		case pid := <-pidChan:
			fmt.Printf("Transcoding progress pid: %d \n", pid)
		}

		if done {
			break // 终止内部循环
		}
	}

	fmt.Println("完成", pid, done)

}
func TranscodingByPid(inPath, outPath string, pidchan chan<- int) {
	videoPath := inPath
	outputPath := outPath

	var done bool
	var pid int
	var err error
	var progress float64
	progressChan := make(chan float64)
	defer close(progressChan)

	go func() {
		progress, pid, err = TranscodeVideo(videoPath, outputPath, progressChan, pidchan)
		fmt.Println("in main progress go run1", progress, pid, err)
		if err != nil {
			fmt.Println("Error transcoding video:", err)
			progressChan <- -1 // Signal error
			pidchan <- -1
			return
		}

	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastProgress float64

outerLoop:
	for {
		select {
		case progress := <-progressChan:
			if progress >= 0 {
				lastProgress = progress
				//progressFyne.SetValue(lastProgress)

				fmt.Printf("Transcoding progress: %.2f%% \n", progress*100)
			} else if progress >= 1.00 {
				//fmt.Printf("Transcoding progress: %.2f%%\n", progress*100)
				done = true
				break outerLoop // 使用带有标签的 break 终止外部循环
			} else {
				fmt.Println("Transcoding failed.")
				break outerLoop // 使用带有标签的 break 终止外部循环
			}
		case <-ticker.C:
			//fmt.Printf("Transcoding progress: %.2f%% (no recent updates) %v\n", lastProgress*100, lastProgress)
			if lastProgress >= 1.00 {
				fmt.Printf("Transcoding >= 1.00 progress: %.2f%%\n", lastProgress*100)
				done = true
				break outerLoop // 使用带有标签的 break 终止外部循环
			}
			// case pid:= <-pidChan:
			// 	fmt.Printf("Transcoding progress pid: %d \n", pid)
=======
>>>>>>> remotes/origin/main
		}

		if done {
			break // 终止内部循环
		}
	}

	fmt.Println("完成", pid, done)

}

<<<<<<< HEAD
func TranscodeVideo(videoPath, outputPath string, progressChanOut chan<- float64, pidChan chan<- int) (float64, int, error) {
=======
func TranscodeVideo(videoPath, outputPath string, progressChanOut chan<- float64) (float64, int, error) {
>>>>>>> remotes/origin/main
	var pid int

	// 文件时长
	fileDuration, err := GetMediaDuration(videoPath)
	if err != nil {
		return fileDuration, pid, err
	}

	var progress float64
	progressChan := make(chan float64)
	defer close(progressChan) // 关闭 progressChan 通道

	// cmd := exec.Command("ffmpeg",
	// 	"-y",
	// 	"-i", videoPath,
	// 	"-c:v", "libx264",
	// 	"-preset", "medium",
	// 	"-crf", "23",
	// 	"-c:a", "aac",
	// 	"-b:a", "128k",
	// 	"-movflags", "faststart",
	// 	outputPath,
	// )

	cmd := exec.Command("ffmpeg",
		"-y",
		"-i", videoPath,
		outputPath,
	)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return progress, pid, err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	re := regexp.MustCompile(`time=(\d+:\d+:\d+\.\d+)`)

	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)
			if err != nil {
				break
			}
			output := string(buf[:n])
			match := re.FindStringSubmatch(output)
			if len(match) > 1 {
				timeStr := match[1]
				progress = CalculateProgress(timeStr, fileDuration)
				fmt.Println("in func progress ", progress, pid)
				progressChan <- progress
				if progress >= 1.00 {
					break
				}
			}
		}
	}()

	go func() {
		err := cmd.Start()
		if cmd.Process != nil {
			pid = cmd.Process.Pid
<<<<<<< HEAD
			pidChan <- pid
=======
>>>>>>> remotes/origin/main
		}
		//fmt.Println(cmd.Args, cmd.Process.Pid)
		if err != nil {
			fmt.Println("Error starting ffmpeg:", err)
			return
		}
		defer wg.Wait()

		for {
			select {
			case progress = <-progressChan:
				// Do nothing, just update progress value
<<<<<<< HEAD
				//fmt.Println("in func progress2 ", progress, pid)
=======
				fmt.Println("in func progress2 ", progress, pid)
>>>>>>> remotes/origin/main

				progressChanOut <- progress
				if progress >= 1.00 {
					return
				}
			default:
				// If no new progress value received, continue waiting
				//fmt.Println("in func progress2-1: ", progress, pid)
				if progress >= 1.10 {
					return
				}
			}
		}
	}()

	wg.Wait()
	err = cmd.Wait()

	return progress, pid, err
}

func CalculateProgress(timeStr string, fileDuration float64) float64 {
	totalDurationInSeconds := fileDuration
	parts := strings.Split(timeStr, ":")
	hours, _ := strconv.ParseFloat(parts[0], 64)
	minutes, _ := strconv.ParseFloat(parts[1], 64)
	seconds, _ := strconv.ParseFloat(parts[2], 64)
	totalSeconds := hours*3600 + minutes*60 + seconds
	return totalSeconds / totalDurationInSeconds
}
