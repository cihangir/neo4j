package neo4j

import "testing"

func TestSendCypherQuery(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]string{
			"query": `
        START k=node(2635, 2637)
        return k.event as event, id(k) as eventNodeId
		  `,
		},
		Payload: map[string]interface{}{},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	_, err := batch.Execute()
	if err != nil {
		t.Error(err)
	}

	if cypher.Payload.(map[string]interface{})["data"] == nil {
		t.Error("no data")
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
