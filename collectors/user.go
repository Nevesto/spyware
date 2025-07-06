package collectors

import "os/user"

func CollectUserInfo() (string, string) {
	currentUser, err := user.Current()
	if err != nil {
		return "", ""
	}
	return currentUser.Username, currentUser.HomeDir
}
