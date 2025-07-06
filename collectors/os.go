package collectors

import "runtime"

func CollectOSInfo() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
