package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func Opener(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("브라우저 자동 실행 실패: %v\n", err)
	}
}
