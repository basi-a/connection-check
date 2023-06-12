package main

import (
	"connection-check/cmd"

)

func main()  {
	Name, Type, _, _ := cmd.NetConnInfo()
	if Name =="ccdx-wifi" && Type == "wifi" {
		net_check := cmd.Ping()
		if !net_check{
			cmd.OpenBrowser()
		}

	}
}