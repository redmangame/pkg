/*
使用 GetProcessID 函数获取特定应用程序的 PID。例如：

pid, err := processutil.GetProcessID("firefox")
if err != nil {
    log.Fatalf("Error getting process ID: %v", err)
}
fmt.Printf("PID of firefox: %d\n", pid)
这将输出 Firefox 浏览器的 PID。确保将 "firefox" 替换为你想要查找的进程名称。

使用 GetProcessPath 函数获取特定 PID 的应用程序路径。例如：

path, err := processutil.GetProcessPath(pid)
if err != nil {
    log.Fatalf("Error getting process path: %v", err)
}
fmt.Printf("Path of process with PID %d: %s\n", pid, path)
这将输出具有指定 PID 的进程的路径。

*/
package processutil

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// GetProcessID 根据应用程序名称获取PID
func GetProcessID(processName string) (int, error) {
	switch runtime.GOOS {
	case "windows":
		return GetProcessIDWindows(processName)
	case "linux":
		return GetProcessIDLinux(processName)
	default:
		return 0, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// GetProcessID 根据PID获取应用程序名称
func GetProcessName(pid int) (string, error) {
	switch runtime.GOOS {
	case "windows":
		return GetProcessPathWindows(pid)
	case "linux":
		return GetProcessPathLinux(pid)
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// GetProcessIDLinux 根据应用程序名称获取PID（Linux系统）
func GetProcessIDLinux(processName string) (int, error) {
	cmd := exec.Command("pgrep", "-x", processName)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	pidStr := strings.TrimSpace(string(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, err
	}

	return pid, nil
}

// GetProcessPathLinux 根据PID获取应用程序路径（Linux系统）
func GetProcessPathLinux(pid int) (string, error) {
	procPath := fmt.Sprintf("/proc/%d/exe", pid)
	path, err := os.Readlink(procPath)
	if err != nil {
		return "", err
	}
	return path, nil
}

// GetProcessIDWindows 根据应用程序名称获取PID（Windows系统）
func GetProcessIDWindows(processName string) (int, error) {
	cmd := exec.Command("tasklist", "/fo", "csv", "/nh")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, processName) {
			fields := strings.Split(line, ",")
			pidStr := strings.TrimSpace(fields[1])
			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				return 0, err
			}
			return pid, nil
		}
	}

	return 0, fmt.Errorf("process %s not found", processName)
}

// GetProcessPathWindows 根据PID获取应用程序路径（Windows系统）
func GetProcessPathWindows(pid int) (string, error) {
	cmd := exec.Command("wmic", "process", "where", fmt.Sprintf("processid=%d", pid), "get", "ExecutablePath", "/format:list")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parsing the output to get the path
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ExecutablePath=") {
			return strings.TrimSpace(strings.TrimPrefix(line, "ExecutablePath=")), nil
		}
	}

	return "", fmt.Errorf("executable path not found for PID %d", pid)
}
