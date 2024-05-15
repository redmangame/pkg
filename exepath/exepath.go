<<<<<<< HEAD
/*
 * @Author: redmangame redmangame@163.com
 * @Date: 2024-05-14 08:53:01
 * @LastEditors: redmangame redmangame@163.com
 * @LastEditTime: 2024-05-14 08:55:25
 * @FilePath: \fyneguiffmpeg\pkg\exepath\exepath.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
=======
>>>>>>> remotes/origin/main
package exepath

import (
	"log"
	"os"
	"path/filepath"
)

<<<<<<< HEAD
var ExeDir string

func init() {
	ExeDir = GetCurrentDir()
}

=======
>>>>>>> remotes/origin/main
func GetCurrentDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	currentDir := filepath.Dir(exePath)
	return currentDir
}
