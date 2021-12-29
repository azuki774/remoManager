package main

import (
	"remoManager/internal/exporter"
	"remoManager/internal/logger"
)

func main() {
	logger.InfoPrint("remoManager start")
	go exporter.GetSensorValueRoutine()
	exporter.ServerStart()
}
