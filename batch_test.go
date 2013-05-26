package neo4j

import (
	"fmt"
	"testing"
)

func TestBatch(t *testing.T) {
	neo4jConnection := Connect("", 0)
	batch := neo4jConnection.NewBatch()

	node := &Node{}
	relationship := &Relationship{}

	batch.Get(node).Create(node).Delete(node).Update(node).Get(relationship).Create(relationship).Delete(relationship).Update(relationship).Execute()
	fmt.Println("...")
}
