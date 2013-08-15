// +build graphity

package neo4j

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"net/url"
)

type ManuelRequest struct {
	Method string
	To     string
	Params map[string]string
}

func (neo4j *Neo4j) Request(mr *ManuelRequest) error {
	params := url.Values{}
	for key, value := range mr.Params {
		params.Add(key, value)
	}

	finalUrl := fmt.Sprintf("%s?%s", mr.To, params.Encode())

	req, err := http.NewRequest(mr.Method, finalUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := neo4j.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 204 {
		return fmt.Errorf(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Println(string(body))

	return nil
}
