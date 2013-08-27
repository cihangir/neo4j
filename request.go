package neo4j

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ManuelRequest struct {
	Method string
	To     string
	Params map[string]string
}

func (neo4j *Neo4j) Request(mr *ManuelRequest) ([]string, error) {
	params := url.Values{}
	for key, value := range mr.Params {
		params.Add(key, value)
	}

	finalUrl := fmt.Sprintf("%s?%s", mr.To, params.Encode())

	req, err := http.NewRequest(mr.Method, finalUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := neo4j.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 && res.StatusCode != 204 {
		return nil, fmt.Errorf(res.Status)
	}

	if res.StatusCode == 200 {
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
	}

	return nil, nil
}
