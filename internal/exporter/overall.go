package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"remoManager/internal/api"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

const (
	namespace = "remoMetrics"
)

type myCollector struct{}

var (
	roomTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "room_temperature",
		Help:      "room temperature",
	})
	roomMotion = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "room_motion",
		Help:      "room motion sensor",
	})
	roomIllumi = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "room_illuminance",
		Help:      "room illuminance",
	})
	roomHumid = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "room_humidity",
		Help:      "room humidity",
	})
)

func GetSensorValueRoutine() {
	for {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		resp, err := client.Get("http://remo-manager/sensor")
		if err != nil {
			logrus.Errorf("Get sensor data error %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("Get sensor data read error %v", err)
		}

		fmt.Println(resp.Status)
		fmt.Println("body: " + string(body))
		refreshSensorValue(body)

		time.Sleep(5 * time.Second)
	}
}

func refreshSensorValue(body []byte) {
	var inputJson api.SensorValues
	err := json.Unmarshal(body, &inputJson)
	if err != nil {
		logrus.Errorf("refresh json error")
	}

	roomTemp.Set(inputJson.Te)
	roomIllumi.Set(float64(inputJson.Il))
	roomMotion.Set(float64(inputJson.Mo))
	roomHumid.Set(float64(inputJson.Hu))
}

func init() {
	prometheus.MustRegister(roomTemp)
	prometheus.MustRegister(roomIllumi)
	prometheus.MustRegister(roomMotion)
	prometheus.MustRegister(roomHumid)
}
