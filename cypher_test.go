package neo4j

import "testing"

func TestSendCypherQuery(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]string{
			"query": `
        START k=node(2635, 2637)
        return k.event
		  `,
		},
		Payload: []string{},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload == nil {
		t.Error("No cypher results")
	}
}

func TestSendCypherQueryWithNoResults(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]string{
			"query": `
        START k=node(239494)
        return k
		  `,
		},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload != nil {
		t.Error("Got cypher results")
	}
}
