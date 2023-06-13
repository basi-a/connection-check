package cmd

import (
	"os/exec"
)

// 定义常量，其中AuthenticationURL为认证地址
const (
    AuthenticationURL = "http://1.1.1.2"
    chrome      = "google-chrome-stable"
    firefox     = "firefox"
    chromium    = "chromium"
    edge        = "microsoft-edge-stable"
)

// 定义浏览器选择顺序
var BrowserChoices = []string{edge, firefox, chromium, chrome}

// OpenBrowser函数用于打开浏览器并访问指定URL，返回打开浏览器的所有错误
func OpenBrowser() []error{
	// 创建一个有缓存通道
	ch := make(chan []error)

	go func ()  {
		var errs []error
		// 循环检查浏览器是否存在，若不存在则记录错误
		for _, browser := range BrowserChoices {
			_, err := exec.LookPath(browser)
			if err != nil {
				errs = append(errs, err)
			}else {
				// 创建一个打开指定URL的命令，并执行
				cmd := exec.Command(browser, AuthenticationURL)
				err := cmd.Run()
				if err != nil {
					errs = append(errs, err)
				}
				break
			}
		}
		// 将所有错误信息发送到通道中
		ch <- errs
	}()
	// 从通道中获取所有错误信息，并返回
	err := <-ch
	return err
}
