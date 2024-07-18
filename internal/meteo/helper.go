package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (m *meteo) setDefaultParams(q url.Values) url.Values {
	q.Add("key", m.key)
	q.Add("units", "metric")
	q.Add("language", "en")

	return q
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
