package psn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type headers map[string]string

func (p *psn) post(formData url.Values, url string, headers headers, value interface{}) error {
	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		return fmt.Errorf("can't create new request %w: ", err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := p.http.Do(req)
	if err != nil {
		return fmt.Errorf("can't execute request %w: ", err)
	}

	defer func() {
		err = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(&value)
	if err != nil {
		return fmt.Errorf("can't decode request %w: ", err)
	}

	return nil
}
