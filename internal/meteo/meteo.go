package meteo

import (
	"context"
	"encoding/json"
	"i3/config"
	"i3/pkg/datasource"
	"i3/pkg/logger"
	"net/http"
	"time"
)

type Meteo interface {
	FetchLocationWeather(id, lat, lon string) (*Weather, error)
	FetchNearestPlace(lat, lon string) (*Place, error)
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
	if err := m.checkRedisCache(id, &data); err == nil {
		return &data, nil
	}

	req, err := http.NewRequest(http.MethodGet, m.url+"/point", nil)
	if err != nil {
		return nil, err
	}

	q := m.setDefaultParams(req.URL.Query())
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()

	exp, err := m.doRequest(req, &data)
	if err != nil {
		return nil, err
	}

	m.mustSaveRedisCache(id, data, exp)

	return &data, nil
}

func (m *meteo) FetchNearestPlace(lat, lon string) (*Place, error) {
	req, err := http.NewRequest(http.MethodGet, m.url+"/nearest_place", nil)
	if err != nil {
		return nil, err
	}

	q := m.setDefaultParams(req.URL.Query())
	q.Add("lat", lat)
	q.Add("lon", lat)
	req.URL.RawQuery = q.Encode()

	var data Place
	_, err = m.doRequest(req, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *meteo) checkRedisCache(key string, v any) error {
	res, err := m.redis.Get(context.Background(), key)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(res), v); err != nil {
		return err
	}

	return nil
}

func (m *meteo) mustSaveRedisCache(key string, v any, exp time.Duration) {
	if _, err := m.redis.Set(context.Background(), key, v, exp); err != nil {
		logger.Zap().Sugar().Error(err)
	}
}
