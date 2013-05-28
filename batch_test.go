package neo4j

import (
	"reflect"
	"testing"
)

func TestBatchCreation(t *testing.T) {
	neo4jConnection := Connect("", 0)
	batch := neo4jConnection.NewBatch()
	tt := reflect.TypeOf(batch).String()
	if tt != "*neo4j.Batch" {
		t.Error("Batch instance is not valid!")
	}
}

func TestBatchWithOneNode(t *testing.T) {
	neo4jConnection := Connect("", 0)
	batch := neo4jConnection.NewBatch()

	node := &Node{}
	data := make(map[string]interface{})
	data["hede"] = "debe"
	node.Data = data

	batch.Create(node)
	res, err := batch.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) < 1 {
		t.Error("Response length is not valid")
	}

	if node.Id == "" {
		t.Error("Node id is empty")
	}

	if node.Data == nil {
		t.Error("Data id is empty")
	}

	if node.Payload == nil {
		t.Error("Payload id is empty")
	}

}

func TestBatchWithOneRelationship(t *testing.T) {
	neo4jConnection := Connect("", 0)

	//create node
	node := &Node{}
	data := make(map[string]interface{})
	data["hede"] = "debe"
	node.Data = data

	//create node
	node2 := &Node{}
	data2 := make(map[string]interface{})
	data2["hede"] = "debe"
	node2.Data = data

	//create batch request for node
	batchNode := neo4jConnection.NewBatch()
	batchNode.Create(node)
	batchNode.Create(node2)
	res, err := batchNode.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 2 {
		t.Error(len(res), "Response length is not valid")
	}

	//create batch request for relationship
	batchRel := neo4jConnection.NewBatch()

	//create relationship
	relationship := &Relationship{}
	dataRel := make(map[string]interface{})
	dataRel["dada"] = "gaga"
	relationship.Data = dataRel
	relationship.Type = "sampleType"
	relationship.StartNodeId = node.Id
	relationship.EndNodeId = node2.Id

	batchRel.Create(relationship)

	res, err = batchRel.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error(len(res), "Response length is not valid")
	}
	if relationship.Id == "" {
		t.Error("Relationhip is not created")
	}
}

func TestBatchWithNodeForAllRequests(t *testing.T) {
	neo4jConnection := Connect("", 0)

	//create node
	node := &Node{}
	data := make(map[string]interface{})
	data["hede"] = "debe"
	node.Data = data

	//create batch request for node
	batchNode := neo4jConnection.NewBatch()

	batchNode.Create(node)
	res, err := batchNode.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error(len(res), "Response length is not valid")
	}

	if node.Data["hede"] != "debe" {
		t.Error("Node data is not valid")
	}

	if node.Payload == nil {
		t.Error("Payload in nil")
	}

	data["seconData"] = "secondVariable"
	node.Data = data

	//create batch request for node
	updateBatchNode := neo4jConnection.NewBatch()
	updateBatchNode.Update(node)
	res, err = updateBatchNode.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error(len(res), "Response length is not valid")
	}

	if len(node.Data) != 2 {
		t.Error("Node data is not valid")
	}

	if node.Payload == nil {
		t.Error("Payload in nil")
	}

	data["deleteNode"] = "yes"
	node.Data = data

	node2 := node
	//create batch request for node
	deleteBatchNode := neo4jConnection.NewBatch()
	deleteBatchNode.Delete(node2)
	res, err = deleteBatchNode.Execute()
	if err != nil {
		t.Error(err)
	}

	getBatchNode := neo4jConnection.NewBatch()
	getBatchNode.Get(node2)
	res, err = getBatchNode.Execute()
	if err == nil {
		t.Error("trying to get non-existent node succeeded")
	}
}

func TestBatchWithRelationshipForAllRequests(t *testing.T) {
	neo4jConnection := Connect("", 0)

	data := make(map[string]interface{})
	data["hede"] = "debe"

	//create node
	node := &Node{}
	node.Data = data
	//copy node
	node2 := node

	//create batch request for node
	batchNode := neo4jConnection.NewBatch()
	batchNode.Create(node)
	batchNode.Create(node2)
	res, err := batchNode.Execute()
	if err != nil {
		t.Error(err)
	}
	if len(res) != 2 {
		t.Error(len(res), "Response length is not valid")
	}

	//create batch request for relationship
	batchRel := neo4jConnection.NewBatch()
	//create relationship
	relationship := &Relationship{}
	dataRel := make(map[string]interface{})
	dataRel["dada"] = "gaga"
	relationship.Data = dataRel
	relationship.Type = "sampleType"
	relationship.StartNodeId = node.Id
	relationship.EndNodeId = node2.Id

	batchRel.Create(relationship)
	res, err = batchRel.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error(len(res), "Response length is not valid")
	}

	if relationship.Id == "" {
		t.Error("Relationship id is not set")
	}

	if relationship.Data["dada"] != "gaga" {
		t.Error("relationship data is not valid")
	}

	if relationship.Payload == nil {
		t.Error("Payload in nil")
	}

	data["seconData"] = "secondVariable"
	relationship.Data = data

	//create batch request for relationship
	updateBatchRelationship := neo4jConnection.NewBatch()
	updateBatchRelationship.Update(relationship)
	res, err = updateBatchRelationship.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error(len(res), "Response length is not valid")
	}

	if len(relationship.Data) != 2 {
		t.Error("relationship data is not valid")
	}

	if relationship.Payload == nil {
		t.Error("Payload in nil")
	}

	dataRel["deleteRelationship"] = "yes"
	relationship.Data = dataRel

	relationship2 := relationship
	//create batch request for relationship
	deleteBatchRelationship := neo4jConnection.NewBatch()
	deleteBatchRelationship.Delete(relationship2)
	res, err = deleteBatchRelationship.Execute()
	if err != nil {
		t.Error(err)
	}

	getBatchRelationship := neo4jConnection.NewBatch()
	getBatchRelationship.Get(relationship2)
	res, err = getBatchRelationship.Execute()
	if err == nil {
		t.Error("trying to get non-existent relationship succeeded")
	}
}
