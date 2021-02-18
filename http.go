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
		return fmt.Errorf("can't create new POST request %w: ", err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := p.http.Do(req)
	if err != nil {
		return fmt.Errorf("can't execute POST request %w: ", err)
	}

	defer func() {
		err = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad request")
	}

	err = json.NewDecoder(resp.Body).Decode(&value)
	if err != nil {
		return fmt.Errorf("can't decode POST request %w: ", err)
	}

	return nil
}

func (p *psn) get(url string, headers headers, value interface{}) error {
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return fmt.Errorf("can't create new GET request %w: ", err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := p.http.Do(req)
	if err != nil {
		return fmt.Errorf("can't execute GET request %w: ", err)
	}

	defer func() {
		err = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad request")
	}

	err = json.NewDecoder(resp.Body).Decode(&value)
	if err != nil {
		return fmt.Errorf("can't decode POST request %w: ", err)
	}

	return nil
}

func (p *psn) get(url string, headers headers, value interface{}) error {
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return fmt.Errorf("can't create new GET request %w: ", err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := p.http.Do(req)
	if err != nil {
		return fmt.Errorf("can't execute GET request %w: ", err)
	}

	defer func() {
		err = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(&value)
	if err != nil {
		return fmt.Errorf("can't decode GET request %w: ", err)
	}

	return nil
}
