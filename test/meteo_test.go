package test

import (
	"encoding/json"
	"i3/config"
	"i3/internal/meteo"
	"i3/pkg/datasource"
	"testing"
)

func init() {
	config.InitConfig("../.env")
}

func TestMeteo(t *testing.T) {
	redis := datasource.NewRedis()
	m := meteo.New(redis)

	id := "a1a739d4-a9fc-4951-af12-04430fc20dff"
	lat := "3.6422714"
	lon := "98.5043403"

	res, err := m.FetchLocationWeather(id, lat, lon)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.Marshal(res)
	t.Log(string(b))
}
