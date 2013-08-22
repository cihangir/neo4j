package neo4j

import (
	"testing"
)

func TestSendCypherQuery(t *testing.T) {
	neo4jConnection := Connect("")

	cyper := &Cypher{
		Query  : `
			START group=node:koding(id={groupId})
			MATCH group-[r:member]->members<-[:follower]-currentUser
			WHERE currentUser.id = {currentUserId}
			RETURN members
			ORDER BY {orderByQuery} DESC
			SKIP {skipCount}
			LIMIT {limitCount}
		`
		Params : map[string]{

		}
	}
	data := make(map[string]interface{})
	data["hede"] = "debe"

	node := &Node{}
	node.Data = data
	node2 := &Node{}
	node2.Data = data

	//create batch request for node
	batch := neo4jConnection.NewBatch()
	batch.Create(node)
	batch.Create(node2)
	batch.Execute()
}
