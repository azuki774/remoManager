package main

import (
	"os"
	"remoManager/internal/api"
	"remoManager/internal/logger"
	"remoManager/internal/server"
)

func main() {
	logger.InfoPrint("remoManager start")
	var rapi api.RemoAPI

	if os.Getenv("TOKEN") == "" {
		logger.FatalPrint("TOKEN is not set.")
		os.Exit(1)
	}

	logger.InfoPrint("Use Appliance ID : " + os.Getenv("APPLIANCE_ID"))
	logger.InfoPrint("Use Token : " + os.Getenv("TOKEN"))

	var testMode bool = true
	if os.Getenv("API_ENV") == "prd" {
		testMode = false
		logger.InfoPrint("Use actual Nature Remo API")
	}

	rapi = api.MakeRemoAPI(testMode, os.Getenv("APPLIANCE_ID"), os.Getenv("TOKEN"))
	go api.RoutineFetchDevices(rapi)
	server.ServerStart()
}
