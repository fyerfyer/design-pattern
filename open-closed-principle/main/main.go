package main

import (
	"alert/alert"
	"fmt"
)

func main() {
	appContext := alert.NewApplicationContext()
	appContext.InitializeBeans()

	apiStatInfo := &alert.ApiStatInfo{
		API:               "SomeAPI",
		RequestCount:      1000,
		ErrorCount:        10,
		DurationOfSeconds: 100,
	}

	apiStatInfo.SetTimeoutCount(289)

	appContext.GetAlert().Check(*apiStatInfo)

	fmt.Println("Alert system triggered based on the given ApiStatInfo.")
}
