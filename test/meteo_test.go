package test

import (
	"encoding/json"
	"i3/config"
	"i3/internal/meteo"
	"testing"
)

func TestMeteo(t *testing.T) {
	config.InitConfig("../.env")

	m := meteo.New()
	res, err := m.FetchLocationWeather("3.6422714", "98.5043403")
	if err != nil {
		t.Error(err)
	}

	b, _ := json.Marshal(res)
	t.Log(string(b))
}
