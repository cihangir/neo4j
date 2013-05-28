package neo4j

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Batcher interface {
	getBatchQuery(operation string) (map[string]interface{}, error)
	mapBatchResponse(neo4j *Neo4j, data map[string]interface{}) (bool, error)
}

type Batch struct {
	Neo4j *Neo4j
	Stack []*BatchRequest
}

type BatchRequest struct {
	Operation string
	Data      Batcher
}

type BatchResponse struct {
	Id       int                    `json:"id"`
	Location string                 `json:"location"`
	Body     map[string]interface{} `json:"body"`
	From     string                 `json:"from"`
}

type ManuelBatchRequest struct {
	To   string
	Body map[string]interface{}
}

func (mbr *ManuelBatchRequest) getBatchQuery(operation string) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	query["to"] = mbr.To
	query["body"] = mbr.Body

	switch operation {
	case BATCH_GET:
		query["method"] = "GET"
	case BATCH_UPDATE:
		query["method"] = "PUT"
	case BATCH_CREATE:
		query["method"] = "POST"
	case BATCH_DELETE:
		query["method"] = "DELETE"
	}

	return query, nil
}

func (batch *Batch) GetLastIndex() string {
	return strconv.Itoa(len(batch.Stack) - 1)
}

var (
	BATCH_GET    = "get"
	BATCH_CREATE = "create"
	BATCH_DELETE = "delete"
	BATCH_UPDATE = "update"
)

func (neo4j *Neo4j) NewBatch() *Batch {

	stack := make([]*BatchRequest, 0, 2)
	batch := &Batch{}
	batch.Neo4j = neo4j
	batch.Stack = stack

	return batch
}

func (batch *Batch) Get(obj Batcher) *Batch {
	batch.addToStack(BATCH_GET, obj)
	return batch
}

func (batch *Batch) Create(obj Batcher) *Batch {
	batch.addToStack(BATCH_CREATE, obj)
	return batch
}

func (batch *Batch) Delete(obj Batcher) *Batch {
	batch.addToStack(BATCH_DELETE, obj)
	return batch
}

func (batch *Batch) Update(obj Batcher) *Batch {
	batch.addToStack(BATCH_UPDATE, obj)
	return batch
}

func (batch *Batch) addToStack(operation string, obj Batcher) {

	stack := batch.Stack
	length := len(stack)

	if length+1 > cap(stack) {
		newStack := make([]*BatchRequest, len(stack), (cap(stack)+1)*2) // +1 in case cap(s) == 0
		copy(newStack, stack)
		stack = newStack
	}
	stack = stack[0 : length+1]

	batchRequest := &BatchRequest{}
	batchRequest.Operation = operation
	batchRequest.Data = obj
	stack[len(stack)-1] = batchRequest
	batch.Stack = stack

}

func (batch *Batch) Execute() ([]*BatchResponse, error) {
	if batch.Neo4j == nil {
		return nil, errors.New("Batch request is not created by NewBatch method!")
	}
	// cache batch stack lengh
	stackLength := len(batch.Stack)

	//create result array
	response := make([]*BatchResponse, stackLength)

	if stackLength == 0 {
		return response, nil
	}

	request := make([]map[string]interface{}, stackLength)
	for i, value := range batch.Stack {
		// interface has this method getBatchQuery()
		query, err := value.Data.getBatchQuery(value.Operation)
		if err != nil {
			fmt.Println(err)
			continue
		}
		query["id"] = i
		request[i] = query
	}

	encodedRequest, err := jsonEncode(request)
	res, err := batch.Neo4j.doBatchRequest("POST", batch.Neo4j.BatchUrl, encodedRequest)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(res), &response)
	if err != nil {
		return response, err
	}

	//do mapping here for later usage
	batch.mapResponse(response)
	batch.Stack = make([]*BatchRequest, 0)
	return response, nil
}

func (batch *Batch) mapResponse(response []*BatchResponse) {

	for _, val := range response {
		id := val.Id
		batch.Stack[id].Data.mapBatchResponse(batch.Neo4j, val.Body)
	}
}
