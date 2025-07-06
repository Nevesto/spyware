package collectors

import "os"

func CollectHostInfo() (string, string) {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	return hostname, wd
}
