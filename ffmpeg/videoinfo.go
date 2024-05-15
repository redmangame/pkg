/*
使用了ffprobe命令来获取视频的宽度和高度信息。然后，通过比较宽度和高度的值，判断视频的方向。如果宽度大于高度，则为横屏；如果宽度小于高度，则为竖屏；如果宽度和高度相等，则为正方形。

请确保您的系统中已经安装了ffprobe命令，并且其可执行文件位于环境变量中。如果没有安装，请根据您的操作系统下载并安装FFmpeg工具包，并将ffprobe添加到您的环境变量中

*/
package ffmpegv1

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//获取视频的形状
func GetVideoOrientation(filename string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=p=0", filename)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// 使用正则表达式提取宽度和高度
	pattern := regexp.MustCompile(`(\d+),(\d+)`)
	match := pattern.FindStringSubmatch(string(output))
	if len(match) != 3 {
		return "", fmt.Errorf("Failed to parse video width and height")
	}

	width := match[1]
	height := match[2]

	// 根据宽度和高度比较来确定视频方向
	if width > height {
		return "横屏", nil
	} else if width < height {
		return "竖屏", nil
	} else {
		return "正方形", nil
	}
}

//获取视频的形状
func GetVideoWH(filename string) (width float64, height float64, err error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=p=0", filename)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, 0, fmt.Errorf("执行ffprobe命令失败：%v", err)
	}

	// 使用正则表达式提取宽度和高度
	pattern := regexp.MustCompile(`(\d+),(\d+)`)
	match := pattern.FindStringSubmatch(string(output))
	if len(match) != 3 {
		return 0, 0, fmt.Errorf("无法解析视频宽度和高度")
	}

	widthS := match[1]
	heightS := match[2]

	width, err = strconv.ParseFloat(widthS, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("宽度转换失败：%v", err)
	}
	height, err = strconv.ParseFloat(heightS, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("高度转换失败：%v", err)
	}

	return width, height, nil
}

//原始视频文件得出生成的对应gif文件 包含完整路径
func GetConvertToGifPath(inputFile string) string {
	// 获取输入文件的路径和文件名信息
	outputFilePath := filepath.Dir(inputFile)
	outputFileName := filepath.Base(inputFile)
	outputFileNameWithoutExt := strings.TrimSuffix(outputFileName, filepath.Ext(outputFileName))
	outputFileExt := ".gif"

	// 拼接输出文件的完整路径
	outputFile := filepath.Join(outputFilePath, outputFileNameWithoutExt+outputFileExt)

	// 返回输出文件的路径
	return outputFile
}

//原始视频文件得出生成的对应缩放后的文件 包含完整路径
func GetConvertToScalePath(inputFile string) string {
	// 获取输入文件的路径和文件名信息
	outputFilePath := filepath.Dir(inputFile)
	outputFileName := filepath.Base(inputFile)
	outputFileNameWithoutExt := strings.TrimSuffix(outputFileName, filepath.Ext(outputFileName))
	outputFileExt := filepath.Ext(outputFileName)

	// 拼接输出文件的完整路径
	outputFile := filepath.Join(outputFilePath, outputFileNameWithoutExt+"_scale_"+outputFileExt)

	// 返回输出文件的路径
	return outputFile
}

//视频文件得出生成的对应画板的文件 包含完整路径
func GetConvertToPalettegenPath(inputFile string) string {
	// 获取输入文件的路径和文件名信息
	outputFilePath := filepath.Dir(inputFile)
	outputFileName := filepath.Base(inputFile)
	outputFileNameWithoutExt := strings.TrimSuffix(outputFileName, filepath.Ext(outputFileName))
	//outputFileExt := filepath.Ext(outputFileName)
	outputFileExt := ".png"
	// 拼接输出文件的完整路径
	outputFile := filepath.Join(outputFilePath, outputFileNameWithoutExt+"_palettegen_"+outputFileExt)

	// 返回输出文件的路径
	return outputFile
}
