package meteo

import (
	"i3/config"
	"i3/pkg/datasource"
	"net/http"
	"time"
)

type Meteo interface {
	FetchLocationWeather(id, lat, lon string) (*Weather, error)
	FetchNearestPlace(id, lat, lon string) (*Place, error)
	FetchPlacesPrefix(id, text string) (Places, error)
	FetchPlaces(id, text string) (Places, error)
}

type meteo struct {
	client *http.Client
	url    string
	key    string
	redis  datasource.Redis
}

func New(redis datasource.Redis) Meteo {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &meteo{
		client: client,
		url:    config.ReadConfig().MeteoUrl,
		key:    config.ReadConfig().MeteoKey,
		redis:  redis,
	}
}

func (m *meteo) FetchLocationWeather(id, lat, lon string) (*Weather, error) {
	var data Weather
	params := map[string]string{
		"lat":   lat,
		"lon":   lon,
		"units": "metric",
	}

	if err := m.fetchMeteoData(id+":weather", "/point", params, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *meteo) FetchNearestPlace(id, lat, lon string) (*Place, error) {
	var data Place
	params := map[string]string{
		"lat": lat,
		"lon": lon,
	}

	if err := m.fetchMeteoData(id+":nearest_place", "/nearest_place", params, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *meteo) FetchPlacesPrefix(id, text string) (Places, error) {
	var data Places
	params := map[string]string{
		"text": text,
	}

	if err := m.fetchMeteoData(id+":places_prefix", "/find_places_prefix", params, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *meteo) FetchPlaces(id, text string) (Places, error) {
	var data Places
	params := map[string]string{
		"text": text,
	}

	if err := m.fetchMeteoData(id+":places", "/find_places", params, &data); err != nil {
		return nil, err
	}

	return data, nil
}
