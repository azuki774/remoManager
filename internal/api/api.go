package api

import "time"

type SensorValues struct {
	Hu     int
	Il     int
	Mo     int
	Te     float64
	Update time.Time
}

type SensorJson []struct {
	BtMacAddress    string    `json:"bt_mac_address"`
	CreatedAt       time.Time `json:"created_at"`
	FirmwareVersion string    `json:"firmware_version"`
	HumidityOffset  int       `json:"humidity_offset"`
	ID              string    `json:"id"`
	MacAddress      string    `json:"mac_address"`
	Name            string    `json:"name"`
	NewestEvents    struct {
		Hu SensorJsonAtVal `json:"hu"`
		Il SensorJsonAtVal `json:"il"`
		Mo SensorJsonAtVal `json:"mo"`
		Te struct {
			CreatedAt time.Time `json:"created_at"`
			Val       float64   `json:"val"`
		} `json:"te"`
	} `json:"newest_events"`
	SerialNumber      string    `json:"serial_number"`
	TemperatureOffset int       `json:"temperature_offset"`
	UpdatedAt         time.Time `json:"updated_at"`
	Users             []struct {
		ID        string `json:"id"`
		Nickname  string `json:"nickname"`
		Superuser bool   `json:"superuser"`
	} `json:"users"`
}
type SensorJsonAtVal struct {
	CreatedAt time.Time `json:"created_at"`
	Val       int       `json:"val"`
}

var nowSensorValues SensorValues

// Unmarshalしたデータからセンサー情報を取り出して格納する
func PickSensorValue(devicesJson SensorJson) (sensorvalue SensorValues) {
	sensorvalue.Hu = int(devicesJson[0].NewestEvents.Hu.Val)
	sensorvalue.Il = int(devicesJson[0].NewestEvents.Il.Val)
	sensorvalue.Mo = int(devicesJson[0].NewestEvents.Mo.Val)
	sensorvalue.Te = devicesJson[0].NewestEvents.Te.Val
	sensorvalue.Update = time.Now()
	return sensorvalue
}

// 現時点で最新のデータを返す
func GetSensorValues() SensorValues {
	return nowSensorValues
}
