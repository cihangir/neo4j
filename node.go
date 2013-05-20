package neo4j

import (
	"encoding/json"
	"errors"
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

// gets node from neo4j with given unique node id
// response will be object
func (neo4j *Neo4j) GetNode(id string) (*Node, error) {

	url := neo4j.NodeUrl + "/" + id

	//if node not found Neo4j returns 404
	response, err := neo4j.doRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	node := &Node{}

	payload, err := node.decodeResponse(response)
	if err != nil {
		return nil, err
	}

	node.Id = id
	node.Data = payload.Data
	node.Payload = payload

	return node, nil
}

// creates a unique node with given id and node name
// response will be Object
func (neo4j *Neo4j) CreateNode(node *Node) (bool, error) {

	postData, err := node.encodeData()
	if err != nil {
		return false, err
	}

	response, err := neo4j.doRequest("POST", neo4j.NodeUrl, postData)
	if err != nil {
		return false, err
	}

	payload, err := node.decodeResponse(response)
	if err != nil {
		return false, err
	}

	id, err := getIdFromUrl(neo4j.NodeUrl, payload.Self)

	node.Id = id
	node.Data = payload.Data
	node.Payload = payload

	return true, nil
}

// creates a unique node with given id and node name
// response will be Object
func (neo4j *Neo4j) UpdateNode(node *Node) (bool, error) {

	if node.Id == "" {
		return false, errors.New("Invalid node id")
	}

	postData, err := node.encodeData()
	if err != nil {
		return false, err
	}

	url := neo4j.NodeUrl + "/" + node.Id + "/properties"

	_, err = neo4j.doRequest("PUT", url, postData)
	if err != nil {
		return false, err
	}

	return true, nil
}

// gets node from neo4j with given unique node id
// response will be object
// todo add force deletion, because nodes with relationships can not be deleted
func (neo4j *Neo4j) DeleteNode(id string) (bool, error) {

	url := neo4j.NodeUrl + "/" + id

	//if node not found Neo4j returns 404
	_, err := neo4j.doRequest("DELETE", url, "")
	if err != nil {
		return false, err
	}

	return true, nil
}

func (node *Node) encodeData() (string, error) {
	result, err := jsonEncode(node.Data)
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
