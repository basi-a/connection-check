package cmd

import (
	"bytes"
	"strings"

	"os/exec"
)
const(
	getNameCmd = "nmcli connection show --active | awk '{if (NR==2){print $1}}'"
	getTypeCmd = "nmcli connection show --active | awk '{if (NR==2){print $3}}'"
)
func NetConnInfo() (Name, Type, GetNameErr, GetTypeErr string) {
	ch := make(chan []string)
	go func ()  {
		cmd := exec.Command("/bin/bash", "-c", getNameCmd)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		
		cmd.Run()

		outStr, errStr := strings.Replace(stdout.String(),"\n","", -1), strings.Replace(stderr.String(),"\n","", -1)
		ch <- []string{outStr, errStr}
	}()
	cmd_info1 := <-ch
	Name, GetNameErr = cmd_info1[0], cmd_info1[1]
	go func ()  {
		cmd := exec.Command("/bin/bash", "-c", getTypeCmd)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		
		cmd.Run()

		outStr, errStr := strings.Replace(stdout.String(),"\n","", -1), strings.Replace(stderr.String(),"\n","", -1)
		ch <- []string{outStr, errStr}
	}()
	cmd_info2 := <-ch
	Type, GetTypeErr = cmd_info2[0], cmd_info2[1]
	return
}