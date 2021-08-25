package gutil

import (
	"bytes"
	"os/exec"
)

func CallScript(script string) (string, string, error) {
	cmd := exec.Command("/bin/bash", script)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr, nil
}
