package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

// AddExecuteChmod 添加执行权限
func AddExecuteChmod(executePath string) ([]byte, error) {
	// 添加执行权限
	var chmodCmdString string
	switch runtime.GOOS {
	case "windows":
		return nil, nil
	case "darwin":
		chmodCmdString = fmt.Sprintf("chmod +x %s && xattr -cr %s", executePath, executePath)
	default:
		chmodCmdString = fmt.Sprintf("chmod +x %s", executePath)
	}
	chmodCmd := exec.Command("bash", "-c", chmodCmdString)
	return chmodCmd.CombinedOutput()
}
