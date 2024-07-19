package test

import (
	"i3/config"
	"i3/internal/meteo"
	"i3/pkg/datasource"
	"testing"

	"github.com/stretchr/testify/assert"
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
	text := "jakarta"
	textPrefix := "jak"

	t.Run("Weathers", func(t *testing.T) {
		res, err := m.FetchLocationWeather(id, lat, lon)
		if err != nil {
			t.Error(err)
		}

		assert.NotNil(t, res)
	})

	t.Run("Nearest Place", func(t *testing.T) {
		res, err := m.FetchNearestPlace(id, lat, lon)
		if err != nil {
			t.Error(err)
		}

		assert.NotNil(t, res)
	})

	t.Run("Places Prefix", func(t *testing.T) {
		res, err := m.FetchPlacesPrefix(id, textPrefix)
		if err != nil {
			t.Error(err)
		}

		assert.NotNil(t, res)
	})

	t.Run("Places", func(t *testing.T) {
		res, err := m.FetchPlaces(id, text)
		if err != nil {
			t.Error(err)
		}

		assert.NotNil(t, res)
	})

}
