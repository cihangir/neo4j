package neo4j

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Node struct {
	Id      string
	Data    map[string]interface{}
	Payload *NodeResponse
}

type NodeResponse struct {
	PagedTraverse              string                 `json:"paged_traverse"`
	OutgoingRelationships      string                 `json:"outgoing_relationships"`
	Traverse                   string                 `json:"traverse"`
	AllTypedRelationships      string                 `json:"all_typed_relationships"`
	Property                   string                 `json:"property"`
	AllRelationships           string                 `json:"all_relationships"`
	Self                       string                 `json:"self"`
	Properties                 string                 `json:"properties"`
	OutgoingTypedRelationships string                 `json:"outgoing_typed_relationships"`
	IncomingRelationships      string                 `json:"incoming_relationships"`
	IncomingTypedRelationships string                 `json:"incoming_typed_relationships"`
	CreateRelationship         string                 `json:"create_relationship"`
	Data                       map[string]interface{} `json:"data"`
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

func (node *Node) getBatchQuery(operation string) (map[string]interface{}, error) {
	switch operation {
	case BATCH_GET:
		query, err := prepareNodeGetBatchMap(node)
		return query, err
	case BATCH_UPDATE:
		query, err := prepareNodeUpdateBatchMap(node)
		return query, err
	case BATCH_CREATE:
		query, err := prepareNodeCreateBatchMap(node)
		return query, err
	case BATCH_DELETE:
		query, err := prepareNodeDeleteBatchMap(node)
		return query, err
	case BATCH_CREATE_UNIQUE:
		query, err := prepareNodeCreateUniqueBatchMap(node)
		return query, err
	default:
		return map[string]interface{}{}, nil
	}
}

func prepareNodeGetBatchMap(n *Node) (map[string]interface{}, error) {
	query := make(map[string]interface{})

	if n.Id == "" {
		return query, errors.New("Id field is empty")
	}

	query["method"] = "GET"
	query["to"] = fmt.Sprintf("/node/%s", n.Id)

	return query, nil
}

func prepareNodeDeleteBatchMap(n *Node) (map[string]interface{}, error) {
	query := make(map[string]interface{})

	if n.Id == "" {
		return query, errors.New("Id not valid")
	}

	query["method"] = "DELETE"
	query["to"] = fmt.Sprintf("/node/%s", n.Id)

	return query, nil
}

func prepareNodeCreateBatchMap(node *Node) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	query["method"] = "POST"
	query["to"] = "/node"
	query["body"] = node.Data

	return query, nil
}

func prepareNodeUpdateBatchMap(n *Node) (map[string]interface{}, error) {
	query := make(map[string]interface{})

	if n.Id == "" {
		return query, errors.New("Id not valid")
	}

	query["method"] = "PUT"
	query["to"] = fmt.Sprintf("/node/%s/properties", n.Id)
	query["body"] = n.Data

	return query, nil
}

func prepareNodeCreateUniqueBatchMap(node *Node) (map[string]interface{}, error) {

	query := make(map[string]interface{})
	query["method"] = "POST"
	query["to"] = "/index/node"
	query["body"] = map[string]interface{}{
		"properties": node.Data,
	}

	return query, nil
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
