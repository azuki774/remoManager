package api

import (
	"reflect"
	"testing"
)

func Test_pickSensorValue(t *testing.T) {
	// テスト用mockからjson取得
	var devJson SensorJson
	rapi := MakeRemoAPI(true, "", "")
	devJson, _ = rapi.getDevices()

	type args struct {
		devicesJson SensorJson
	}
	tests := []struct {
		name            string
		args            args
		wantSensorvalue SensorValues
	}{
		{
			name:            "test #1",
			args:            args{devicesJson: devJson},
			wantSensorvalue: SensorValues{Hu: 36, Il: 44, Mo: 1, Te: 18.8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSensorvalue := PickSensorValue(tt.args.devicesJson)
			tt.wantSensorvalue.Update = gotSensorvalue.Update
			if !reflect.DeepEqual(gotSensorvalue, tt.wantSensorvalue) {
				t.Errorf("pickSensorValue() = %v, want %v", gotSensorvalue, tt.wantSensorvalue)
			}
		})
	}
}
