package neo4j

import (
	"strconv"
	"testing"
)

func TestDefaultConnection(t *testing.T) {
	neo4jConnection := Connect("", 0)

	if neo4jConnection.Client == nil {
		t.Error("Connection client is not set")
	}

	if neo4jConnection.BaseUrl == "" {
		t.Error("BaseUrl is not set")
	}

	if neo4jConnection.NodeUrl == "" {
		t.Error("NodeUrl is not set")
	}

	if neo4jConnection.RelationshipUrl == "" {
		t.Error("RelationshipUrl is not set")
	}

	if neo4jConnection.IndexNodeUrl == "" {
		t.Error("IndexNodeUrl is not set")
	}
}

func TestGetNodeWithEmptyId(t *testing.T) {
	neo4jConnection := Connect("", 0)

	node, err := neo4jConnection.GetNode("")
	if err == nil {
		t.Error("Error is nil")
	}

	if node != nil {
		t.Error("Node is not nil")
	}
}

func TestGetNodeWithInvalidId(t *testing.T) {
	neo4jConnection := Connect("", 0)

	node, err := neo4jConnection.GetNode("asdfasdfas")
	if err == nil {
		t.Error("Error is nil")
	}

	if node != nil {
		t.Error("Node is not nil")
	}
}

func TestGetNodeWithIdZero(t *testing.T) {
	neo4jConnection := Connect("", 0)

	node, err := neo4jConnection.GetNode("0")

	if len(node.Data) > 0 {
		t.Error("node data is not nil")
	}

	if node.Id != "0" {
		t.Error("Assigning node id doesnt work")
	}

	checkForSetValues(t, node, err)
}

func TestGetNodeReturnsNodeObject(t *testing.T) {
	neo4jConnection := Connect("", 0)

	node, _ := neo4jConnection.GetNode("0")

	tt := reflect.TypeOf(node).String()
	// find a better way to check this
	if tt != "*neo4j.Node" {
		t.Error("Response is not a Node object!")
	}
}

func TestGetNodeReturnsErrorObjectOnError(t *testing.T) {
	neo4jConnection := Connect("", 0)

	_, err := neo4jConnection.GetNode("asdfasdfas")

	tt := reflect.TypeOf(err).String()
	// find a better way to check this
	if tt != "*errors.errorString" {
		t.Error("Error is not valid!")
	}
}

func TestGetNodeWithIntMaxId(t *testing.T) {
	maxInt := strconv.Itoa(int(^uint(0) >> 1))
	neo4jConnection := Connect("", 0)

	node, err := neo4jConnection.GetNode(maxInt)
	if err == nil {
		t.Error("Error is nil")
	}

	if node != nil {
		t.Error("Node is not nil")
	}
}

func checkForSetValues(t *testing.T, node *Node, err error) {
	if err != nil {
		t.Error("Error is not nil on valid test")
	}

	if node == nil {
		t.Error("Node is nil on valid test")
	}

	if node.Payload.PagedTraverse == "" {
		t.Error("PagedTraverse on valid node is nil")
	}
	if node.Payload.OutgoingRelationships == "" {
		t.Error("OutgoingRelationships on valid node is nil")
	}

	if node.Payload.Traverse == "" {
		t.Error("Traverse on valid node is nil")
	}

	if node.Payload.AllTypedRelationships == "" {
		t.Error("AllTypedRelationships on valid node is nil")
	}

	if node.Payload.Property == "" {
		t.Error("Property on valid node is nil")
	}

	if node.Payload.AllRelationships == "" {
		t.Error("AllRelationships on valid node is nil")
	}

	if node.Payload.Self == "" {
		t.Error("Self on valid node is nil")
	}

	if node.Payload.Properties == "" {
		t.Error("Properties on valid node is nil")
	}

	if node.Payload.OutgoingTypedRelationships == "" {
		t.Error("OutgoingTypedRelationships on valid node is nil")
	}

	if node.Payload.IncomingRelationships == "" {
		t.Error("IncomingRelationships on valid node is nil")
	}

	if node.Payload.IncomingTypedRelationships == "" {
		t.Error("IncomingTypedRelationships on valid node is nil")
	}

	if node.Payload.CreateRelationship == "" {
		t.Error("CreateRelationship on valid node is nil")
	}
}

func checkForNil(t *testing.T, node *Node) {

	if node.Id != "" {
		t.Error("Id is set")
	}

	if node.Data != nil {
		t.Error("node data is not nil")
	}

	if node.Payload.PagedTraverse != "" {
		t.Error("PagedTraverse on valid node is not nil")
	}
	if node.Payload.OutgoingRelationships != "" {
		t.Error("OutgoingRelationships on valid node is not nil")
	}

	if node.Payload.Traverse != "" {
		t.Error("Traverse on valid node is not nil")
	}

	if node.Payload.AllTypedRelationships != "" {
		t.Error("AllTypedRelationships on valid node is not nil")
	}

	if node.Payload.Property != "" {
		t.Error("Property on valid node is not nil")
	}

	if node.Payload.AllRelationships != "" {
		t.Error("AllRelationships on valid node is not nil")
	}

	if node.Payload.Self != "" {
		t.Error("Self on valid node is not nil")
	}

	if node.Payload.Properties != "" {
		t.Error("Properties on valid node is not nil")
	}

	if node.Payload.OutgoingTypedRelationships != "" {
		t.Error("OutgoingTypedRelationships on valid node is not nil")
	}

	if node.Payload.IncomingRelationships != "" {
		t.Error("IncomingRelationships on valid node is not nil")
	}

	if node.Payload.IncomingTypedRelationships != "" {
		t.Error("IncomingTypedRelationships on valid node is not nil")
	}

	if node.Payload.CreateRelationship != "" {
		t.Error("CreateRelationship on valid node is not nil")
	}
}

func TestCreateNodeWithPassingInvalidObject(t *testing.T) {
	t.Log("complete this method")
}

func TestCreateNodeWithPassingValidObjectAndData(t *testing.T) {

	node := &Node{}
	data := make(map[string]interface{})
	data["stringData"] = "firstData"
	data["integerData"] = 3
	data["floatData"] = 3.0
	node.Data = data

	neo4jConnection := Connect("", 0)
	res, err := neo4jConnection.CreateNode(node)
	testCreatedNodeDeafultvalues(t, node, res, err)
	t.Log("test integer values, all numbers are in float64 format")

	if node.Data["stringData"] != "firstData" {
		t.Error("string value has changed")
	}

	checkForSetValues(t, node, err)

}

func TestCreateNodeWithPassingValidObjectAndEmptyData(t *testing.T) {

	node := &Node{}
	neo4jConnection := Connect("", 0)
	res, err := neo4jConnection.CreateNode(node)
	testCreatedNodeDeafultvalues(t, node, res, err)

	if len(node.Data) != 0 {
		t.Error("node data len must be 0")
	}

	checkForSetValues(t, node, err)
}

func testCreatedNodeDeafultvalues(t *testing.T, node *Node, res bool, err error) {
	if !res {
		t.Error("node creation failed")
	}

	if err != nil {
		t.Error("node creation returned err")
	}

	if node.Id == "" {
		t.Error("Assigning node id doesnt work")
	}
}

	if err == nil {
		t.Error("Error is nil")
	}
}
