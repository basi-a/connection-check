package cmd

import "os/exec"


// Ping函数用于检测网络是否正常，返回值类型为bool
func Ping() bool {
	// 创建一个有缓存通道
	ch := make(chan error)
	go func ()  {
		// 创建一个ping指定地址的命令
		cmd := exec.Command("ping", "-c", "1", "-w", "5", "baidu.com")
		// 运行cmd指令，将错误信息发送到通道中
		err := cmd.Run()
		ch <- err
	}()
	// 从通道中获取错误信息
	err := <-ch
	// 如果错误信息为nil，表示网络正常，返回true
	// 否则返回false
	return err == nil
}