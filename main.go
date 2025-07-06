package main

import (
	"fmt"

	"github.com/Nevesto/spyware/services"
)

func main() {
	systemInfo := services.CollectSystemInfo()
	services.SaveToFile(systemInfo)
	services.DisplayInfo(systemInfo) // debug
	fmt.Println("Program finished.")
}
