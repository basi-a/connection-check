package cmd

import (
	"os/exec"
)

const (
    AuthenticationURL = "http://1.1.1.2"
    chrome      = "google-chrome-stable"
    firefox     = "firefox"
    chromium    = "chromium"
    edge        = "microsoft-edge-stable"
)

var BrowserChoices = []string{edge, firefox, chromium, chrome}

func OpenBrowser() []error{
	ch := make(chan []error)
	
	go func ()  {
		var errs []error
		for _, browser := range BrowserChoices {
			_, err := exec.LookPath(browser)
			if err != nil {
				errs = append(errs, err)
			}else {
				cmd := exec.Command(browser, AuthenticationURL)
				err := cmd.Run()
				if err != nil {
					errs = append(errs, err)
				}
				break
			}
		}
		ch <- errs
	}()
	err := <-ch
	return err
}
