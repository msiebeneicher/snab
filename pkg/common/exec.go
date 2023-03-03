package common

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"snab/pkg/logger"
)

// Exec inits a command and execute it
func Exec(exec_cmd string, args []string, dir string, env string) {
	//cmd := newCmd(exec_cmd, args, env)
	cmd := exec.Command(
		exec_cmd,
		args...,
	)

	if dir != "" {
		cmd.Dir = dir
	}

	if env != "" {
		newEnv := append(os.Environ(), env)
		cmd.Env = newEnv
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	logger.WithField("args", args).Tracef("Exec '%s'", exec_cmd)
	err := cmd.Run()
	_, errStr := stdoutBuf.String(), stderrBuf.String()

	if err != nil {
		logger.WithField("err", err).WithField("errStr", errStr).Errorf("Error during execution of '%s'", exec_cmd)
		logger.Fatalf(string(errStr))
	}

	//fmt.Print(outStr)
	fmt.Println()

	// variant: Capture output via default CombinedOutput()
	//
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	logger.WithField("err", err).Errorf("Error during execution of '%s'", exec_cmd)
	// 	logger.Fatalf(string(out))
	// }
	//
	// fmt.Print(string(out))
}
