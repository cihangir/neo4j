package neo4j

import (
	"errors"
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

	payload, err := relationship.decode(response)
	if err != nil {
		return nil, err
	}

	return relationship, nil
}

// creates a unique Relationship with given id and Relationship name
// response will be Object
func (neo4j *Neo4j) CreateRelationship(relationship *Relationship) (bool, error) {

	if relationship.StartNodeId == "" {
		return false, errors.New("Start Node Id not valid")
	}

	if relationship.EndNodeId == "" {
		return false, errors.New("End Node Id not valid")
	}

	url := neo4j.NodeUrl + "/" + relationship.StartNodeId + "/relationships"

	endNodeUrl := neo4j.NodeUrl + "/" + relationship.EndNodeId

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

	if Relationship.Id == nil {
		return false, errors.New("Invalid Relationship id")
	}

	postData, err := relationship.encodeData()
	if err != nil {
		return false, err
	}

	url := neo4j.RelationshipUrl + "/" + relationship.Id + "/properties"

	response, err := neo4j.doRequest("PUT", url, postData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (neo4j *Neo4j) DeleteRelationship(id string) (bool, error) {

	url := neo4j.RelationshipUrl + "/" + id

	//if Relationship not found Neo4j returns 404
	response, err := neo4j.doRequest("DELETE", url, "")
	if err != nil {
		return false, err
	}

	return true, nil
}

func (neo4j *Neo4j) GetRelationshipTypes() ([]string, error) {

	url := neo4j.RelationshipUrl + "/types"

	result := make([]string, 0)

	response, err := neo4j.doRequest("GET", url, "")
	if err != nil {
		return result, err
	}

	err = jsonDecode(response, &result)

	return result, err
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

func (relationship *Relationship) encodeData() string {
	result, err := jsonEncode(relationship.Data)
	return result, err
}

func (relationship *Relationship) decode(neo4j *Neo4j, data string) (bool, error) {

	payload := RelationshipResponse{}

	err := jsonDecode(data, &payload)
	// err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		return false, err
	}

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
// 		err := mapRelationship(neo4j, relationship, payload)
// 		if err != nil {
// 			return false, err
// 		}
// 	}

// 	return true, nil
// }

func mapRelationship(neo4j *Noe4j, relationship *Relationship, payload RelationshipResponse) error {

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

	relationship.Id = id
	relationship.StartNodeId = startNodeId
	relationship.EndNodeId = endNodeId
	relationship.Type = payload.Type
	relationship.Data = payload.Data
	relationship.Payload = payload

	return nil

}
