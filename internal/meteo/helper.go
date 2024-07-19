package meteo

import (
	"context"
	"encoding/json"
	"fmt"
	"i3/pkg/logger"
	"net/http"
	"time"
)

func (m *meteo) fetchMeteoData(redisKey, path string, params map[string]string, v any) error {
	// Check for redis cache
	if err := m.checkRedisCache(redisKey, v); err == nil {
		return nil
	}

	// Create HTTP Request
	req, err := http.NewRequest(http.MethodGet, m.url+path, nil)
	if err != nil {
		return nil
	}

	q := req.URL.Query()
	q.Add("key", m.key)
	q.Add("language", "en")
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	exp, err := m.doRequest(req, v)
	if err != nil {
		return err
	}

	m.mustSaveRedisCache(redisKey, v, exp)

	return nil
}

func (m *meteo) doRequest(req *http.Request, v any) (expiresIn time.Duration, err error) {
	res, err := m.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("http request failed with status code: %v", res.StatusCode)
		return
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return
	}

	ts, err := time.Parse(http.TimeFormat, res.Header.Get("expires"))
	if err != nil {
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	local := ts.In(loc)
	expiresIn = time.Until(local)

	return
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
