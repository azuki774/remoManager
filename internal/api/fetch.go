package api

import (
	"encoding/json"
	"os/exec"

	"github.com/sirupsen/logrus"
)

type RemoAPI interface {
	getAppliances() (res interface{}, err error)
	getDevices() (res SensorJson, err error)
	postLightOperate() (res interface{}, err error)
}

type remoAPI struct {
	applianceId string
	token       string
}

type mockremoAPI struct {
}

func MakeRemoAPI(testMode bool, napplianceId string, ntoken string) RemoAPI {
	// testMode = Trueならダミー応答mockに接続
	if testMode {
		return &mockremoAPI{}
	}
	return &remoAPI{applianceId: napplianceId, token: ntoken}
}

// Raw Data -> Json (interface)
func (rapi *mockremoAPI) getAppliances() (res interface{}, err error) {
	output := `[
		{
		  "id": "b9676d66-a962-47ff-a35f-8553d50db74f",
		  "device": {
			"name": "Remo",
			"id": "fa6a01c6-ecd1-4c42-8a4a-8f6e5e9ab35a",
			"created_at": "2021-12-11T02:03:40Z",
			"updated_at": "2021-12-26T00:57:50Z",
			"mac_address": "ac:67:b2:f0:ae:cc",
			"bt_mac_address": "ac:67:b2:f0:ae:ce",
			"serial_number": "1W320100010231",
			"firmware_version": "Remo/1.6.4",
			"temperature_offset": 0,
			"humidity_offset": 0
		  },
		  "model": {
			"id": "4de59377-0e4a-4eb4-a73e-5c5fbbf740c4",
			"country": "JP",
			"manufacturer": "daikin",
			"remote_name": "arc478a52",
			"series": "",
			"name": "Daikin AC 074",
			"image": "ico_ac_1"
		  },
		  "type": "AC",
		  "nickname": "エアコン",
		  "image": "ico_ac_1",
		  "settings": {
			"temp": "21",
			"temp_unit": "c",
			"mode": "warm",
			"vol": "1",
			"dir": "1",
			"dirh": "still",
			"button": "",
			"updated_at": "2021-12-26T00:57:49Z"
		  },
		  "aircon": {
			"range": {
			  "modes": {
				"auto": {
				  "temp": [
					"-5",
					"-4.5",
					"-4",
					"-3.5",
					"-3",
					"-2.5",
					"-2",
					"-1.5",
					"-1",
					"-0.5",
					"0",
					"0.5",
					"1",
					"1.5",
					"2",
					"2.5",
					"3",
					"3.5",
					"4",
					"4.5",
					"5"
				  ],
				  "dir": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"swing"
				  ],
				  "dirh": [
					"still",
					"swing"
				  ],
				  "vol": [
					"1",
					"auto"
				  ]
				},
				"blow": {
				  "temp": [
					""
				  ],
				  "dir": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"swing"
				  ],
				  "dirh": [
					"still",
					"swing"
				  ],
				  "vol": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"auto"
				  ]
				},
				"cool": {
				  "temp": [
					"18",
					"18.5",
					"19",
					"19.5",
					"20",
					"20.5",
					"21",
					"21.5",
					"22",
					"22.5",
					"23",
					"23.5",
					"24",
					"24.5",
					"25",
					"25.5",
					"26",
					"26.5",
					"27",
					"27.5",
					"28",
					"28.5",
					"29",
					"29.5",
					"30",
					"30.5",
					"31",
					"31.5",
					"32"
				  ],
				  "dir": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"swing"
				  ],
				  "dirh": [
					"still",
					"swing"
				  ],
				  "vol": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"auto"
				  ]
				},
				"dry": {
				  "temp": [
					"-2",
					"-1.5",
					"-1",
					"-0.5",
					"0",
					"0.5",
					"1",
					"1.5",
					"2"
				  ],
				  "dir": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"swing"
				  ],
				  "dirh": [
					"still",
					"swing"
				  ],
				  "vol": [
					""
				  ]
				},
				"warm": {
				  "temp": [
					"14",
					"14.5",
					"15",
					"15.5",
					"16",
					"16.5",
					"17",
					"17.5",
					"18",
					"18.5",
					"19",
					"19.5",
					"20",
					"20.5",
					"21",
					"21.5",
					"22",
					"22.5",
					"23",
					"23.5",
					"24",
					"24.5",
					"25",
					"25.5",
					"26",
					"26.5",
					"27",
					"27.5",
					"28",
					"28.5",
					"29",
					"29.5",
					"30"
				  ],
				  "dir": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"swing"
				  ],
				  "dirh": [
					"still",
					"swing"
				  ],
				  "vol": [
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"auto"
				  ]
				}
			  },
			  "fixedButtons": [
				"power-off"
			  ]
			},
			"tempUnit": "c"
		  },
		  "signals": []
		},
		{
		  "id": "e8c7d05a-b07b-4565-a97d-02242d06525b",
		  "device": {
			"name": "Remo",
			"id": "fa6a01c6-ecd1-4c42-8a4a-8f6e5e9ab35a",
			"created_at": "2021-12-11T02:03:40Z",
			"updated_at": "2021-12-26T00:57:50Z",
			"mac_address": "ac:67:b2:f0:ae:cc",
			"bt_mac_address": "ac:67:b2:f0:ae:ce",
			"serial_number": "1W320100010231",
			"firmware_version": "Remo/1.6.4",
			"temperature_offset": 0,
			"humidity_offset": 0
		  },
		  "model": {
			"id": "5499e595-33f1-4bca-b48e-fc1a9b58cd74",
			"country": "JP",
			"manufacturer": "nec",
			"remote_name": "re0402-ch1",
			"name": "NEC LIGHT 001",
			"image": "ico_light"
		  },
		  "type": "LIGHT",
		  "nickname": "照明",
		  "image": "ico_light",
		  "settings": null,
		  "aircon": null,
		  "signals": [
			{
			  "id": "7b4758db-1dea-4771-b833-a4902a23596a",
			  "name": "3",
			  "image": "ico_number_3"
			}
		  ],
		  "light": {
			"buttons": [
			  {
				"name": "on",
				"image": "ico_on",
				"label": "Light_on"
			  },
			  {
				"name": "off",
				"image": "ico_off",
				"label": "Light_off"
			  },
			  {
				"name": "on-100",
				"image": "ico_light_all",
				"label": "Light_all"
			  },
			  {
				"name": "on-favorite",
				"image": "ico_light_favorite",
				"label": "Light_favorite"
			  },
			  {
				"name": "night",
				"image": "ico_light_night",
				"label": "Light_night"
			  },
			  {
				"name": "bright-up",
				"image": "ico_arrow_top",
				"label": "Light_bright"
			  },
			  {
				"name": "bright-down",
				"image": "ico_arrow_bottom",
				"label": "Light_dark"
			  },
			  {
				"name": "colortemp-down",
				"image": "ico_arrow_right",
				"label": "Light_warm"
			  },
			  {
				"name": "colortemp-up",
				"image": "ico_arrow_left",
				"label": "Light_cold"
			  }
			],
			"state": {
			  "brightness": "100",
			  "power": "on",
			  "last_button": "on-100"
			}
		  }
		}
	  ]

	  `
	outputByte := []byte(output)
	_ = json.Unmarshal(outputByte, &res)
	return res, nil
}

