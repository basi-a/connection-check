package main

import (
	"connection-check/cmd"
)

func main()  {
	// 获取当前网络连接信息
	Name, Type, _, _ := cmd.NetConnInfo()
	// 如果连接名称为“ccdx-wifi”，并且连接类型为“wifi”
	if Name =="ccdx-wifi" && Type == "wifi" {
		// 使用Ping函数检测网络是否正常
		net_check := cmd.Ping()
		// 如果网络不正常
		if !net_check{
			// 使用OpenBrowser函数打开浏览器
			cmd.OpenBrowser()
		}

	}
}