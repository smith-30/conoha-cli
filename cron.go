package main

import (
	"runtime"

	"os/exec"

	"os"

	"github.com/carlescere/scheduler"
	"github.com/smith-30/conoha-cli/subcmd"
	"go.uber.org/zap"
	"github.com/smith-30/conoha-cli/helper/env"
)

func main() {
	env.LoadEnv()
	logger, _ := zap.NewDevelopment()

	boot := func() {
		out, err := exec.Command(os.Getenv("CMD"), subcmd.BOOT).Output()
		logger.Sugar().Info(out)
		logger.Sugar().Info(err)
	}

	halt := func() {
		out, err := exec.Command(os.Getenv("CMD"), subcmd.HALT).Output()
		logger.Sugar().Info(out)
		logger.Sugar().Info(err)
	}

	scheduler.Every().Day().At("19:30").Run(boot)
	scheduler.Every().Day().At("01:05").Run(halt)

	// Keep the program from not exiting.
	runtime.Goexit()
}
