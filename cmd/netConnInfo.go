package cmd

import (
	"bytes"
	"strings"

	"os/exec"
)

// 定义获取网络连接信息命令的常量
const(
	getNameCmd = "nmcli connection show --active | awk '{if (NR==2){print $1}}'"
	getTypeCmd = "nmcli connection show --active | awk '{if (NR==2){print $3}}'"
)

// NetConnInfo函数用于获取当前网络连接的名称和类型信息，并返回获取到的信息和错误信息
func NetConnInfo() (Name, Type, GetNameErr, GetTypeErr string) {
	ch := make(chan []string)

	// 获取网络连接名称
	go func ()  {
		cmd := exec.Command("/bin/bash", "-c", getNameCmd)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		
		cmd.Run()

		// 将输出流和错误流转换为字符串，并去除换行符
		outStr, errStr := strings.Replace(stdout.String(),"\n","", -1), strings.Replace(stderr.String(),"\n","", -1)
		ch <- []string{outStr, errStr}
	}()
	cmd_info1 := <-ch
	Name, GetNameErr = cmd_info1[0], cmd_info1[1]

	// 获取网络连接类型
	go func ()  {
		cmd := exec.Command("/bin/bash", "-c", getTypeCmd)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		
		cmd.Run()

		// 将输出流和错误流转换为字符串，并去除换行符
		outStr, errStr := strings.Replace(stdout.String(),"\n","", -1), strings.Replace(stderr.String(),"\n","", -1)
		ch <- []string{outStr, errStr}
	}()
	cmd_info2 := <-ch
	Type, GetTypeErr = cmd_info2[0], cmd_info2[1]
	
	// 返回获取到的网络连接名称和类型信息以及可能出现的错误信息
	return
}