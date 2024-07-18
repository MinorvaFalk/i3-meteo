package meteo

import (
	"encoding/json"
	"fmt"
	"i3/config"
	"i3/internal/model"
	"net/http"
	"net/url"
	"time"
)

type Meteo struct {
	client *http.Client
	url    string
	key    string
}

func New() *Meteo {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Meteo{
		client: client,
		url:    config.ReadConfig().MeteoUrl,
		key:    config.ReadConfig().MeteoKey,
	}
}

func (m *Meteo) FetchLocationWeather(lat, lon string) (*model.Weather, error) {
	req, err := http.NewRequest(http.MethodGet, m.url+"/point", nil)
	if err != nil {
		return nil, err
	}

	q := m.setDefaultParams(req.URL.Query())
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()

	var data model.Weather
	if err := m.doRequest(req, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *Meteo) FetchAirQuality(lat, lon string) (*model.Air, error) {
	panic(fmt.Errorf("method not implemented"))
}

func (m *Meteo) FetchNearestPlace(lat, lon string) (*model.Place, error) {
	req, err := http.NewRequest(http.MethodGet, m.url+"/nearest_place", nil)
	if err != nil {
		return nil, err
	}

	q := m.setDefaultParams(req.URL.Query())
	q.Add("lat", lat)
	q.Add("lon", lat)
	req.URL.RawQuery = q.Encode()

	var data model.Place
	if err := m.doRequest(req, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *Meteo) doRequest(req *http.Request, v any) error {
	res, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed with status code: %v", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (m *Meteo) setDefaultParams(q url.Values) url.Values {
	q.Add("key", m.key)
	q.Add("units", "metric")
	q.Add("language", "en")

	return q
}
