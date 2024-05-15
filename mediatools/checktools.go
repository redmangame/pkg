/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2024-04-16 10:29:24
<<<<<<< HEAD
 * @LastEditors: redmangame redmangame@163.com
<<<<<<< HEAD
 * @LastEditTime: 2024-05-14 15:50:42
=======
 * @LastEditTime: 2024-05-14 14:54:09
>>>>>>> 57c1a529f170b5b38b9c894a33ef9114ee6767b0
=======
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-04-16 10:40:21
>>>>>>> remotes/origin/main
 * @FilePath: \splitfile\mediatools\checktools.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mediatools

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
=======
	"os"
	"os/exec"
	"path/filepath"
>>>>>>> 57c1a529f170b5b38b9c894a33ef9114ee6767b0

	"github.com/redmangame/pkg/exepath"
=======
	"os"
	"os/exec"
	"path/filepath"

	"github.com/redmangame/splitfile/pkg/exepath"
>>>>>>> remotes/origin/main
)

var exeDir string

func init() {
	exeDir = exepath.GetCurrentDir()
}

// HasFFMPEG checks if ffmpeg is available in the system or in the current directory
func HasFFMPEG() bool {
	// Check in system PATH
	_, err := exec.LookPath("ffmpeg")
	if err == nil {
		return true
	}

	// Check in current directory
	exePath := filepath.Join(exeDir, "ffmpeg")
	if _, err := os.Stat(exePath); err == nil {
		return true
	}

	return false
}

// HasFFProbe checks if ffprobe is available in the system or in the current directory
func HasFFProbe() bool {
	// Check in system PATH
	_, err := exec.LookPath("ffprobe")
	if err == nil {
		return true
	}

	// Check in current directory
	exePath := filepath.Join(exeDir, "ffprobe")
	if _, err := os.Stat(exePath); err == nil {
		return true
	}

	return false
}
<<<<<<< HEAD
<<<<<<< HEAD

func StopFFmpegNoPid() {
	var cmd *exec.Cmd

	// 根据操作系统选择合适的命令
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "Get-Process -Name ffmpeg | Select-Object -ExpandProperty Id")
	case "linux":
		cmd = exec.Command("pgrep", "-x", "ffmpeg")
	default:
		log.Println("Unsupported operating system")
		return
	}

	// 获取进程 PID
	pidBytes, err := cmd.Output()
	if err != nil {
		log.Println("Error finding ffmpeg process:", err)
		return
	}

	// 清理字符串中的空白字符
	pidStr := strings.TrimSpace(string(pidBytes))

	// 将 PID 转换为整数
	pidInt, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Println("Error converting PID to integer:", err)
		return
	}

	// 查找并终止 FFmpeg 进程
	log.Println("Stopping ffmpeg process...")
	process, err := os.FindProcess(pidInt)
	if err != nil {
		log.Println("Error finding ffmpeg process:", err)
		return
	}

	if err := process.Kill(); err != nil {
		log.Println("Error stopping ffmpeg process:", err)
		return
	}

	log.Println("FFmpeg process stopped.")
}

func StopFFmpegByPid(pid int) {

	// 查找并终止 FFmpeg 进程
	log.Println("Stopping ffmpeg process...")
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Println("Error finding ffmpeg process:", err)
		return
	}

	if err := process.Kill(); err != nil {
		log.Println("Error stopping ffmpeg process:", err)
		return
	}

	log.Println("FFmpeg process stopped.")
}
=======
>>>>>>> 57c1a529f170b5b38b9c894a33ef9114ee6767b0
=======
>>>>>>> remotes/origin/main