func (rapi *mockremoAPI) getDevices() (res SensorJson, err error) {
	output := `[
		{
		  "name": "Remo",
		  "id": "fa6a01c6-ecd1-4c42-8a4a-8f6e5e9ab35a",
		  "created_at": "2021-12-11T02:03:40Z",
		  "updated_at": "2021-12-26T00:57:50Z",
		  "mac_address": "ac:67:b2:f0:ae:cc",
		  "bt_mac_address": "ac:67:b2:f0:ae:ce",
		  "serial_number": "1W320100010231",
		  "firmware_version": "Remo/1.6.4",
		  "temperature_offset": 0,
		  "humidity_offset": 0,
		  "users": [
			{
			  "id": "12ee4f2f-a59a-4884-95bb-b75f7a9623c6",
			  "nickname": "Naruki Taniguchi",
			  "superuser": true
			}
		  ],
		  "newest_events": {
			"hu": {
			  "val": 36,
			  "created_at": "2021-12-26T03:15:30Z"
			},
			"il": {
			  "val": 44,
			  "created_at": "2021-12-26T02:51:46Z"
			},
			"mo": {
			  "val": 1,
			  "created_at": "2021-12-26T03:18:39Z"
			},
			"te": {
			  "val": 18.8,
			  "created_at": "2021-12-26T03:13:30Z"
			}
		  }
		}
	  ]
	  
	  `

	outputByte := []byte(output)
	_ = json.Unmarshal(outputByte, &res)
	return res, nil
}

func (rapi *mockremoAPI) postLightOperate() (res interface{}, err error) {
	output := `{"brightness":"100","power":"on","last_button":"on-100"}`
	outputByte := []byte(output)
	_ = json.Unmarshal(outputByte, &res)
	return output, nil
}

func (rapi *remoAPI) getAppliances() (res interface{}, err error) {
	return "", nil
}

func (rapi *remoAPI) getDevices() (res SensorJson, err error) {
	output, err := exec.Command("sh", "/getDevices.sh").Output()
	if err != nil {
		logrus.Errorf("getDevices error : %v", err)
		return nil, err
	}
	err = json.Unmarshal(output, &res)
	return res, err
}

func (rapi *remoAPI) postLightOperate() (res interface{}, err error) {
	return "", nil
}
