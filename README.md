connection-check
====

## 用途
当使用`Linux + 窗口管理器`时，系统不会自动检测所连接的网络是否需要认证，于是便有了这个
## 食用方法
```bash
git clone git@github.com:basi-a/connection-check.git
cd connection-check
go build && go install
```
然后添加到自启脚本中，以dwm为例
```bash
echo "${HOME}/go/bin/connection-check &" >> ${HOME}/.dwm/autostart.sh
```
## 工作方式
1. 检查网络连接方式是否是`wifi`, SSID是不是`ccdx-wifi`
2. 进行`ping`测试，若无错误则结束，否则继续
3. 若有错误，打开浏览器访问我们学校的校园网认证页

~~代码很少，很简陋就不写注释了~~