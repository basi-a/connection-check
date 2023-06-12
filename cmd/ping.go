package cmd

import "os/exec"

func Ping() bool {
	ch := make(chan error)
	go func ()  {
		cmd := exec.Command("ping", "-c", "1", "-w", "5", "baidu.com")
		err := cmd.Run()
		ch <- err
	}()
	err := <-ch
	return err == nil
}