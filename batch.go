package neo4j

import (
	"errors"
	"fmt"
)

type Batch struct {
	Neo4j *Neo4j
	Stack []*BatchRequest
}

type BatchRequest struct {
	Operation string
	Data      interface{}
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

func (batch *Batch) Get(obj interface{}) *Batch {
	batch.addToStack(BATCH_GET, obj)
	return batch
}

func (batch *Batch) Create(obj interface{}) *Batch {
	batch.addToStack(BATCH_CREATE, obj)
	return batch
}

func (batch *Batch) Delete(obj interface{}) *Batch {
	batch.addToStack(BATCH_DELETE, obj)
	return batch
}

func (batch *Batch) Update(obj interface{}) *Batch {
	batch.addToStack(BATCH_UPDATE, obj)
	return batch
}

func (batch *Batch) addToStack(operation string, obj interface{}) {

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

func (batch *Batch) Execute() ([]interface{}, error) {
	if batch.Neo4j == nil {
		return nil, errors.New("Batch request is not created by NewBatch method!")
	}
	// cache batch stack lengh
	stackLength := len(batch.Stack)

	//create result array
	result := make([]interface{}, stackLength)

	if stackLength == 0 {
		return result, nil
	}

	for i, value := range batch.Stack {
		fmt.Println(value.Operation)
		result[i] = value.Data
	}

	return result, nil
}
