package neo4j

import "testing"

func TestSendCypherQuery(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]string{
			"query": `
        START k=node(*)
        return k
		  `,
		},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload == nil {
		t.Error("No cypher results")
	}
}
