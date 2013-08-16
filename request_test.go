// +build graphity

package neo4j

import "testing"

func TestCreateAndDeleteEvents(t *testing.T) {
	sourceId, eventId := createNodes("source", "event", t)

	req := &ManuelRequest{
		Method: "POST",
		To:     "http://localhost:7474/graphity/events",
		Params: map[string]string{
			"source": sourceId,
			"event":  eventId,
		},
	}

	_, err := Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}

	req = &ManuelRequest{
		Method: "DELETE",
		To:     "http://localhost:7474/graphity/events",
		Params: map[string]string{
			"source": sourceId,
			"event":  eventId,
		},
	}

	_, err = Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateAndDeleteSubscriptions(t *testing.T) {
	streamId, sourceId := createNodes("stream", "source", t)

	req := &ManuelRequest{
		Method: "POST",
		To:     "http://localhost:7474/graphity/subscriptions",
		Params: map[string]string{
			"stream": streamId,
			"source": sourceId,
		},
	}

	_, err := Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}

	req = &ManuelRequest{
		Method: "DELETE",
		To:     "http://localhost:7474/graphity/subscriptions",
		Params: map[string]string{
			"stream": streamId,
			"source": sourceId,
		},
	}

	_, err = Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}
}

func TestGetEvents(t *testing.T) {
	streamId, sourceId := createNodes("stream", "source", t)

	req := &ManuelRequest{
		Method: "POST",
		To:     "http://localhost:7474/graphity/subscriptions",
		Params: map[string]string{
			"stream": streamId,
			"source": sourceId,
		},
	}

	_, err := Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}

	_, eventId := createNodes("_", "event", t)

	req = &ManuelRequest{
		Method: "POST",
		To:     "http://localhost:7474/graphity/events",
		Params: map[string]string{
			"source":    sourceId,
			"event":     eventId,
			"timestamp": "1",
		},
	}

	_, err = Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}

	_, eventId = createNodes("_", "event", t)

	req = &ManuelRequest{
		Method: "POST",
		To:     "http://localhost:7474/graphity/events",
		Params: map[string]string{
			"source":    sourceId,
			"event":     eventId,
			"timestamp": "2",
		},
	}

	_, err = Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}

	req = &ManuelRequest{
		Method: "GET",
		To:     "http://localhost:7474/graphity/events",
		Params: map[string]string{
			"stream": streamId,
			"count":  "10",
		},
	}

	nodeIds, err := Connect("").Request(req)
	if err != nil {
		t.Error(err)
	}

	if len(nodeIds) < 2 {
		t.Error("not enough results")
	}
}

func createNodes(nodeOneName, nodeTwoName string, t *testing.T) (nodeOneId, nodeTwoId string) {
	nodeOne := &Node{
		Data: map[string]interface{}{"name": nodeOneName},
	}

	nodeTwo := &Node{
		Data: map[string]interface{}{"name": nodeTwoName},
	}

	batch := Connect("").NewBatch()
	_, err := batch.Create(nodeOne).
		Create(nodeTwo).
		Execute()

	if err != nil {
		t.Error(err)
	}

	return nodeOne.Payload.Self, nodeTwo.Payload.Self
}
