package neo4j

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Relationship struct {
	Id          string
	StartNodeId string
	EndNodeId   string
	Type        string
	Data        map[string]interface{}
	Payload     *RelationshipResponse
}

type RelationshipResponse struct {
	Start      string                 `json:"start"`
	Property   string                 `json:"property"`
	Self       string                 `json:"self"`
	Properties string                 `json:"properties"`
	Type       string                 `json:"type"`
	End        string                 `json:"end"`
	Data       map[string]interface{} `json:"data"`
}

// gets Relationship from neo4j with given unique Relationship id
// response will be object
func (neo4j *Neo4j) GetRelationship(id string) (*Relationship, error) {

	url := neo4j.RelationshipUrl + "/" + id

	response, err := neo4j.doRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	relationship := &Relationship{}

	_, err = relationship.decode(neo4j, response)
	if err != nil {
		return nil, err
	}

	return relationship, nil
}

// creates a unique Relationship with struct
func (neo4j *Neo4j) CreateRelationship(relationship *Relationship) (bool, error) {

	if relationship.StartNodeId == "" {
		return false, errors.New("Start Node Id not valid")
	}

	if relationship.EndNodeId == "" {
		return false, errors.New("End Node Id not valid")
	}

	url := fmt.Sprintf("%s/%s/relationships", neo4j.NodeUrl, relationship.StartNodeId)

	endNodeUrl := fmt.Sprintf("%s/%s", neo4j.NodeUrl, relationship.EndNodeId)

	relData, err := relationship.encodeData()
	if err != nil {
		return false, err
	}

	postData := fmt.Sprintf(`{"to" : "%s", "type" : "%s", "data" : %s }`, endNodeUrl, relationship.Type, relData)
	response, err := neo4j.doRequest("POST", url, postData)
	if err != nil {
		return false, err
	}

	result, err := relationship.decode(neo4j, response)
	if err != nil {
		return false, err
	}

	return result, nil
}

// creates a unique Relationship with given id and Relationship name
// response will be Object
func (neo4j *Neo4j) UpdateRelationship(relationship *Relationship) (bool, error) {

	if relationship.Id == "" {
		return false, errors.New("Invalid Relationship id")
	}

	postData, err := relationship.encodeData()
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf("%s/%s/properties", neo4j.RelationshipUrl, relationship.Id)

	_, err = neo4j.doRequest("PUT", url, postData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (neo4j *Neo4j) DeleteRelationship(id string) (bool, error) {

	url := fmt.Sprintf("%s/%s", neo4j.RelationshipUrl, id)
	//if Relationship not found Neo4j returns 404
	_, err := neo4j.doRequest("DELETE", url, "")
	if err != nil {
		return false, err
	}

	return true, nil
}

func (neo4j *Neo4j) GetRelationshipTypes() ([]string, error) {

	url := fmt.Sprintf("%s/types", neo4j.RelationshipUrl)
	result := make([]string, 0)
	response, err := neo4j.doRequest("GET", url, "")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(response), &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func (relationship *Relationship) mapBatchResponse(neo4j *Neo4j, data map[string]interface{}) (bool, error) {
	// because data is a map, convert back to Json
	encodedData, err := jsonEncode(data)
	result, err := relationship.decode(neo4j, data)

	return result, err
}

func (relationship *Relationship) getBatchQuery(operation string) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	switch operation {
	case BATCH_GET:
		query, err := prepareRelationshipGetBatchMap(relationship)
		return query, err
	case BATCH_UPDATE:
		query, err := prepareRelationshipUpdateBatchMap(relationship)
		return query, err
	case BATCH_CREATE:
		query, err := prepareRelationshipCreateBatchMap(relationship)
		return query, err
	case BATCH_DELETE:
		query, err := prepareRelationshipDeleteBatchMap(relationship)
		return query, err
	}
	return query, nil
}

func prepareRelationshipGetBatchMap(relationship *Relationship) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	if relationship.Id == "" {
		return query, errors.New("Id not valid")
	}

	query["method"] = "GET"
	query["to"] = "/relationship/" + relationship.Id

	return query, nil
}

func prepareRelationshipDeleteBatchMap(relationship *Relationship) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	if relationship.Id == "" {
		return query, errors.New("Id not valid")
	}

	query["method"] = "DELETE"
	query["to"] = "/relationship/" + relationship.Id

	return query, nil
}

func prepareRelationshipCreateBatchMap(relationship *Relationship) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	if relationship.StartNodeId == "" {
		return query, errors.New("Start Node Id not valid")
	}

	if relationship.EndNodeId == "" {
		return query, errors.New("End Node Id not valid")
	}

	if relationship.Type == "" {
		return query, errors.New("Relationship type is not valid")
	}

	url := "/node/" + relationship.StartNodeId + "/relationships"

	endNodeUrl := "/node/" + relationship.EndNodeId

	query["method"] = "POST"
	query["to"] = url

	body := make(map[string]interface{})
	body["to"] = endNodeUrl
	body["type"] = relationship.Type
	body["data"] = relationship.Data
	query["body"] = body
	return query, nil
}

func prepareRelationshipUpdateBatchMap(relationship *Relationship) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	if relationship.Id == "" {
		return query, errors.New("Id not valid")
	}

	query["method"] = "PUT"
	query["to"] = "/node/" + relationship.Id + "/properties"
	query["body"] = relationship.Data
	return query, nil
}

// func (neo4j *Neo4j) GetAllRelationships(id string) (*[]Relationship, error) {

// 	url := neo4j.NodeUrl + "/" + id + "/relationships/all"

// 	//if node not found Neo4j returns 404
// 	response, err := neo4j.doRequest("GET", url, "")
// 	if err != nil {
// 		return false, err
// 	}

// 	relationships := &[]Relationship{}

// 	result, err := relationships.decodeArray(response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return relationships, nil

// }

func (relationship *Relationship) encodeData() (string, error) {
	result, err := jsonEncode(relationship.Data)
	return result, err
}

func (relationship *Relationship) decode(neo4j *Neo4j, data string) (bool, error) {

	payload := &RelationshipResponse{}

	// Map json to our RelationshipResponse struct
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		return false, err
	}

	// Map returning result to our relationship struct
	err = mapRelationship(neo4j, relationship, payload)
	if err != nil {
		return false, err
	}

	return true, nil
}

// func (relationship *[]Relationship) decodeArray(neo4j *Neo4j, data string) (bool, error) {

// 	payload := []RelationshipResponse{}

// 	err := jsonDecode(data, &payload)
// 	// err := json.Unmarshal([]byte(data), payload)
// 	if err != nil {
// 		return false, err
// 	}

// 	for k, v := range payload {
// 		err := mapRelationship(neo4j, relationship, &payload)
// 		if err != nil {
// 			return false, err
// 		}
// 	}

// 	return true, nil
// }

func mapRelationship(neo4j *Neo4j, relationship *Relationship, payload *RelationshipResponse) error {

	relationshipId, err := getIdFromUrl(neo4j.RelationshipUrl, payload.Self)
	if err != nil {
		return err
	}

	startNodeId, err := getIdFromUrl(neo4j.NodeUrl, payload.Start)
	if err != nil {
		return err
	}

	endNodeId, err := getIdFromUrl(neo4j.NodeUrl, payload.End)
	if err != nil {
		return err
	}

	relationship.Id = relationshipId
	relationship.StartNodeId = startNodeId
	relationship.EndNodeId = endNodeId
	relationship.Type = payload.Type
	relationship.Data = payload.Data
	relationship.Payload = payload

	return nil

}
