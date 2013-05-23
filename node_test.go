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

	if node != nil {
		t.Error("Node is not nil")
	}

	if err == nil {
		t.Error("Error is nil")
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

func TestGetNodeWithInvalidId(t *testing.T) {
	neo4jConnection := Connect("", 0)
	node, err := neo4jConnection.GetNode("asdfasdfas")

	if node != nil {
		t.Error("Node is not nil")
	}

	if err == nil {
		t.Error("Error is nil")
	}

}

func TestGetNodeWithIdZero(t *testing.T) {
	neo4jConnection := Connect("", 0)
	node, err := neo4jConnection.GetNode("0")

	if node == nil {
		t.Error("Node is nil on valid test")
	}

	if err != nil {
		t.Error("Error is not nil on valid test")
	}

	if node.Id != "0" {
		t.Error("Assigning node id doesnt work")
	}

	if len(node.Data) > 0 {
		t.Error("node data is not nil")
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

func TestGetNodeWithIntMaxId(t *testing.T) {
	maxInt := strconv.Itoa(int(^uint(0) >> 1))
	neo4jConnection := Connect("", 0)
	node, err := neo4jConnection.GetNode(maxInt)

	if node != nil {
		t.Error("Node is not nil")
	}

	if err == nil {
		t.Error("Error is nil")
	}
}
