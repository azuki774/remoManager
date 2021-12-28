package api

import (
	"fmt"
	"remoManager/internal/logger"
	"time"
)

// Remo Cloud APIに取りに行く頻度(s)
const FetchInterval = 60

// 指定時間ごとにデータを取得しに行き、nowSensorValuesの値を更新する
func RoutineFetchDevices(rapi RemoAPI) {
	for {
		devicesJson, err := rapi.getDevices()
		if err != nil {
			logger.FatalPrint("fetch device data error")
			time.Sleep(FetchInterval * time.Second)
			continue
		}

		nowSensorValues = PickSensorValue(devicesJson)
		fmt.Println(nowSensorValues)

		time.Sleep(FetchInterval * time.Second)
	}
}
