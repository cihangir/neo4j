package neo4j

import (
	"errors"
	"fmt"
)

type Index struct {
	Name   string
	Config map[string]interface{}
}

func (neo4j *Neo4j) CreateIndex(index *Index) (bool, error) {

	if index.Name == "" {
		return false, errors.New("Name must be set!")
	}

	postData := ""

	if len(index.Config) > 0 {
		config, err := jsonEncode(index.Config)
		if err != nil {
			return false, err
		}
		postData = fmt.Sprintf(`{"name" : "%s", "config" : %s }`, index.Name, config)
	} else {
		postData = fmt.Sprintf(`{"name" : "%s" }`, index.Name)
	}

	_, err := neo4j.doRequest("POST", neo4j.IndexNodeUrl, postData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (neo4j *Neo4j) DeleteIndex(name string) (bool, error) {

	url := neo4j.IndexNodeUrl + "/" + name

	//if node not found Neo4j returns 404
	_, err := neo4j.doRequest("DELETE", url, "")
	if err != nil {
		return false, err
	}

	return true, nil
}
