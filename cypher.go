package neo4j

import "encoding/json"

type Cypher struct {
	Query  string
	Params map[string]interface{}
}

type CypherResponse struct {
	Columns map[string]interface{} `json:"columns"`
	Data    map[string]interface{} `json:"data"`
}

func (node *Node) mapBatchResponse(neo4j *Neo4j, data interface{}) (bool, error) {
	encodedData, err := jsonEncode(data)
	payload, err := node.decodeResponse(encodedData)
	if err != nil {
		return false, err
	}
	id, err := getIdFromUrl(neo4j.NodeUrl, payload.Self)
	if err != nil {
		return false, nil
	}
	node.Id = id
	node.Data = payload.Data
	node.Payload = payload

	return true, nil
}

func (c *Cypher) getBatchQuery(operation string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"method": "POST",
		"to":     "/cypher",
		"body":   c,
	}, nil
}

func (n *Node) encodeData() (string, error) {
	result, err := jsonEncode(n.Data)
	return result, err
}

func (node *Node) decodeResponse(data string) (*NodeResponse, error) {
	nodeResponse := &NodeResponse{}

	err := json.Unmarshal([]byte(data), nodeResponse)
	if err != nil {
		return nil, err
	}

	return nodeResponse, nil
}
