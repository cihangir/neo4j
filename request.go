package neo4j

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"
	"strings"
)

type ManuelRequest struct {
	Method string
	To     string
	Params map[string]string
	Body   map[string]string
}

func (mr *ManuelRequest) Get() ([]string, error) {
	urlWithParams := mr.encodeParams()
	req, err := http.NewRequest(mr.Method, urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := mr.decodeResponse(res)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mr *ManuelRequest) Post() error {
	body, err := jsonEncode(mr.Body)
	if err != nil {
		return err
	}

	jsonBody := strings.NewReader(body)
	req, err := http.NewRequest(mr.Method, mr.To, jsonBody)
	if err != nil {
		return err
	}

	mr.encodeForm(req)
	client := &http.Client{}
	res, err := client.PostForm(mr.To, req.Form)
	if err != nil {
		return err
	}

	_, err = mr.decodeResponse(res)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (mr *ManuelRequest) decodeResponse(res *http.Response) ([]string, error) {
	switch res.StatusCode {
	case 200, 500:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		result := make([]string, 0)
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	case 204:
	default:
		return nil, fmt.Errorf(res.Status)
	}

	return nil, nil
}

func (mr *ManuelRequest) encodeParams() string {
	var urlWithParams string

	if mr.Params != nil {
		params := url.Values{}
		for key, value := range mr.Params {
			params.Add(key, value)
		}

		urlWithParams = fmt.Sprintf("%s?%s", mr.To, params.Encode())
	} else {
		urlWithParams = mr.To
	}

	return urlWithParams
}

func (mr *ManuelRequest) encodeForm(req *http.Request) {
	if mr.Body != nil {
		req.Form = url.Values{}

		for k, v := range mr.Body {
			req.Form.Add(k, v)
		}
	}
}
